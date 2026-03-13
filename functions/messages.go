package functions

import (
	"fmt"
)

// i'll change this because it'll be outdated after the migration
// but we can keep this now
var ShowHelp = fmt.Sprintf(
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

var ConfigHelp = fmt.Sprintf(
	"\nGenerates a config file if one doesn't exists\n\n" +
		"Comes with default values for:\n" +
		"\tdefaultBranch	 	'main'\n" +
		"\temojis	 			true\n" +
		"\tuseAI	 			false (TO BE IMPLEMENTED)")

var LoginHelp = fmt.Sprintf(
	"\nLogs you into your Git account\n" +
		"Commita will only ask for your username and email.")

var SetHelp = fmt.Sprintf(
	"\nSets commita's configurations\n" +
		"Simply use set and pass the config through the flag.",
)

func TestMsg() {
	fmt.Print("This is a test message")
}
