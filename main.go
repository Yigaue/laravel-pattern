// main
// Description: This program provides a command-line interface for Linux commands assistance.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yigaue/linux/commands"
)

func parseTLDRFile(path string) (*commands.TLDRPage, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var page commands.TLDRPage
	var currentExample commands.CommandEntry

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") {
			page.Title = strings.TrimPrefix(line, "#")
		} else if strings.HasPrefix(line, "> ") {
			page.Description = strings.TrimPrefix(line, "> ")
		} else if strings.HasPrefix(line, "- ") {
			
			if currentExample.Description != "" {
				page.Examples = append(page.Examples, currentExample)
				currentExample = commands.CommandEntry{}
			}
			currentExample.Description = strings.TrimPrefix(line, "- ")
		} else if strings.HasPrefix(line, "`") && strings.HasSuffix(line, "`") {
			currentExample.Command = strings.Trim(line, "`")
		}
	}

	if currentExample.Description != "" {
		page.Examples = append(page.Examples, currentExample)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &page, nil
}

func loadAllTLDRPages(dir string) (map[string]*commands.TLDRPage, error) {
	pages := make(map[string]*commands.TLDRPage)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			page, err := parseTLDRFile(path)

			if err != nil {
				return err
			}

			filename := strings.TrimSuffix(info.Name(), ".md")

			pages[strings.ToLower(filename)] = page
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return pages, nil
}

func main() {
	allPages, err := loadAllTLDRPages("commands/data/linux")
	if err != nil {
		fmt.Println("Error loading pages:", err)
		return
	}

	// --- Parse user input ---
	if len(os.Args) < 3 {
		fmt.Println("Usage: linux help <topic>")
		return
	}

	if strings.ToLower(os.Args[1]) != "help" {
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Usage: linux help <topic>")
		return
	}

	topic := strings.ToLower(os.Args[2])

	page, exists := allPages[topic]

	if !exists {
		fmt.Printf("No help found for: %s\n", topic)
		return
	}

	commands.PrintPage(page)
}