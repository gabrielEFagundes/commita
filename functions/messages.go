package functions

import (
	"fmt"
)

var ShowHelp = fmt.Sprintf(
	"\nCommita - v2.0.0-beta" +
		"\nWelcome to Commita, now powered by GOlang!" +
		"\n\nYou shall start by setupping your config file for your project:\n" +
		"	commita config set -u/--url <<your remote URL>> -b/--branch <<your default branch>>\n" +
		"	commita config login -l/--local/-g/--global <<username>> <<email>>\n" +
		"Usage:\n\n" +
		"   commit\n" +
		"		[flags] --feat		<<message>>\n" +
		"				--fix		<<message>>\n" +
		"				--docs		<<message>>\n" +
		"				--chore		<<message>>\n" +
		"				--refactor 	<<message>>\n" +
		"				--remove 	<<message>>\n" +
		"				--type 		<<commit type>> <<message>>\n\n")

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
		"Simply use set and pass the config through a flag.")
