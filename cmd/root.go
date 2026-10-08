package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fahrul/github-trending-cli/internal/display"
	"github.com/fahrul/github-trending-cli/internal/github"
	"github.com/spf13/cobra"
)

var (
	// Version is the current version of ghtrend CLI.
	Version = "0.1.0"

	durationFlag string
	limitFlag    int
	languageFlag string
)

// NewRootCmd initializes and returns the root cobra command.
func NewRootCmd(client *github.Client) *cobra.Command {
	if client == nil {
		client = github.NewClient("", nil)
	}

	rootCmd := &cobra.Command{
		Use:           "ghtrend",
		Short:         "Discover trending GitHub repositories right from your terminal",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Validate duration
			if !github.ValidDurations[durationFlag] {
				err := fmt.Errorf("Invalid duration '%s'. Allowed values: day, week, month, year.", durationFlag)
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				return err
			}

			// Validate limit
			if limitFlag < github.MinLimit || limitFlag > github.MaxLimit {
				err := fmt.Errorf("Limit must be between 1 and 100.")
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				return err
			}

			ctx := context.Background()
			repos, err := client.FetchRepositories(ctx, durationFlag, limitFlag, languageFlag)
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				return err
			}

			if err := display.RenderTable(cmd.OutOrStdout(), repos, durationFlag, languageFlag); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				return err
			}

			return nil
		},
	}

	rootCmd.Flags().StringVarP(&durationFlag, "duration", "d", github.DefaultDuration, "Time range: day, week, month, year")
	rootCmd.Flags().IntVarP(&limitFlag, "limit", "l", github.DefaultLimit, "Number of repositories to display (1-100)")
	rootCmd.Flags().StringVar(&languageFlag, "language", "", "Filter by programming language (e.g., go, python, rust)")

	return rootCmd
}

// Execute runs the root CLI command.
func Execute() {
	cmd := NewRootCmd(nil)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
