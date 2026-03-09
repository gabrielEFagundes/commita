## v1.3.0 - 2026.03.09

### Features
Added commit logic, commita can now stage, commit and push changes to a remote repository

Now commita works with a config file, without it, it does not work.
- The configuration file holds the default branch, remote repository URL and if it should use AI and Emojis on the predefined semantical commits.
- The configuration file is located at the same folder as commita, that being the one the user chooses
    - On Windows, the file is located at `%USERPROFILE%/commita/`, on Unix-based systems, such as Linux and MacOS, the file is located at `~/.config/commita/`
    - Note that these directories might change, especially on Unix-based systems.

### Fixes
Commit logic now has a few simple handlers to simple commit errors, such as `status 1: nothing to commit` and a few 2-options commands, such as creating or updating remote repository URL and creating or changing to the branch on the config file.

## v1.2.2 - 2026.02.27

### Changes
Changed main library from `github.com/jonathonwebb/getopt` to `github.com/spf13/cobra`

Changed a bit of the project's structure to adapt to Golang.

## v1.2.1 - 2026.02.20

### Changes
Started Commita's migration from **C** to **Golang**

Created `CHANGELOGS.md`