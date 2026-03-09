package exec

import (
	"fmt"
	"os/exec"

	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/gabrielefagundes/commita/structs"
	"github.com/gabrielefagundes/commita/utils"
)

func CommitAndPush(commitMsg string) (string, error) {
	conf, errConf := setup.LoadConfig()
	if errConf != nil {
		return "", fmt.Errorf("error in configuration: %w", errConf)
	}

	cmds := structs.MountTask(commitMsg, conf.DefaultBranch, conf.Url)

	for _, arg := range cmds {
		label := fmt.Sprintf("\n-> %s", arg.Label)
		fmt.Print(utils.LeftAlign(label, 50))

		var err error
		switch arg.Cmd[0] {
		case "remote":
			err = setup.SetupRemote(conf.Url, arg)

		case "commit":
			status := setup.SetupCommit()
			if status != nil {
				fmt.Printf("%s", status)
				continue
			}
			err = exec.Command("git", arg.Cmd...).Run()

		case "checkout":
			err = setup.SetupBranch(conf.DefaultBranch, arg)

		default:
			err = exec.Command("git", arg.Cmd...).Run()
		}

		if err != nil {
			return "", fmt.Errorf("\nfailed trying %s: %w", arg.Cmd[0], err)
		}

		fmt.Print("[ OK ]")
	}

	return "\nDone!", nil
}
