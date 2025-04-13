# Linux Command Helper

A lightweight CLI tool for quickly looking up commonly used Linux terminal commands grouped by topic. Perfect for those who forget syntax often and want instant, offline help.

---

## Features

- **Fuzzy Search**: Get suggestions even with partial or misspelled topics.
- **Help by Topic**: Run `linux help <topic>` to see relevant commands.
- **Smart Suggestions**: Interactive mode to explore and discover commands.
- **Keyword Search**: Look up any word in command descriptions.
- **Extensible**: Add your own command snippets.
- **Pretty Output**: Color-coded and readable CLI.
- **Offline-first**: No network calls—just pure local JSON.

See [`FEATURES.md`](./FEATURES.md) for upcoming features.

---

## Installation

```bash
# Clone the repository
git clone https://github.com/yigaue/linux-commands-assistance.git

# Navigate into the folder
cd linux-command-helper

# Build the binary
go build -o linux
```

---

## 🧪 Usage

```bash
# Basic command
./linux help create

# With fuzzy matching
./linux help cre

# Invalid usage hint
./linux help
> Usage: linux help <topic>
```

---

## JSON Format (commands.json)

Commands are defined in a simple, editable JSON format:

```json
{
  "create": [
    { "command": "mkdir <dir>", "description": "Create a directory" },
    { "command": "touch <file>", "description": "Create a file" }
  ],
  "delete": [
    { "command": "rm <file>", "description": "Delete a file" },
    { "command": "rm -r <dir>", "description": "Delete a directory recursively" }
  ]
}
```

---

## Dev Guide

```bash
# Format and lint
go fmt ./...

# Run with test args
go run main.go help move
```

## Contributing

PRs welcome! Feel free to submit commands, improve the tool, or open an issue.

---

## [License](LICENCE)

MIT © [@yigaue](https://github.com/yigaue)

---

> Built with Go & ❤️ for the forgetful terminal warrior.