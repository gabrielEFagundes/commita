package main

import (
	"fmt"
	"os"

	"github.com/jonathonwebb/getopt"
)

var state = getopt.NewState(os.Args)
var config = getopt.Config{
	Opts:     getopt.OptStr(`hv`),
	LongOpts: getopt.LongOptStr("help,version"),
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
	fmt.Print("\nCommita - v1.2.2\n")
}

func wrong_opt() {
	fmt.Print("\nUnknown flag parsed, the available options are:\n\n" +
		"		-v		Version\n" +
		"		-h		Help\n\n")
}

func main() {
	// this prints a lot of wrong messages when using the long flag
	// it must be the loop iterating over every single character
	// solve this ;)
	state.Parse(config)
	for opt, err := range state.All(config) {
		if err != nil {
			wrong_opt()
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
