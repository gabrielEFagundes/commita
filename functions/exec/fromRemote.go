package exec

import "os/exec"

// for now it's only a git pull, but I'll implement diagnostics and problem solving.
func PullFromRemote(u string) error {
	return exec.Command("git", "pull", u).Run()
}
