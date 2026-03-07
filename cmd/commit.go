package cmd

import (
	"fmt"
	"strings"

	"github.com/gabrielefagundes/commita/functions/exec"
	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/gabrielefagundes/commita/structs"
	"github.com/gabrielefagundes/commita/utils"
	"github.com/spf13/cobra"
)

var (
	commitType string
	feat       bool
	fix        bool
	docs       bool
	chore      bool
	refactor   bool
	remove     bool
)

var commitCmd = &cobra.Command{
	Use:           "commit <message>",
	Short:         "Commits your changes",
	Long:          `'commit' argument commits your changes to your remote repository, while also granting a semantical commit message.`,
	Args:          cobra.MaximumNArgs(2),
	Run:           runCommit,
	SilenceErrors: true,
}

func init() {
	commitCmd.Flags().BoolVar(&feat, "feat", false, "Adds the semantical tag 'feat'")
	commitCmd.Flags().BoolVar(&fix, "fix", false, "Adds the semantical tag 'fix'")
	commitCmd.Flags().BoolVar(&docs, "docs", false, "Adds the semantical tag 'docs'")
	commitCmd.Flags().BoolVar(&chore, "chore", false, "Adds the semantial tag 'chore'")
	commitCmd.Flags().BoolVar(&refactor, "refactor", false, "Adds the semantical tag 'refactor'")
	commitCmd.Flags().BoolVar(&remove, "remove", false, "Adds the semantical tag 'remove'")

	commitCmd.Flags().StringVar(&commitType, "type", "", "Defines the type of the semantical commit")

	rootCmd.AddCommand(commitCmd)
}

func runCommit(cmd *cobra.Command, args []string) {
	opts := structs.CommitOptions{}
	conf, _ := setup.LoadConfig()
	types := cmd.Flags().Changed("type")

	if len(args) == 0 {
		err := fmt.Sprint("Invalid\n" +
			"Please, provide a valid input\n" +
			"\t--type [commit type] [msg]\n" +
			"\t--[semantic] [msg]")

		fmt.Print(err)
		return
	}

	opts.Msg = strings.Join(args, " ")

	if feat {
		opts.CommitType = "feat"
	}

	if fix {
		opts.CommitType = "fix"
	}

	if docs {
		opts.CommitType = "docs"
	}

	if chore {
		opts.CommitType = "chore"
	}

	if refactor {
		opts.CommitType = "refactor"
	}

	if remove {
		opts.CommitType = "remove"
	}

	if types {
		opts.CommitType = commitType
	}

	opts.Emoji = conf.UseEmoji

	opts.Msg = utils.Parse(opts.CommitType, opts.Msg, opts.Emoji)

	fmt.Print(exec.CommitAndPush(opts.Msg))
}
