// WORKSHOP GUIDE: follow the STEP 1-9 comments in order, starting with the struct below.
package main

import (
	// "encoding/json"
	"fmt"
	// "os"
	// "strconv"
)

// STEP 1 — the Task model: every task is just an ID, a title, and a done flag.
type Task struct {

}

// STEP 2: entry point — loads tasks, routes the subcommand, then saves.
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
func loadTasks() []Task {
	return []Task{}
}

// STEP 5 — READ: prints every task with a checkbox-style status marker.
func listTasks(tasks []Task) {

}

// STEP 6 — CREATE: builds a new Task and appends it to the slice.
func addTask(tasks []Task, title string) []Task {
	return tasks
}

// STEP 7 — PERSISTENCE: saveTasks writes the tasks slice to tasks.json.
func saveTasks(tasks []Task) {

}

// STEP 8 — UPDATE: finds the task by ID and flips Done to true.
func completeTask(tasks []Task, id int) []Task {
	return tasks
}

// STEP 9 — DELETE: removes the task at index i by slicing around it.
func deleteTask(tasks []Task, id int) []Task {
	return tasks
}
