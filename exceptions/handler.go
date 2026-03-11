package exceptions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/gabrielefagundes/commita/internal"
	"github.com/gabrielefagundes/commita/structs"
	"github.com/gabrielefagundes/commita/utils"
	"github.com/logrusorgru/aurora"
)

func DiagnoseErr(out string, conf structs.Config) *internal.CommitaErr {
	patterns := []internal.CommitaErr{
		{Raw: "failed to push some refs", Type: internal.MissingRemoteChanges, Msg: "You have changes on your remote that aren't in your local repo!"},
		{Raw: "cannot do a partial commit during a merge", Type: internal.CommitAfterMerge, Msg: "You have to finish your merge before commiting."},
		{Raw: "insufficient permission", Type: internal.MissingPerms, Msg: "You don't have permissions to write to the remote repository."},
		{Raw: "Another git process seems to be running in this repository", Type: internal.InterruptedOperation, Msg: "There's another Git process happening! Stop that process and try to commit again."},
		{Raw: "Authentication failed", Type: internal.InvalidRemoteCredentials, Msg: "Git couldn't recognize you, please, try login again using commita login"},
		{Raw: "refusing to merge unrelated histories", Type: internal.UnrelatedHists, Msg: "Your local and remote repositories don't have the same commit histories."},
		{Raw: "does not appear to be a git repository", Type: internal.NotARepo, Msg: ("The URL from your repository does not seems to be a Git Repository: " + conf.Url)},
		{Raw: "exceeds GitHub's file size limit", Type: internal.TooBig, Msg: "The files you want to upload are too big!"},
	}

	for _, p := range patterns {
		if strings.Contains(out, p.Raw) {
			return &internal.CommitaErr{
				Raw:  p.Raw,
				Type: p.Type,
				Msg:  fmt.Sprintf("%s\n\n%s", aurora.BrightRed("[ FATAL ]"), p.Msg),
			}
		}
	}

	return &internal.CommitaErr{Raw: "", Type: internal.Unknown, Msg: fmt.Sprintf("%s\nAn unexpected Git error happened: %s", aurora.BrightRed("fatal!"), out)}
}

func AttemptSolve(e internal.GitErr, conf structs.Config) error {
	// only those I think commita can solve

	switch e {
	case internal.NAE:
		return nil

	case internal.MissingRemoteChanges:
		if utils.Confirm("\nWould you like me to pull your remote data?") {
			return exec.Command("git", "pull", conf.Url).Run() // this can cause merging error (yes, it goes straight to 'main|MERGING') TODO
		}

		return fmt.Errorf("\n%s", aurora.BrightRed("Pull cancelled."))

	case internal.InterruptedOperation:
		if utils.Confirm("\nWould you like me to remove '.git/index.lock'?") {
			return os.Remove("./.git/index.lock")
		}

		return fmt.Errorf("\n%s", aurora.BrightRed("Remove cancelled."))

	case internal.InvalidRemoteCredentials:
		if utils.Confirm("\nWould you like to login through commita?") {
			// I need to do this logic yet
		}

		return fmt.Errorf("\n%s", aurora.BrightRed("Login cancelled."))

	case internal.NotARepo:
		if utils.Confirm("\nWould you like to see and change your url configuration?") {
			fmt.Printf("Remote repository URL: %s", conf.Url)

			if utils.Confirm("\nWould you like to change that?") {
				fmt.Scanf("New Url: %s", conf.Url)
				return setup.SaveConfig(conf)
			}
		}

		return fmt.Errorf("\n%s", aurora.BrightRed("Configurations not altered."))

	default:
		return fmt.Errorf("\ncannot solve error %d", e)
	}
}
