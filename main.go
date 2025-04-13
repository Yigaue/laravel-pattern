// main
// Description: This program provides a command-line interface for Linux commands assistance.
package main

import (
	"fmt"
	"os"

	"github.com/yigaue/linux-commands-assistance/app"
	"github.com/yigaue/linux-commands-assistance/commands"
)

func main() {
	const EXPECTED_HELP_ARGUMENTS_COUNT = 2;
	const EXPECTED_ARGUMENTS_INDEX = 2;
	const EXPECTED_HELP_INDEX = 1;

	if len(os.Args) < 2 {
		app.PrintAbout()
		return
	}

	if len(os.Args) == EXPECTED_HELP_ARGUMENTS_COUNT {
		commands.ShowHelp("")
		fmt.Printf("Run 'linux help <topic>' for more information on a specific topic.\n")
		return
	} 
	commands.ShowHelp(os.Args[EXPECTED_ARGUMENTS_INDEX])
}

func init() {
	if _, err := os.Stat("data/commands.json"); os.IsNotExist(err) {
		fmt.Println("Error: commands.json file not found in the data directory.")
		os.Exit(1)
	}
}