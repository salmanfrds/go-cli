# Task CLI

A minimal command-line task manager written in Go. Add, list, complete, and delete tasks right from your terminal — tasks are saved to a `tasks.json` file so they persist between runs.

## Requirements

- [Go](https://go.dev/dl/) 1.26 or newer (matches `go.mod`)

Check your installed Go version:

```bash
go version
```

## Getting started

Clone the repo and run the app directly — no manual build step needed:

```bash
git clone <this-repo-url>
cd <this-repo-folder>
go run . list
```

`go run .` compiles and runs the app on the fly, so this works identically across macOS, Linux, and Windows.

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

## Building a Binary (OS-Specific Guide)

If you prefer executing a compiled binary instead of typing `go run .` every time, follow the instructions for your operating system:

###  macOS & 🐧 Linux (Terminal / Bash / Zsh)

1. Build the executable binary:
   ```bash
   go build -o task .
   ```
2. Run the binary:
   ```bash
   ./task list
   ./task add "New task"
   ./task done 1
   ```

### 🪟 Windows (Command Prompt / PowerShell)

1. Build the `.exe` executable binary:
   - **PowerShell / CMD**:
     ```powershell
     go build -o task.exe .
     ```
2. Run the binary:
   - **PowerShell**:
     ```powershell
     .\task.exe list
     .\task.exe add "New task"
     .\task.exe done 1
     ```
   - **Command Prompt (CMD)**:
     ```cmd
     task.exe list
     task.exe add "New task"
     task.exe done 1
     ```

---

## Cross-Compiling for Other Operating Systems

Go makes it effortless to build binaries for other operating systems from any platform:

* **Build for Windows (from macOS/Linux):**
  ```bash
  GOOS=windows GOARCH=amd64 go build -o task.exe .
  ```

* **Build for Linux (from macOS/Windows):**
  - macOS/Linux (Bash/Zsh):
    ```bash
    GOOS=linux GOARCH=amd64 go build -o task .
    ```
  - Windows (PowerShell):
    ```powershell
    $env:GOOS="linux"; $env:GOARCH="amd64"; go build -o task .
    ```

* **Build for macOS (from Windows/Linux):**
  - Windows (PowerShell):
    ```powershell
    $env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o task .
    ```

---

## How It Works

- **Data Model**: Tasks are stored as plain Go structs (`Task{ID, Title, Done}`) in memory while the program runs.
- **Persistence**: On every run, tasks are loaded from `tasks.json` (if present) and saved back automatically after any modification.
- **Zero Third-Party Dependencies**: Built entirely using Go's standard library (`encoding/json`, `os`, `fmt`, `strconv`).
