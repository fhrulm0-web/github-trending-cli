package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fhrulm0-web/github-trending-cli/internal/github"
)

func newMockServer() *httptest.Server {
	mockResponse := `{
		"total_count": 1,
		"incomplete_results": false,
		"items": [
			{
				"id": 123,
				"name": "trending-project",
				"full_name": "developer/trending-project",
				"description": "An awesome trending tool",
				"html_url": "https://github.com/developer/trending-project",
				"stargazers_count": 4500,
				"language": "Go"
			}
		]
	}`

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockResponse))
	}))
}

func TestRootCmd_Success(t *testing.T) {
	server := newMockServer()
	defer server.Close()

	client := github.NewClient(server.URL, server.Client())
	rootCmd := NewRootCmd(client)

	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs([]string{"--duration", "month", "--limit", "5", "--language", "go"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	outStr := outBuf.String()
	if !strings.Contains(outStr, "developer/trending-project") {
		t.Errorf("expected repo in output, got: %s", outStr)
	}
	if !strings.Contains(outStr, "Trending go repositories (month)") {
		t.Errorf("expected title in output, got: %s", outStr)
	}
}

func TestRootCmd_InvalidDuration(t *testing.T) {
	rootCmd := NewRootCmd(nil)

	errBuf := new(bytes.Buffer)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs([]string{"--duration", "decade"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error for invalid duration, got nil")
	}

	errStr := errBuf.String()
	if !strings.Contains(errStr, "Invalid duration 'decade'. Allowed values: day, week, month, year.") {
		t.Errorf("unexpected error output: %s", errStr)
	}
}

func TestRootCmd_InvalidLimit(t *testing.T) {
	tests := []struct {
		name  string
		limit string
	}{
		{"below min", "0"},
		{"above max", "101"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := NewRootCmd(nil)
			errBuf := new(bytes.Buffer)
			rootCmd.SetErr(errBuf)
			rootCmd.SetArgs([]string{"--limit", tt.limit})

			err := rootCmd.Execute()
			if err == nil {
				t.Fatalf("expected error for limit %s, got nil", tt.limit)
			}

			errStr := errBuf.String()
			if !strings.Contains(errStr, "Limit must be between 1 and 100.") {
				t.Errorf("unexpected error output: %s", errStr)
			}
		})
	}
}

func TestRootCmd_Help(t *testing.T) {
	rootCmd := NewRootCmd(nil)
	outBuf := new(bytes.Buffer)
	rootCmd.SetOut(outBuf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error on --help: %v", err)
	}

	outStr := outBuf.String()
	if !strings.Contains(outStr, "--duration") || !strings.Contains(outStr, "--limit") || !strings.Contains(outStr, "--language") {
		t.Errorf("help output missing flag descriptions:\n%s", outStr)
	}
}
