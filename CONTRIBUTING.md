# Contributing to github-trending-cli

Thank you for considering contributing to `github-trending-cli`! Here are the guidelines to help get you started.

---

## Code of Conduct

Please be respectful and constructive in all discussions, issues, and pull requests.

---

## How Can I Contribute?

### Reporting Bugs

Before creating a bug report, please check existing issues to ensure the problem hasn't already been reported.

When creating an issue, please include:
- A clear and descriptive title.
- Steps to reproduce the behavior.
- Expected vs. actual behavior.
- Your OS, terminal, and runtime version (e.g. Node.js or Python).

### Suggesting Enhancements

Feature requests are always welcome! Please open an issue using the feature request template and provide:
- A clear description of the proposed feature.
- Why this feature would be useful to users.
- Potential CLI syntax or UI mockups if applicable.

### Submitting Pull Requests

1. Fork the repository and create your branch from `main`:
   ```bash
   git checkout -b feature/my-cool-feature
   ```
2. Follow standard coding standards defined in `.editorconfig`.
3. Write clean, descriptive commit messages adhering to [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat: add filter by stars count`
   - `fix: handle network timeout gracefully`
   - `docs: update command examples in README`
4. Test your changes thoroughly.
5. Push your branch and open a Pull Request against the `main` branch.

---

## Commit Guidelines

We recommend Conventional Commits format:

```text
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `chore`.
