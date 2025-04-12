// commands is holds the logic to load commands from a JSON file and display help information
// for a specific command or a list of available commands.
package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CommandEntry struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type CommandData map[string][]CommandEntry

func loadCommandData() (CommandData, error) {
	path := filepath.Join("data", "commands.json")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data CommandData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	return data, err
}

func ShowHelp(topic string) {
	data, err := loadCommandData()
	if err != nil {
		fmt.Println("Failed to load command data:", err)
		return
	}

	if topic == "" {
		fmt.Println("Available topics:")
		for t := range data {
			fmt.Printf("  - %s\n", t)
		}
		fmt.Println("Example: linux help delete")
		return
	}

	topic = strings.ToLower(topic)
	commands, ok := data[topic]
	if !ok {
		fmt.Printf("No commands found for topic: '%s'\n", topic)
		return
	}

	fmt.Printf("Commands for '%s':\n", topic)
	for _, cmd := range commands {
		fmt.Printf("%-20s - %s\n", cmd.Command, cmd.Description)
	}
}
