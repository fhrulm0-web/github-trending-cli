package display

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/fhrulm0-web/github-trending-cli/internal/github"
)

const (
	MaxDescriptionLength = 60
	TablePadding         = 3
)

// FormatNumber formats an integer with comma separators (e.g. 12345 -> "12,345").
func FormatNumber(n int) string {
	str := fmt.Sprintf("%d", n)
	if len(str) <= 3 {
		return str
	}

	var result []byte
	offset := len(str) % 3
	if offset == 0 {
		offset = 3
	}

	result = append(result, str[:offset]...)
	for i := offset; i < len(str); i += 3 {
		result = append(result, ',')
		result = append(result, str[i:i+3]...)
	}

	return string(result)
}

// Truncate safely truncates a string to maxLen runes and appends an ellipsis if truncated.
func Truncate(s string, maxLen int) string {
	// Clean newlines or carriage returns in description
	clean := strings.ReplaceAll(s, "\r\n", " ")
	clean = strings.ReplaceAll(clean, "\n", " ")
	clean = strings.ReplaceAll(clean, "\t", " ")
	clean = strings.TrimSpace(clean)

	runes := []rune(clean)
	if len(runes) <= maxLen {
		return clean
	}

	return string(runes[:maxLen]) + "…"
}

// RenderTable writes the formatted trending repositories table to the provided writer.
func RenderTable(w io.Writer, repos []github.Repository, duration string, language string) error {
	if len(repos) == 0 {
		_, err := fmt.Fprintln(w, "No repositories to display.")
		return err
	}

	// Title
	var title string
	if language != "" {
		title = fmt.Sprintf("Trending %s repositories (%s) — sorted by stars", language, duration)
	} else {
		title = fmt.Sprintf("Trending repositories (%s) — sorted by stars", duration)
	}

	divider := strings.Repeat("─", 80)

	fmt.Fprintln(w, title)
	fmt.Fprintln(w, divider)

	tw := tabwriter.NewWriter(w, 0, 0, TablePadding, ' ', 0)
	fmt.Fprintln(tw, " #\tRepository\tStars\tLanguage\tDescription")
	fmt.Fprintf(tw, " %s\t%s\t%s\t%s\t%s\n", "─", "────────────────────", "───────", "──────────", "────────────────────────────")

	for i, repo := range repos {
		lang := repo.Language
		if lang == "" {
			lang = "-"
		}
		desc := Truncate(repo.Description, MaxDescriptionLength)
		stars := FormatNumber(repo.StargazersCount)

		fmt.Fprintf(tw, " %d\t%s\t%s\t%s\t%s\n", i+1, repo.FullName, stars, lang, desc)
	}

	if err := tw.Flush(); err != nil {
		return err
	}

	fmt.Fprintln(w, divider)
	fmt.Fprintf(w, "Showing %d results | Source: GitHub Search API | Duration: %s\n", len(repos), duration)

	return nil
}
