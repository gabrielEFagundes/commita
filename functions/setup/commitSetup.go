package setup

import (
	"fmt"
	"os/exec"

	"github.com/gabrielefagundes/commita/structs"
)

func SetupBranch(b string, arg structs.Task) error {
	err := exec.Command("git", arg.Cmd...).Run()
	if err == nil {
		return nil
	}

	return exec.Command("git", "checkout", b).Run()
}

func SetupRemote(u string, arg structs.Task) error {
	err := exec.Command("git", arg.Cmd...).Run()
	if err == nil {
		return nil
	}

	return exec.Command("git", "remote", "set-url", "origin", u).Run()
}

func SetupCommit() error {
	status, _ := exec.Command("git", "status", "--porcelain").Output()
	if len(status) != 0 {
		return nil
	}

	return fmt.Errorf("[ UP TO DATE ]")
}
