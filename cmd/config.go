package cmd

import (
	"fmt"

	"github.com/gabrielefagundes/commita/functions"
	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/spf13/cobra"
)

var (
	url    string
	branch string
	// ai	bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generates (if not exists) a config file",
	Long:  functions.ConfigHelp(),
	Run:   changeConfigs,
}

func init() {
	configCmd.Flags().StringVarP(&url, "url", "u", "", "Adds the repository URL to your config file")
	configCmd.Flags().StringVarP(&branch, "branch", "b", "main", "Changes the default branch for commits without specified branch")
	configCmd.Flags().Bool("use-emoji", false, "Teels if commita should use emojis for predefined commit semantics") // maybe set this to true
	// I won't implement AI flag yet
	rootCmd.AddCommand(configCmd)
}

func changeConfigs(cmd *cobra.Command, args []string) {
	setup.CreateConfig()

	conf, _ := setup.LoadConfig()

	emoji, _ := cmd.Flags().GetBool("use-emoji")

	if cmd.Flags().Changed("url") {
		conf.Url = url
		fmt.Println("Remote repo set to " + url)
	}

	if url != "" && branch == "main" {
		fmt.Println("Remote repo updated and no branch specified\nBranch set to 'main'")
	}

	if cmd.Flags().Changed("branch") {
		conf.DefaultBranch = branch
		fmt.Println("Branch set to '" + branch + "'")
	}

	if emoji {
		conf.UseEmoji = !conf.UseEmoji
		fmt.Printf("Emojis set to %t", conf.UseEmoji)
	}

	setup.SaveConfig(conf)
}
