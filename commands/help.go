// commands is holds the logic to load commands from a JSON file and display help information
// for a specific command or a list of available commands.
package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"

	"github.com/lithammer/fuzzysearch/fuzzy"
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

	// If no topic is given, list all available topics
	if topic == "" {
		title := color.New(color.FgGreen, color.Bold).SprintFunc()
		topicColor := color.New(color.FgYellow).SprintFunc()

		fmt.Println(title("\n Available Topics:\n"))
		for t := range data {
			fmt.Printf("  • %s\n", topicColor(t))
		}
		return
	}

	topic = strings.ToLower(topic)

	// Exact match
	if commands, ok := data[topic]; ok {
		printCommands(topic, commands)
		return
	}

	// Fuzzy match
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	matches := fuzzy.Find(topic, keys)

	if len(matches) > 0 {
		bestMatch := matches[0]
		fmt.Printf("Did you mean '%s'?\n\n", bestMatch)
		printCommands(bestMatch, data[bestMatch])
		return
	}

	// No matches — suggest closest
	warn := color.New(color.FgRed, color.Bold).SprintFunc()
	fmt.Printf("%s Unknown topic: '%s'\n", warn(""), topic)
	fmt.Println("Use 'linux help' to see all available topics.")

	suggestions := fuzzy.RankFindNormalizedFold(topic, keys)
	for i, s := range suggestions {
		if i >= 3 {
			break
		}
		fmt.Printf("  - %s\n", s.Target)
	}
}

func printCommands(topic string, commands []CommandEntry) {
	header := color.New(color.FgGreen, color.Bold).SprintFunc()
	cmdColor := color.New(color.FgCyan).SprintFunc()
	descColor := color.New(color.FgWhite).SprintFunc()

	fmt.Printf("\n%s %s\n\n", header("Topic:"), header(strings.Title(topic)))

	for _, cmd := range commands {
		fmt.Printf("  %s  %s\n", cmdColor(cmd.Command), descColor("- "+cmd.Description))
	}

	fmt.Println()
}
