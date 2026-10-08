package display

import (
	"bytes"
	"strings"
	"testing"

	"github.com/fahrul/github-trending-cli/internal/github"
)

func TestFormatNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{5, "5"},
		{999, "999"},
		{1000, "1,000"},
		{12345, "12,345"},
		{1234567, "1,234,567"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			got := FormatNumber(tt.input)
			if got != tt.expected {
				t.Errorf("FormatNumber(%d) = %s; want %s", tt.input, got, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{"empty string", "", 10, ""},
		{"short string", "hello", 10, "hello"},
		{"exact length", "12345", 5, "12345"},
		{"longer string", "1234567890", 5, "12345…"},
		{"newlines removed", "line 1\nline 2", 20, "line 1 line 2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Truncate(tt.input, tt.maxLen)
			if got != tt.expected {
				t.Errorf("Truncate(%q, %d) = %q; want %q", tt.input, tt.maxLen, got, tt.expected)
			}
		})
	}
}

func TestRenderTable(t *testing.T) {
	repos := []github.Repository{
		{
			FullName:        "facebook/react",
			StargazersCount: 220100,
			Language:        "JavaScript",
			Description:     "A declarative, efficient, and flexible JavaScript library for building user interfaces.",
		},
		{
			FullName:        "golang/go",
			StargazersCount: 110500,
			Language:        "Go",
			Description:     "The Go programming language",
		},
		{
			FullName:        "unspecified/repo",
			StargazersCount: 42,
			Language:        "",
			Description:     "No language specified",
		},
	}

	var buf bytes.Buffer
	err := RenderTable(&buf, repos, "week", "Go")
	if err != nil {
		t.Fatalf("RenderTable returned error: %v", err)
	}

	output := buf.String()

	// Check title
	if !strings.Contains(output, "Trending Go repositories (week)") {
		t.Errorf("expected title in output, got:\n%s", output)
	}

	// Check table headers
	if !strings.Contains(output, "Repository") || !strings.Contains(output, "Stars") {
		t.Errorf("missing headers in output:\n%s", output)
	}

	// Check repository contents
	if !strings.Contains(output, "facebook/react") || !strings.Contains(output, "220,100") {
		t.Errorf("missing repo react data in output:\n%s", output)
	}
	if !strings.Contains(output, "golang/go") || !strings.Contains(output, "110,500") {
		t.Errorf("missing repo go data in output:\n%s", output)
	}

	// Check empty language fallback to "-"
	if !strings.Contains(output, "-") {
		t.Errorf("expected empty language to be '-', got:\n%s", output)
	}

	// Check footer
	if !strings.Contains(output, "Showing 3 results") {
		t.Errorf("missing footer summary in output:\n%s", output)
	}
}

func TestRenderTable_Empty(t *testing.T) {
	var buf bytes.Buffer
	err := RenderTable(&buf, []github.Repository{}, "week", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "No repositories to display.") {
		t.Errorf("expected empty notice, got: %s", buf.String())
	}
}
