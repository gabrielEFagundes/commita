<div align="center">
    <h1>Commita</h1>
    <sub>"Why waste time with commiting and pushing, Commita gets you done within seconds".</sub>
    <br>
    <br>
    <img src="https://img.shields.io/github/v/release/gabrielEFagundes/commita?logo=devbox">
    <img src="https://img.shields.io/badge/Go-1.25.0-00ADD8?logo=go&logoColor=white">
    <img src="https://img.shields.io/badge/Open-Source-orange?logo=git&logoColor=white">
    <img src="https://img.shields.io/github/license/gabrielEFagundes/commita">
</div>

## What's Commita?

Commita is a framework and a wrapper for Git, you can call it as you wish.

Being born as an idea from a friend and with the thought that a visual interface isn't that nice (thanks, MQ), especially to those used to CLI-like programs, Commita's sole purpose is to simplify Git, starting from the process of commiting, while also featuring self-healing to most of Git Errors.

## Features

Commita's on a early stage of development, having a complete rewrite from C to Golang, so it still needs some improvement, including the self-healing features.

I do not recommend using Commita on big projects yet, start by using it on side projects, at least until I'm sure commita's not going to delete everything or execute an irreversible action.

As for now, Commita is capable of commiting and pushing changes from a single command line.

## Quick Start

```bash
commita config set -u "your_repo_url" -b "your_default_branch"
commita config login --local/--global "your_user" "your_email"
```

Those commands create Commita's config file for your project. They are necessary for Commita to work.

After that, Commita's ready to Go! ;]

Example:
```bash
commita commit --feat "Adding new feature"

# ✨ feat: Adding new feature
```

If you need further info, check out our [Docs](https://gabrielefagundes.github.io/commita)

## License

You're free to use Commita anyway you'd like, the only restriction is redistribution, please, do not distribute Commita under your own name, give me some credit!

## Contributing

You're free to open [Issues](https://github.com/gabrielEFagundes/commita/issues) and [Pull Requests](https://github.com/gabrielEFagundes/commita/pulls)! Have a look on the [Contributing Page](In-Dev!) for more information.