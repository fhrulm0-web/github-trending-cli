package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestClient_FetchRepositories_Success(t *testing.T) {
	mockResponse := `{
		"total_count": 2,
		"incomplete_results": false,
		"items": [
			{
				"id": 1,
				"name": "react",
				"full_name": "facebook/react",
				"description": "The library for web and native user interfaces",
				"html_url": "https://github.com/facebook/react",
				"stargazers_count": 220000,
				"language": "JavaScript"
			},
			{
				"id": 2,
				"name": "go",
				"full_name": "golang/go",
				"description": "The Go programming language",
				"html_url": "https://github.com/golang/go",
				"stargazers_count": 120000,
				"language": "Go"
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search/repositories" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("User-Agent") != DefaultUserAgent {
			t.Errorf("expected User-Agent %s, got %s", DefaultUserAgent, r.Header.Get("User-Agent"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := NewClient(server.URL, server.Client())
	repos, err := client.FetchRepositories(context.Background(), DurationWeek, 10, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repos) != 2 {
		t.Fatalf("expected 2 repositories, got %d", len(repos))
	}

	if repos[0].FullName != "facebook/react" || repos[0].StargazersCount != 220000 {
		t.Errorf("unexpected repo 0: %+v", repos[0])
	}
	if repos[1].Language != "Go" {
		t.Errorf("unexpected language for repo 1: %s", repos[1].Language)
	}
}

func TestClient_FetchRepositories_RateLimit(t *testing.T) {
	resetTimestamp := time.Now().Add(15 * time.Minute).Unix()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(resetTimestamp, 10))
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"message": "API rate limit exceeded"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, server.Client())
	_, err := client.FetchRepositories(context.Background(), DurationDay, 10, "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "rate limit exceeded") {
		t.Errorf("expected rate limit error, got: %v", err)
	}
}

func TestClient_FetchRepositories_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "Internal Server Error"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, server.Client())
	_, err := client.FetchRepositories(context.Background(), DurationWeek, 10, "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "500") {
		t.Errorf("expected status 500 error, got: %v", err)
	}
}

func TestClient_FetchRepositories_NoResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"total_count": 0, "items": []}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, server.Client())
	_, err := client.FetchRepositories(context.Background(), DurationYear, 10, "nonexistentlang123")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrNoRepositories {
		t.Errorf("expected ErrNoRepositories, got: %v", err)
	}
}

func TestClient_FetchRepositories_InvalidLimit(t *testing.T) {
	client := NewClient("", nil)

	_, err := client.FetchRepositories(context.Background(), DurationWeek, 0, "")
	if err != ErrInvalidLimit {
		t.Errorf("expected ErrInvalidLimit for limit 0, got %v", err)
	}

	_, err = client.FetchRepositories(context.Background(), DurationWeek, 101, "")
	if err != ErrInvalidLimit {
		t.Errorf("expected ErrInvalidLimit for limit 101, got %v", err)
	}
}

func TestClient_FetchRepositories_NetworkFailure(t *testing.T) {
	// Point to an invalid unreachable address
	client := NewClient("http://127.0.0.1:0", &http.Client{Timeout: 100 * time.Millisecond})
	_, err := client.FetchRepositories(context.Background(), DurationWeek, 10, "")
	if err == nil {
		t.Fatal("expected error on connection failure, got nil")
	}

	if err != ErrNetworkTimeout {
		t.Errorf("expected ErrNetworkTimeout, got %v", err)
	}
}
