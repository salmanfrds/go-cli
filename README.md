# Task CLI

A minimal command-line task manager written in Go. Add, list, complete, and delete tasks right from your terminal — tasks are saved to a `tasks.json` file so they persist between runs.

## Requirements

- [Go](https://go.dev/dl/) 1.26 or newer (matches `go.mod`)

Check your version:

```bash
go version
```

## Getting Started & Usage

Clone the repo and run the app directly — works on **macOS**, **Linux**, and **Windows**:

```bash
git clone <this-repo-url>
cd <this-repo-folder>
```

### Run directly with `go run .`

```bash
go run . add "Buy milk"     # Add a new task
go run . list               # List all tasks
go run . done 1             # Mark task #1 as complete
go run . delete 1           # Delete task #1
```

Running `go run .` with no command prints usage help.

---

## Building a Binary (Optional)

If you prefer building an executable binary to run directly on your device:

### macOS / Linux
```bash
go build -o task .
./task list
```

### Windows (PowerShell / Command Prompt)
```powershell
go build -o task.exe .
.\task.exe list
```

---

## How It Works

- **Data Model**: Tasks are stored as plain Go structs (`Task{ID, Title, Done}`) in memory while the program runs.
- **Persistence**: On every run, tasks are loaded from `tasks.json` (if present) and saved back automatically after any modification.
- **Zero Dependencies**: Built entirely using Go's standard library (`encoding/json`, `os`, `fmt`, `strconv`).
