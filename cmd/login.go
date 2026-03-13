package cmd

import (
	"fmt"
	"os/exec"

	"github.com/gabrielefagundes/commita/functions"
	"github.com/gabrielefagundes/commita/functions/setup"
	"github.com/gabrielefagundes/commita/utils"
	"github.com/spf13/cobra"
)

var (
	local  bool
	global bool
)

// ! this still has no effect over commita's process to commit, but I'll it implement later.
var configLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into your Git account",
	Args:  cobra.MinimumNArgs(2),
	Long:  functions.LoginHelp,
	RunE:  login,
}

func init() {
	configLoginCmd.Flags().BoolVarP(&local, "local", "l", false, "Locally applies your login to Git")
	configLoginCmd.Flags().BoolVarP(&global, "global", "g", false, "Globally applies your login to Git")

	configCmd.AddCommand(configLoginCmd)
}

func login(cmd *cobra.Command, args []string) error {
	if err := setup.CreateConfig(); err != nil {
		return fmt.Errorf("\nerror creating config files: %v", err)
	}

	conf, _ := setup.LoadConfig()

	if local && global {
		return fmt.Errorf("\nplease choose between '--local' or '--global'")
	}

	email, user := utils.ValidateEmail(args[0], args[1])

	scope := "--local"
	if global {
		scope = "--global"
	}

	exec.Command("git", "config", scope, "user.name", user).Run()
	exec.Command("git", "config", scope, "user.email", email).Run()

	conf.Login.Abundance = scope
	conf.Login.Email = email
	conf.Login.User = user

	return setup.SaveConfig(*conf)
}
