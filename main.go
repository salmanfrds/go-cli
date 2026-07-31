// WORKSHOP GUIDE: follow the STEP 1-7 comments in order, starting in main() below.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const fileName = "tasks.json"

// STEP 1: entry point — loads tasks, routes the subcommand, then saves.
func main() {
	tasks := loadTasks()

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {

	// Add Command
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . add \"task title\"")
			return
		}
		tasks = addTask(tasks, os.Args[2])

	// List Command
	case "list":
		listTasks(tasks)

	// Mark as Done Command
	case "done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		tasks = completeTask(tasks, id)

	// Delete Command
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		tasks = deleteTask(tasks, id)

	default:
		printUsage()
		return
	}

	saveTasks(tasks)
}

// STEP 2 — CREATE: builds a new Task and appends it to the slice.
func addTask(tasks []Task, title string) []Task {
	task := Task{ID: len(tasks) + 1, Title: title, Done: false}
	tasks = append(tasks, task)
	fmt.Println("Added:", title)
	return tasks
}

// STEP 3 — READ: prints every task with a checkbox-style status marker.
func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Title)
	}
}

// STEP 4 — UPDATE: finds the task by ID and flips Done to true.
func completeTask(tasks []Task, id int) []Task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			fmt.Println("Completed:", tasks[i].Title)
			return tasks
		}
	}
	fmt.Println("Task not found")
	return tasks
}

// STEP 5 — DELETE: removes the task at index i by slicing around it.
func deleteTask(tasks []Task, id int) []Task {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Deleted:", t.Title)
			return tasks
		}
	}
	fmt.Println("Task not found")
	return tasks
}

// STEP 6 — PERSISTENCE: loadTasks reads tasks.json, saveTasks writes it back.
func loadTasks() []Task {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

// STEP 7 — printed when no command (or an unrecognized one) is given.
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run . add \"task title\"")
	fmt.Println("  go run . list")
	fmt.Println("  go run . done <id>")
	fmt.Println("  go run . delete <id>")
}
