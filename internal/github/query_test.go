package github

import (
	"strings"
	"testing"
	"time"
)

func TestCalculateSinceDate(t *testing.T) {
	refTime := time.Date(2026, 10, 8, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		duration    string
		expected    time.Time
		expectError bool
	}{
		{
			name:        "valid day",
			duration:    DurationDay,
			expected:    time.Date(2026, 10, 7, 12, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "valid week",
			duration:    DurationWeek,
			expected:    time.Date(2026, 10, 1, 12, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "valid month",
			duration:    DurationMonth,
			expected:    time.Date(2026, 9, 8, 12, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "valid year",
			duration:    DurationYear,
			expected:    time.Date(2025, 10, 8, 12, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "invalid duration",
			duration:    "century",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CalculateSinceDate(tt.duration, refTime)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error for duration %s, got nil", tt.duration)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !res.Equal(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, res)
			}
		})
	}
}

func TestBuildSearchQuery(t *testing.T) {
	refTime := time.Date(2026, 10, 8, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		duration    string
		language    string
		expected    string
		expectError bool
	}{
		{
			name:        "week without language",
			duration:    "week",
			language:    "",
			expected:    "created:>2026-10-01",
			expectError: false,
		},
		{
			name:        "day with language",
			duration:    "day",
			language:    "go",
			expected:    "language:go created:>2026-10-07",
			expectError: false,
		},
		{
			name:        "month with trimmed language",
			duration:    "month",
			language:    " python ",
			expected:    "language:python created:>2026-09-08",
			expectError: false,
		},
		{
			name:        "invalid duration",
			duration:    "invalid",
			language:    "rust",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, err := BuildSearchQuery(tt.duration, tt.language, refTime)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if query != tt.expected {
				t.Errorf("expected query '%s', got '%s'", tt.expected, query)
			}
			if !strings.Contains(query, "created:>") {
				t.Errorf("query should contain created:>, got '%s'", query)
			}
		})
	}
}
