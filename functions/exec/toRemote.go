package exec

import (
	"fmt"
	"os/exec"

	"github.com/gabrielefagundes/commita/exceptions"
	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/gabrielefagundes/commita/internal"
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
		var output []byte
		var gitErr *internal.CommitaErr

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
			output, err = exec.Command("git", arg.Cmd...).CombinedOutput()
		}

		if err != nil {
			gitErr = exceptions.DiagnoseErr(string(output), conf)
			fmt.Print(gitErr.Msg)

			solve := exceptions.AttemptSolve(gitErr.Type, conf)
			if solve != nil {
				return "", solve
			}

		} else {
			fmt.Print("[ OK ]")
		}
	}

	return "\nDone!", nil
}
