package main

import (
	"fmt"
	"os"

	"github.com/yigaue/linux-commands-assistance/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: linux <keyword>")
		return
	}

	switch os.Args[1] {
	case "help":
		if len(os.Args) > 2 {
			commands.ShowHelp(os.Args[2])
		} else {
			commands.ShowHelp("")
		}
	default:
		fmt.Println("Usage. Try: linux <topics>")
	}
}
