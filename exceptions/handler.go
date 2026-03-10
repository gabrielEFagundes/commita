package exceptions

import (
	"fmt"
	"strings"

	"github.com/gabrielefagundes/commita/structs"
	"github.com/logrusorgru/aurora"
)

func DiagnoseErr(out string, conf structs.Config) *CommitaErr {
	patterns := []CommitaErr{
		{"failed to push some refs", MissingRemoteChanges, "You have changes on your remote that aren't in your local repo!"},
		{"cannot do a partial commit during a merge", CommitAfterMerge, "You have to finish your merge before commiting."},
		{"insufficient permission", MissingPerms, "You don't have permissions to write to the remote repository."},
		{"Another git process seems to be running in this repository", InterruptedOperation, "There's another Git process happening! Stop that process and try to commit again."},
		{"Authentication failed", InvalidRemoteCredentials, "Git couldn't recognize you, please, try login again using commita login"},
		{"refusing to merge unrelated histories", UnrelatedHists, "Your local and remote repositories don't have the same commit histories."},
		{"does not appear to be a git repository", NotARepo, ("The URL from your repository does not seems to be a Git Repository: " + conf.Url)},
		{"exceeds GitHub's file size limit", TooBig, "The files you want to upload are too big!"},
	}

	for _, p := range patterns {
		if strings.Contains(out, p.Raw) {
			return &CommitaErr{
				Raw:  p.Raw,
				Type: p.Type,
				Msg:  fmt.Sprintf("%s\n%s", aurora.BrightRed("fatal!"), p.Msg),
			}
		}
	}

	return &CommitaErr{"", Unknown, fmt.Sprintf("%s\nAn unexpected Git error happened: %s", aurora.BrightRed("fatal!"), out)}
}

func AttemptSolve(code int) error {
	return nil
}
