package github

import "time"

// Repository represents a GitHub repository returned by the search API.
type Repository struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Description     string    `json:"description"`
	HTMLURL         string    `json:"html_url"`
	StargazersCount int       `json:"stargazers_count"`
	Language        string    `json:"language"`
	ForksCount      int       `json:"forks_count"`
	CreatedAt       time.Time `json:"created_at"`
}

// SearchResult represents the top-level payload returned by GitHub Search Repositories API.
type SearchResult struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

// APIError represents an error payload returned by GitHub API.
type APIError struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}
