// WORKSHOP GUIDE: follow the STEP 1-9 comments in order, starting with the struct below.
package main

import (
	// "encoding/json"
	"fmt"
	// "os"
	// "strconv"
)

// STEP 1 — the Task model: every task is just an ID, a title, and a done flag.
// TODO: add ID (int), Title (string), and Done (bool) fields with json tags.
type Task struct {

}

// STEP 2: entry point — loads tasks, routes the subcommand, then saves.
// TODO: load tasks, check os.Args, switch on os.Args[1] for "add"/"list"/"done"/"delete",
// call the matching helper below, then save tasks before returning.
func main() {

}

// STEP 3 — printed when no command (or an unrecognized one) is given.
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run . add \"task title\"")
	fmt.Println("  go run . list")
	fmt.Println("  go run . done <id>")
	fmt.Println("  go run . delete <id>")
}

// STEP 4 — PERSISTENCE: loadTasks reads tasks.json.
// TODO: read tasks.json, unmarshal into []Task (return empty slice on error).
func loadTasks() []Task {
	return []Task{}
}

// STEP 5 — READ: prints every task with a checkbox-style status marker.
// TODO: loop over tasks and print each one's status, ID, and title.
func listTasks(tasks []Task) {

}

// STEP 6 — CREATE: builds a new Task and appends it to the slice.
// TODO: construct a Task from title, append it, print confirmation, return tasks.
func addTask(tasks []Task, title string) []Task {
	return tasks
}

// STEP 7 — PERSISTENCE: saveTasks writes the tasks slice to tasks.json.
// TODO: marshal tasks to indented JSON and write it to tasks.json with 0644 permissions.
func saveTasks(tasks []Task) {

}

// STEP 8 — UPDATE: finds the task by ID and flips Done to true.
// TODO: find the task with matching ID, set Done = true, print confirmation.
func completeTask(tasks []Task, id int) []Task {
	return tasks
}

// STEP 9 — DELETE: removes the task at index i by slicing around it.
// TODO: find the task with matching ID and remove it from the slice.
func deleteTask(tasks []Task, id int) []Task {
	return tasks
}
