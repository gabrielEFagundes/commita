package cmd

import (
	"fmt"
	"os"

	"github.com/gabrielefagundes/commita/functions"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commita",
	Short: "A Git Framework",
	Long:  functions.ShowHelp,
}

func init() {
	rootCmd.Version = "v2.0.0-beta"
	rootCmd.SetVersionTemplate("Commita - {{.Version}}")
}

func Exec() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}
}
