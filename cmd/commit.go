package cmd

import (
	functions "github.com/gabrielefagundes/commita/functions/flags"
	"github.com/spf13/cobra"
)

var (
	msg string
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commits your changes",
	Run: func(cmd *cobra.Command, args []string) {
		functions.CommitFeat(msg)
	},
}

func init() {
	commitCmd.Flags().StringVar(&msg, "feat", "", "Adds the semantical tag 'feat'")
	rootCmd.AddCommand(commitCmd)
}
