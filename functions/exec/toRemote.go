package exec

import (
	"fmt"
	"os/exec"

	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/gabrielefagundes/commita/structs"
)

func CommitAndPush(commitMsg string) error {
	conf, errConf := setup.LoadConfig()
	if errConf != nil {
		return errConf
	}

	cmds := structs.MountTask(commitMsg, conf.DefaultBranch)

	for _, arg := range cmds {
		fmt.Printf("-> %s", arg.Label)

		var err error
		switch arg.Cmd[0] {
		case "remote":
			err = setup.SetupRemote(conf.Url)

		case "commit":
			status := setup.SetupCommit()
			if status == nil {
				// need to find a way to go through this step
			}
			continue

		case "checkout":
			err = setup.SetupBranch(conf.DefaultBranch)

		default:
			err = exec.Command("git", arg.Cmd...).Run()
		}

		if err != nil {
			return fmt.Errorf("\nfailed trying %s: %w", arg.Cmd[0], err)
		}

		fmt.Println("\t\t\t[ OK ]")
	}

	return nil
}
