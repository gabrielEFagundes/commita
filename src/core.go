package main

import (
	"fmt"
	"os"

	"github.com/jonathonwebb/getopt"
)

var state = getopt.NewState(os.Args)
var config = getopt.Config{
	Opts: getopt.OptStr(`hv`),
}

func show_help() {
	buffer := fmt.Sprintf(
		"\nCommita - v1.2.2" +
			"\nWelcome to Commita!\n\n" +
			"All commands start with 'git'\n\n" +
			"   feat        :sparkles: feat: <<message>>\n" +
			"   fix         :bug: fix: <<message>>\n" +
			"   chore       :art: chore: <<message>>\n" +
			"   docs        :memo: docs: <<message>>\n" +
			"   refactor    :recycle: refactor: <<message>>\n" +
			"   remove      :fire: rm: <<message>>\n\n" +
			"The syntax for most commands are:\n\n" +
			"git [prefix] -u <<repo's url>> -b <<branch>> <<message>>\n\n" +
			"   prefix      the type of commit message you want\n" +
			"   -b          (optional) your desired branch, uses the default 'main' branch if not parsed\n" +
			"   -u          (optional) your repo's url, only parse on your first commit or if not parsed yet\n" +
			"   message     your commit message\n\n" +
			"You can also login into your git account way faster:\n\n" +
			"git login [-l || -g] username email\n\n" +
			"   -l          same as --local\n" +
			"   -g          same as --global\n\n")

	fmt.Print(buffer)
}

func show_version() {
	fmt.Print("Commita - v1.2.2")
}

func wrong_opt() {
	fmt.Print("\nUnknown flag parsed, the available options are:\n\n" +
		"		-v		Version\n" +
		"		-h		Help\n\n")
}

func main() {
	// I'll leave it like this and move on with the migration, we'll look forward to this later
	state.Parse(config)

	if len(os.Args) < 2 {
		show_help()
		os.Exit(0)
	}

	for {
		opt, err := state.GetOpt(config)

		if err != nil {
			wrong_opt()
			fmt.Print("Error:", err)
			os.Exit(0)
		}

		switch {
		case opt.Char == 'h' || opt.Name == "help":
			show_help()
			return

		case opt.Char == 'v' || opt.Name == "version":
			show_version()
			return
		}
	}
}
