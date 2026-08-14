# Task CLI

A minimal command-line task manager written in Go. Add, list, complete, and delete tasks right from your terminal — tasks are saved to a `tasks.json` file so they persist between runs.

## Requirements

- [Go](https://go.dev/dl/) 1.26 or newer (matches `go.mod`)

Check your installed Go version:

```bash
go version
```

## Getting started

Clone the repo and run the app directly — **no build step required**:

```bash
git clone <this-repo-url>
cd <this-repo-folder>
go run . list
```

> 💡 `go run .` automatically detects your operating system and works identically on **macOS**, **Linux**, and **Windows**.

---

## Usage

Run commands with `go run .`:

```bash
go run . add "<task title>"     # Add a new task
go run . list                   # List all tasks
go run . done <id>              # Mark a task as complete
go run . delete <id>            # Remove a task
```

`<id>` is the number shown next to each task in `go run . list`.

### Example Walkthrough

```bash
go run . add "Buy milk"
go run . add "Finish homework"
go run . list
# [ ] 1. Buy milk
# [ ] 2. Finish homework

go run . done 1
go run . list
# [x] 1. Buy milk
# [ ] 2. Finish homework

go run . delete 2
go run . list
# [x] 1. Buy milk
```

Running with no command (or an unknown command) prints usage help:

```bash
go run .
```

---

## Building a Binary (Optional)

If you prefer running a pre-compiled binary instead of `go run .`:

###  macOS & 🐧 Linux
```bash
go build -o task .
./task list
```

### 🪟 Windows (PowerShell / Command Prompt)
```powershell
go build -o task.exe .
.\task.exe list
```

> **Note:** Standard `go build` automatically detects your current operating system and builds the correct binary for your machine. You do **not** need to set any extra environment variables!

---

## Optional: Cross-Compiling for Another OS

Go allows you to build binaries for *other* operating systems (e.g. building a Windows `.exe` while working on macOS):

```bash
# Build for Windows from macOS/Linux:
GOOS=windows GOARCH=amd64 go build -o task.exe .

# Build for Linux from macOS/Windows:
GOOS=linux GOARCH=amd64 go build -o task .
```

---

## How It Works

- **Data Model**: Tasks are stored as plain Go structs (`Task{ID, Title, Done}`) in memory while the program runs.
- **Persistence**: On every run, tasks are loaded from `tasks.json` (if present) and saved back automatically after any modification.
- **Zero Third-Party Dependencies**: Built entirely using Go's standard library (`encoding/json`, `os`, `fmt`, `strconv`).
