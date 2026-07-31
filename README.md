# Task CLI

A minimal command-line task manager written in Go. Add, list, complete, and delete tasks
right from your terminal — tasks are saved to a `tasks.json` file so they persist between
runs.

## Requirements

- [Go](https://go.dev/dl/) 1.26 or newer (matches `go.mod`)

Check your version:

```bash
go version
```

## Getting started

Clone the repo and run the app directly — no build step needed:

```bash
git clone <this-repo-url>
cd <this-repo-folder>
go run . list
```

`go run .` compiles and runs the app in one step, so this works the same way on macOS,
Linux, and Windows.

## Usage

```bash
go run . add "<task title>"     # add a new task
go run . list                   # list all tasks
go run . done <id>              # mark a task as complete
go run . delete <id>            # remove a task
```

`<id>` is the number shown next to each task in `go run . list`.

### Example

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

Running with no command (or an unknown one) prints usage help:

```bash
go run .
```

## How it works

- Tasks are stored as plain structs (`Task{ID, Title, Done}`) in memory while the program runs.
- On every run, tasks are loaded from `tasks.json` (if it exists) and saved back after any change.
- No external dependencies — the entire app uses only Go's standard library (`encoding/json`, `os`, `fmt`, `strconv`).

## Building a binary (optional)

If you'd rather not type `go run .` every time:

```bash
go build -o task .
./task list
```

Note: a binary built on one OS/architecture won't run on another — rebuild with `go build`
on each machine, or use `go run .` instead.
