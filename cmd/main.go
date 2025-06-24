package main

import (
	"fmt"
	"os"
	"strconv"
	"task-traker-cli/internal/model"
	"task-traker-cli/internal/service"
	"time"
)

const fileName = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
	}

	command := os.Args[1]
	args := os.Args[2:]

	tasks, err := service.LoadTasks(fileName)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	switch command {
	case "add":
		if len(args) < 1 {
			fmt.Println("Usage: task-cli add \"Task title\"")
			return
		}
		description := args[0]
		id := service.GetNextID(tasks)
		tasks = append(tasks, model.Task{TaskID: id, Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()})
		fmt.Printf("Task added successfully (ID: %d)\n", id)

	case "update":
		if len(args) < 2 {
			fmt.Println("Usage: task-cli update <id> \"New description\"")
			return
		}
		id, _ := strconv.Atoi(args[0])
		found := false
		for i := range tasks {
			if tasks[i].TaskID == id {
				tasks[i].Description = args[1]
				tasks[i].UpdatedAt = time.Now()
				found = true
				fmt.Println("Task updated successfully")
				break
			}
		}
		if !found {
			fmt.Println("Task not found")
		}

	case "delete":
		if len(args) < 1 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(args[0])
		newTasks := []model.Task{}
		found := false
		for _, task := range tasks {
			if task.TaskID != id {
				newTasks = append(newTasks, task)
			} else {
				found = true
			}
		}
		if found {
			tasks = newTasks
			fmt.Println("Task delete successfully")
		} else {
			fmt.Println("Task not found")
		}

	case "mark-in-progress":
		service.MarkStatus(tasks, args, "in-progress")

	case "mark-done":
		service.MarkStatus(tasks, args, "done")

	case "list":
		if len(args) == 0 {
			service.ListTasks(tasks, "")
		} else {
			service.ListTasks(tasks, args[0])
		}

	default:
		fmt.Println("Unknown command: ", command)
	}

	if err := service.SaveTasks(fileName, tasks); err != nil {
		fmt.Println("Error saving tasks: ", err)
	}

}
