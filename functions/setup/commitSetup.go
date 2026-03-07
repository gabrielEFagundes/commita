package setup

import (
	"fmt"
	"os/exec"
)

func SetupBranch(b string) error {
	err := exec.Command("git", "checkout", b).Run()
	if err == nil {
		return nil
	}

	return exec.Command("git", "checkout", "-b", b).Run()
}

func SetupRemote(u string) error {
	err := exec.Command("git", "remote", "add", "origin", u).Run()
	if err == nil {
		return nil
	}

	return exec.Command("git", "remote", "set-url", "origin", u).Run()
}

func SetupCommit() error {
	status, _ := exec.Command("git", "status", "--porcelain").Output()
	if len(status) == 0 {
		return nil
	}

	return fmt.Errorf("nothing to commit")
}
