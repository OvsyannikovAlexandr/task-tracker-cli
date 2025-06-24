package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"task-traker-cli/internal/model"
	"time"
)

func LoadTasks(fileName string) ([]model.Task, error) {
	data, err := os.ReadFile(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return []model.Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	var tasks []model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(fileName string, tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func GetNextID(tasks []model.Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.TaskID > maxID {
			maxID = task.TaskID
		}
	}
	return maxID + 1
}

func MarkStatus(tasks []model.Task, args []string, status string) {
	if len(args) < 1 {
		fmt.Printf("Usage: task-cli mark-%s <id>\n", status)
		return
	}
	id, _ := strconv.Atoi(args[0])
	for i := range tasks {
		if tasks[i].TaskID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			fmt.Printf("Task marked as %s\n", status)
			return
		}
	}
	fmt.Println("Task not found")
}

func ListTasks(tasks []model.Task, filter string) {
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("[%d] Desc: \"%s\" | Status: (%s) | CreatedAt: %s; UpdatedAt: %s\n",
				task.TaskID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}
}
