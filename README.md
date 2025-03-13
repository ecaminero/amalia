# Amalia

![Go Version](https://img.shields.io/badge/Go-v1.20%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-In%20Development-yellow)

Amalia is an intelligent code analysis agent built in Go that provides automated analysis and feedback for code in Git pull requests.

## 🌟 Overview

Amalia helps development teams improve code quality by automatically analyzing pull requests, detecting issues, suggesting improvements, and ensuring adherence to best practices. It acts as an additional reviewer that provides consistent and thorough feedback.

## ✨ Features

- **Automated PR Analysis**: Scans code changes in pull requests automatically
- **Static Code Analysis**: Identifies code smells, anti-patterns, and potential bugs
- **Best Practices Enforcement**: Ensures code follows team coding standards
- **Security Vulnerability Detection**: Identifies potential security issues in code
- **Performance Impact Assessment**: Evaluates how changes might affect application performance
- **Code Quality Metrics**: Provides metrics on code complexity, test coverage, and maintainability
- **Contextual Feedback**: Delivers suggestions with explanations and learning resources
- **Multiple Git Providers**: Works with GitHub, GitLab, Bitbucket, and more
- **Customizable Rules**: Configure analysis rules based on project-specific requirements

## 🛠️ Technology Stack

- **Go**: Fast and efficient analysis engine
- **Git API Integrations**: Interfaces with various Git providers
- **AST Parsing**: Deep code analysis at the syntax tree level
- **Concurrency Model**: Leverages Go's goroutines for parallel analysis
- **Configuration Management**: YAML-based rule configuration

## 📋 Prerequisites

- Go 1.20 or higher
- Git
- Access tokens for your Git provider (GitHub, GitLab, etc.)

## 🚀 Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/ecaminero/amalia.git
cd amalia

# Build the project
go build -o amalia cmd/main.go

# Make it available in your PATH
mv amalia /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/yourusername/amalia@latest
```

## ⚙️ Configuration

Create a configuration file at `~/action.yml`:

## 🔍 Usage


### GitHub Action Integration

```yaml
# .github/workflows/amalia.yml
name: Amalia Code Analysis
permissions:
  contents: read
  pull-requests: write

on:
  pull_request_target:
    types: [opened, synchronize, reopened]
  pull_request_review_comment:
    types: [created]

concurrency:
  group:
    ${{ github.repository }}-${{ github.event.number || github.head_ref ||
    github.sha }}-${{ github.workflow }}-${{ github.event_name ==
    'pull_request_review_comment' && 'pr_comment' || 'pr' }}
  cancel-in-progress: ${{ github.event_name != 'pull_request_review_comment' }}

jobs:
  review:
    name: AI Reviewer 
    runs-on: 'ubuntu-latest'
    steps:
      - name: Checkout Repo
        uses: actions/checkout@v4
      - name: Ai Revewer
        uses: ./
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          GITHUB_PR_NUMBER: ${{ github.event.pull_request.number }}
        with:
          debug: true
          path_filters: |
            !**/*.lock
            !dist/**
```



## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- The Go community for providing excellent tools and libraries
- All contributors who have helped shape Amalia
- Inspiration from existing code analysis tools like SonarQube, CodeClimate, and Golangci-lint

---

Made with by Edwin

## Contributors
**¡Thanks to all the collaborators who have made this project possible!**

[![Contributors](https://contrib.rocks/image?repo=ecaminero/ai-codereview)](https://github.com/ecaminero/ai-codereview/graphs/contributors)

<p align="right">(<a href="#readme-top">Go back up</a>)</p>

## 🛠️ Stack

- [![Go][go-badge]][go-url] - An open-source programming language supported by Google.

<p align="right">(<a href="#readme-top">Go back up</a>)</p>

[go-url]: https://go.dev/

[githubaction-url]: https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions
[go-badge]: https://img.shields.io/badge/go-fff?style=for-the-badge&logo=go&logoColor=bd303a&color=35256


[contributors-url]: https://github.com/ecaminero/ai-codereview/graphs/contributors
[contributors-shield]: https://img.shields.io/github/contributors/ecaminero/ai-codereview.svg?style=for-the-badge

[forks-url]: https://github.com/ecaminero/ai-codereview/network/members
[forks-shield]: https://img.shields.io/github/forks/ecaminero/ai-codereview.svg?style=for-the-badge

[stars-url]: https://github.com/ecaminero/ai-codereview/stargazers
[stars-shield]: https://img.shields.io/github/stars/ecaminero/ai-codereview.svg?style=for-the-badge

[issues-shield]: https://img.shields.io/github/issues/ecaminero/ai-codereview.svg?style=for-the-badge
[issues-url]: https://github.com/ecaminero/ai-codereview/issues
