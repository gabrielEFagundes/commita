# Contributing

First off, thanks for taking the time to contribute! Every bit helps, whether it's fixing a typo or adding a new feature.

---

## Before you start

- Check the [open issues](https://github.com/gabrielEFagundes/commita/issues) to see if your idea or bug is already being tracked
- For big changes, open an issue first to discuss — this avoids wasted effort if the direction doesn't fit the project
- For small fixes (typos, docs, minor bugs), just open a PR directly

---

## Setting up the project

**Requirements:**

- Go 1.21+
- Git

**Clone and build:**

```bash
git clone https://github.com/gabrielEFagundes/commita
cd commita
go mod download
go build -o /path/to/commita/commita.exe core.go
```

## Making changes

### Branching

Create a branch from `main` with the first letters of your username and what the branch does:

```bash
git checkout -b gef/some-bug
git checkout -b gef/some-feature
git checkout -b gef/improve-readme
```

### Code style

This project follows standard Go conventions, we have to keep it the Go Way!

```bash
# format your code before committing
gofmt -w .

# run the linter
go vet ./...
```

!!! tip
    If you use VS Code, install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) — it formats on save automatically.

### Commits

Keep commits focused and use clear messages:

```
feat: add support for X
fix: handle nil pointer in Y
docs: update installation steps
test: add coverage for Z
refactor: simplify X logic
```

Emojis are optional, you can use them if you want!

---

## Submitting a pull request

1. Make sure all tests pass;
2. Make sure the code is formatted: `gofmt -w .`
3. Push your branch and open a PR against `main`
4. Fill in the PR description — check out the [Template](https://github.com/gabrielEFagundes/commita/blob/main/docs/templates/PR-TEMPLATE.md)
5. Wait for review — I'll get back to you as soon as possible

!!! note
    PRs that don't work on my local machine will be put on stand-by until the reason for that is found.

---

## Reporting bugs

Open an issue and include:

- Go version (`go version`)
- OS and architecture
- A minimal example that shows off the bug
- What you expected vs what actually happened

Check out for the [Issue Template](https://github.com/gabrielEFagundes/commita/blob/main/docs/templates/ISSUE-TEMPLATE.md)

---

## Suggesting features

Open an issue with the `enhancement` label and describe:

- The problem you're trying to solve
- How you'd expect the feature to work
- Any alternatives you've considered

---

## License

By contributing, you agree that your contributions will be licensed under the same [MIT License](https://github.com/gabrielEFagundes/commita/blob/main/LICENSE) as the project.