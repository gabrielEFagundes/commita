package cmd

import (
	"fmt"

	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/spf13/cobra"
)

var (
	url    string
	branch string
	//// ai	bool
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets commita's configurations",
	Run:   changeConfigs,
}

func init() {
	configSetCmd.Flags().StringVarP(&url, "url", "u", "", "Adds the repository URL to your config file")
	configSetCmd.Flags().StringVarP(&branch, "branch", "b", "main", "Changes the default branch for commits without specified branch")
	configSetCmd.Flags().Bool("use-emoji", false, "Teels if commita should use emojis for predefined commit semantics")
	//* I won't implement AI flag yet

	configCmd.AddCommand(configSetCmd)
}

func changeConfigs(cmd *cobra.Command, args []string) {
	if err := setup.CreateConfig(); err != nil {
		fmt.Printf("\nerror creating config files: %v", err)
		return
	}

	conf, errConf := setup.LoadConfig()
	if errConf != nil {
		fmt.Printf("\nerror loading config files: %v", errConf)
		return
	}

	emoji, _ := cmd.Flags().GetBool("use-emoji")

	if cmd.Flags().Changed("url") {
		conf.Url = url
		fmt.Println("Remote repo set to " + url)
	}

	if url != "" && branch == "main" {
		conf.DefaultBranch = "main"
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

	setup.SaveConfig(*conf)
}
