package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	DefaultBaseURL   = "https://api.github.com"
	DefaultTimeout   = 15 * time.Second
	DefaultUserAgent = "github-trending-cli"
)

// Common error messages matching SPEC.md § 6.
var (
	ErrNetworkTimeout = errors.New("failed to connect to GitHub API. Check your internet connection.")
	ErrNoRepositories = errors.New("no repositories found for the given criteria.")
	ErrInvalidLimit   = errors.New("limit must be between 1 and 100.")
)

// Client handles communication with the GitHub API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	UserAgent  string
}

// NewClient creates a new GitHub API client.
func NewClient(baseURL string, httpClient *http.Client) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: DefaultTimeout,
		}
	}
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: httpClient,
		UserAgent:  DefaultUserAgent,
	}
}

// FetchRepositories queries GitHub Search API for trending repositories.
func (c *Client) FetchRepositories(ctx context.Context, duration string, limit int, language string) ([]Repository, error) {
	if limit < MinLimit || limit > MaxLimit {
		return nil, ErrInvalidLimit
	}

	query, err := BuildSearchQuery(duration, language, time.Now())
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/search/repositories", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	params := url.Values{}
	params.Set("q", query)
	params.Set("sort", "stars")
	params.Set("order", "desc")
	params.Set("per_page", strconv.Itoa(limit))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, ErrNetworkTimeout
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode == http.StatusForbidden {
		resetHeader := resp.Header.Get("X-RateLimit-Reset")
		if resetHeader != "" {
			if resetUnix, parseErr := strconv.ParseInt(resetHeader, 10, 64); parseErr == nil {
				resetTime := time.Unix(resetUnix, 0)
				remainingMinutes := int(time.Until(resetTime).Minutes())
				if remainingMinutes <= 0 {
					remainingMinutes = 1
				}
				return nil, fmt.Errorf("rate limit exceeded. Try again in %d minutes.", remainingMinutes)
			}
		}
		return nil, errors.New("rate limit exceeded. Try again in a few minutes.")
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		if err := json.Unmarshal(bodyBytes, &apiErr); err == nil && apiErr.Message != "" {
			return nil, formatStatusError(resp.StatusCode, apiErr.Message)
		}
		return nil, formatStatusError(resp.StatusCode, resp.Status)
	}

	var searchResult SearchResult
	if err := json.Unmarshal(bodyBytes, &searchResult); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}

	if len(searchResult.Items) == 0 {
		return nil, ErrNoRepositories
	}

	return searchResult.Items, nil
}

func formatStatusError(code int, msg string) error {
	return fmt.Errorf("GitHub API returned status %d: %s", code, msg)
}
