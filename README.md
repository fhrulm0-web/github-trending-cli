# github-trending-cli 🚀

A lightweight and powerful command-line interface (CLI) tool to explore trending GitHub repositories and developers directly from your terminal.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

---

## ✨ Features

- 🌟 **Trending Repositories**: Browse trending repositories across all or specific programming languages.
- 🧑‍💻 **Trending Developers**: Discover top trending developers and creators on GitHub.
- ⏱️ **Flexible Time Ranges**: Filter by `daily`, `weekly`, or `monthly` trends.
- 🎨 **Terminal-Friendly UI**: Beautiful terminal formatting with syntax highlighting, tables, and clickable links.
- ⚡ **Local Caching**: Reduces API calls and increases responsiveness.
- 🌐 **Interactive & Non-Interactive Modes**: Run as a one-shot query or browse interactively.

---

## 📋 Prerequisites

Depending on the chosen implementation stack:
- **Node.js** (>= 18.0.0) or **Python** (>= 3.10)
- **Git**

---

## 🚀 Quick Start

### Installation

Clone the repository:

```bash
git clone https://github.com/your-username/github-trending-cli.git
cd github-trending-cli
```

### Environment Configuration

Copy the example environment configuration:

```bash
cp .env.example .env
```

*(Optional)* Add your GitHub personal access token to `.env` to prevent rate limiting:

```env
GITHUB_TOKEN=ghp_your_personal_access_token_here
```

---

## 📖 Usage (Preview)

```bash
# View today's trending repositories
ghtrend

# Filter by programming language
ghtrend --language rust

# Filter by time range (daily | weekly | monthly)
ghtrend --since weekly

# Discover trending developers
ghtrend developers --language python

# Interactive mode
ghtrend --interactive
```

---

## 🛠️ Project Structure

```text
github-trending-cli/
├── .github/                 # GitHub templates & workflows
│   ├── ISSUE_TEMPLATE/     # Bug report & feature request templates
│   └── PULL_REQUEST_TEMPLATE.md
├── .editorconfig            # Cross-editor formatting consistency
├── .env.example             # Example environment variables
├── .gitattributes           # Git line endings and file attributes
├── .gitignore               # Ignored files for Git
├── CHANGELOG.md             # Project changelog
├── CONTRIBUTING.md          # Guidelines for contributing
├── LICENSE                  # MIT License
└── README.md                # Project documentation
```

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check out [CONTRIBUTING.md](CONTRIBUTING.md) and open an issue or pull request.

---

## 📝 License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
