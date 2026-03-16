## v2.0.0-beta - 2026.03.16

### Features

- Completely rewrited the project, migrating everything from C to Go
- Overhauled the way commita wraps git and decided to change it:
  - Before, Commita used `git-aliases`, which allowed commands like `git feat`, this approach wasn't bad, but I felt a lack of originality, so I changed that
  - Now, Commita uses it's own syntax, `commita-aliases`, which are essentially the same thing, but with my own touch on it. No, I don't think people can make their own aliases yet, but I might implement that one day.
- Added configuration folder for internal projects, `.commita`. With it, there's no need to parse the branch and url everytime we need to commit
- Added Github pages for the project's documentation. Does not include the page itself, only the config.
  - The docs pages are located on `/docs` (will be added and worked on later).

### Fixes

- Updated README to match Commita's new rework
- Updated LICENSE from GNU-3.0 to MIT License

## v1.3.0 - 2026.03.09

### Features
Added commit logic, commita can now stage, commit and push changes to a remote repository

Now commita works with a config file, without it, it does not work.
- The configuration file holds the default branch, remote repository URL and if it should use AI and Emojis on the predefined semantical commits.
- The configuration file is located at the same folder as commita, that being the one the user chooses
    - On Windows, the file is located at `%USERPROFILE%/commita/`, on Unix-based systems, such as Linux and MacOS, the file is located at `~/.config/commita/`
    - Note that these directories might change, especially on Unix-based systems.

Added Git Exception diagnoser.
- Not all exceptions are being diagnosed yet, but most of the usual ones, which are fatal, are.

Added Base for commita to fix errors by itself based on the error type, which is also based on the errors from Git's stdout.

### Fixes
Commit logic now has a few simple handlers to handle commit errors, such as `status 1: nothing to commit` and a few 2-options commands, such as creating or updating remote repository URL and creating or changing to the branch on the config file.

## v1.2.2 - 2026.02.27

### Changes
Changed main library from `github.com/jonathonwebb/getopt` to `github.com/spf13/cobra`

Changed a bit of the project's structure to adapt to Golang.

## v1.2.1 - 2026.02.20

### Changes
Started Commita's migration from **C** to **Golang**

Created `CHANGELOGS.md`
