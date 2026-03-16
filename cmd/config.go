package cmd

import (
	"github.com/gabrielefagundes/commita/functions"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generates (if not exists) a config file",
	Long:  functions.ConfigHelp,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
