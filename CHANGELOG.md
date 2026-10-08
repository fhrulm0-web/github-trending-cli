# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-10-08

### Added
- Core CLI tool `ghtrend` implemented in Go using Cobra.
- Integration with GitHub Search API (`/search/repositories`) sorted by stars.
- Filter flags: `--duration` (`day`, `week`, `month`, `year`), `--limit` (1-100), and `--language`.
- Terminal output formatter using Go stdlib `text/tabwriter` with formatted numbers and description truncation.
- Robust error handling for rate limits (HTTP 403), connection failures, and input validation with exit code 1.
- Comprehensive unit tests across `github`, `display`, and `cmd` packages.
- Initial project repository setup:
  - Configuration files (`.gitignore`, `.gitattributes`, `.editorconfig`).
  - Licensing (`LICENSE` - MIT).
  - Documentation (`README.md`, `CONTRIBUTING.md`, `CHANGELOG.md`).
  - GitHub issue and pull request templates.
