# Task CLI

A simple command-line task manager written in Go.

## Installation

```bash
git clone https://github.com/your-username/task-cli.git
cd task-cli
go build -o task-cli .
```

## Usage

```bash
./task-cli <command> [arguments]
```

## Commands

| Command | Arguments | Description |
|---|---|---|
| `add` | `<description>` | Add a new task |
| `update` | `<id> <description>` | Update the description of an existing task |
| `delete` | `<id>` | Delete a task by ID |
| `in-progres` | `<id>` | Mark a task as **In Progress** |
| `done` | `<id>` | Mark a task as **Done** |
| `--help` | — | Show usage information |

## Examples

```bash
# Add a new task
./task-cli add Buy groceries

# Update a task
./task-cli update 1 Buy groceries and cook dinner

# Delete a task
./task-cli delete 1

# Mark a task as in progress
./task-cli in-progres 1

# Mark a task as done
./task-cli done 1

# Show help
./task-cli --help
```

## Task Statuses

- **Todo** — Default status when a task is created
- **In Progress** — Task is currently being worked on
- **Done** — Task has been completed

## Project Structure

```
task-cli/
├── main.go         # Entry point
├── cli.go          # CLI command handler (RunCLI)
├── data.go         # Data model and storage logic
└── README.md
```

## Requirements

- Go 1.18 or higher