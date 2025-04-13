package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func PrintAbout() {
	about := color.New(color.FgMagenta, color.Bold).SprintFunc()
	info := color.New(color.FgYellow).SprintFunc()

	// Get the version from the VERSION file
	version := getVersion()

	fmt.Printf("%s\n", about("🔧 Linux Command Help CLI Tool"))
	fmt.Println(info("A tool for quickly finding Linux commands with easy-to-read descriptions"))
	fmt.Println()
	fmt.Printf("%s\n", about("Usage:"))
	fmt.Println(info("  linux help <topic> - Get help on a specific topic"))
	fmt.Println(info("  linux help          - List all available topics"))
	fmt.Println()
	fmt.Printf("%s\n", about("Examples:"))
	fmt.Println(info("  linux help create   - Shows commands related to creating files/directories"))
	fmt.Println(info("  linux help delete   - Shows commands for deleting files/directories"))
	fmt.Println(info("  linux help move     - Shows commands for moving files/directories"))
	fmt.Println()
	fmt.Printf("%s\n", about("Version:"))
	fmt.Println(info("  " + version))
}

func getVersion() string {
	data, err := os.ReadFile("VERSION")
	if err != nil {
		fmt.Println("Error reading version file:", err)
		return "Unknown"
	}
	return strings.TrimSpace(string(data))
}