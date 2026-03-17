# Commita

> Born from a friend's idea and the thought that interfaces for Git aren't nice, Commita's here to revolutionize.

[![Go](https://img.shields.io/github/go-mod/go-version/gabrielEFagundes/commita?logo=go&logoColor=white)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/gabrielEFagundes/commita?logo=github)](https://github.com/gabrielEFagundes/commita/releases)
[![License](https://img.shields.io/github/license/gabrielEFagundes/commita)](LICENSE)

---

## Overview

Commita is a framework and a wrapper for Git, you can call it as you wish.

Commita's sole purpose is to simplify Git, starting from the process of commiting, while also featuring self-healing to most of Git Errors.

## Features

- **Quicker Pipeline** — Commita simplifies the process of commiting and pushing
- **Self-Healing** — Commita is capable of self-healing errors
    - This is a **WORK IN PROGRESS**! Do not trust Commita for big projects yet!

- **Helpful Commands** — Commita features a few pre-defined semantical commit options, with also the idea of allowing you to define your own semantical commits
- **Simple Configurations** — Commita creates your own config folder `.commita` for all your projects, helping you keep your preferred settings to each project you have.

---

## Installation

> Commita has a installer being worked on, this part can change at any moment!

First, clone this repository:

```bash
git clone https://github.com/gabrielEFagundes/commita
```

Then simply build `core.go` from this repository to your prefered folder and add it to your PATH.

```bash
go build -o /path/to/your/folder/commita.exe core.go # THIS WILL BE CHANGED SOON
```

- On windows, you can use path manager
- On Linux, you can use the following command:

```bash
export PATH=$PATH:/path/to/your/folder
```

> Note that this is only temporary, you need to add that same line to your ~/.profile file for permanent changes.

- On MacOS, add the path to commita into the `/etc/paths` file.

---

## Quick start

Start by setuping your own configurations:

```bash
commita config set -u "https://your-repo-url/" -b "your-default-branch"
```

Then, define your git login for the project:

```bash
commita config login --local/--global "your-user" "your-email"
```

You're ready to Go now!

---

## Usage

### Basic example

```bash
commita commit --feat "Adding new feature"

# ✨ feat: Adding new feature
```

```bash
commita commit --type "perf" "Upgrading performance"

# perf: Upgrading performance
```

### Configuration

```json
{
    "defaultBranch": "main",
    "useEmoji": true,
    "useAi": false,
    "url": "",
    "login": {
        "user": "",
        "email": "",
        "abundance": ""
    }
}
```

Every configuration can be changed from the JSON or with command lines.

!!! tip "Pro tip"
    You don't have to parse your branch if you don't want to, Commita's default branch on the config file is `main`.

### Error handling

```bash
commita commit --docs "Updating README.md"

Initializing your repository...             [ OK ]
Adding your remote to the origin...         [ OK ]
Staging your changes...                     [ OK ]
Commiting your changes...                   [ OK ]
Changing to your branch...                  [ OK ]
Pushing your changes...                     [ FATAL! ]
You have changes on your remote that aren't in your local repo!
Would you like me to pull your remote data? (y/n) yes

Done!
```

---

## Examples

> Examples are coming soon!

---

## Contributing

Contributions are welcome! Please read the [contributing guide](CONTRIBUTING.md) before opening a PR.

---

## License

[MIT](https://github.com/gabrielEFagundes/commita/blob/main/LICENSE) © Gabriel Ehrat Fagundes