# linux Command Helper - Features

This file documents the current and planned features of the Linux Command Helper CLI tool.

---

## Core Features (Implemented)

- **Command Lookup from JSON File**
  - Load commands and their descriptions from `commands.json`.

- **Help by Topic**
  - Usage: `linux help <topic>`
  - Example: `linux help delete`

- **Fuzzy Search for Topics**
  - Types an approximate or partial topic, best-matching suggestions are shown.

- **Friendly Error Messages**
  - For incorrect or incomplete commands, the tool shows helpful usage tips.

---

## Planned Features

### 1. Command Auto-Suggest While Typing

- Interactive mode or shell integration that suggests commands in real-time.
- Command suggestion based on partial input.
- Optional shortcut to copy a command to clipboard.

### 2. Inline Descriptions as Comments

- `--explain` flag appends a description to each command shown.
- Example: `linux help delete --explain`

### 3. Search by Keyword

- Usage: `linux search <term>`
- Matches both command topics and command descriptions.

### 4. Dynamic JSON File Reload

- Automatically reloads updated command list from JSON file without restarting the tool.

### 5. User-defined Command Snippets

- Allow users to add their own commands with:
  - `linux add "topic" "command" "description"`

### 6. Color-coded Terminal Output

- Different colors for command names, parameters, and descriptions.
- Highlight matched search terms.

### 7. Interactive Mode (REPL-like)

- Start with `linux` alone to enter a session:
  - `> help create`
  - `> search network`
  - `> exit`

### 8. Test Mode / Self-Quiz

- Randomly quizzes user on command topics.
- Usage: `linux test`

### 9. AI Explain Mode *(Optional)*

- Usage: `linux explain "<command>"`
- Provides short explanation of any command.
- Optional integration with local ruleset or GPT API.

### 10. Export to Markdown / HTML

- Allow users to generate a cheat sheet or documentation in markdown or HTML.
- Usage: `linux export [markdown|html]`

---

## Contributions

Have an idea or a command to contribute? Submit a pull request with your changes or propose a new feature(using issue on Github)!