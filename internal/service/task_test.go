package service_test

import (
	"task-traker-cli/internal/model"
	"task-traker-cli/internal/service"
	"testing"
	"time"
)

func TestGetNextID(t *testing.T) {
	tasks := []model.Task{
		{TaskID: 1, Description: "model.Task 1", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{TaskID: 3, Description: "model.Task 2", Status: "done", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	next := service.GetNextID(tasks)
	if next != 4 {
		t.Errorf("expected next ID to be 4, got %d", next)
	}
}

func TestAddTask(t *testing.T) {
	tasks := []model.Task{}
	newDesc := "Write unit tests"
	newID := service.GetNextID(tasks)

	tasks = append(tasks, model.Task{
		TaskID:      newID,
		Description: newDesc,
		Status:      "todo",
	})

	if len(tasks) != 1 {
		t.Fatal("task was not added")
	}
	if tasks[0].Description != newDesc || tasks[0].Status != "todo" {
		t.Errorf("unexpected task data: %+v", tasks[0])
	}
}

func TestMarkStatus(t *testing.T) {
	tasks := []model.Task{
		{TaskID: 1, Description: "Learn Go", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	args := []string{"1"}

	service.MarkStatus(tasks, args, "in-progress")
	if tasks[0].Status != "in-progress" && tasks[0].UpdatedAt != time.Now() {
		t.Errorf("expected status to be 'in-progress', got '%s'", tasks[0].Status)
	}
}

func TestSaveAndLoadTasks(t *testing.T) {
	filename := "test_tasks.json"
	testTasks := []model.Task{
		{TaskID: 1, Description: "Test", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	err := service.SaveTasks(filename, testTasks)
	if err != nil {
		t.Fatal("failed to save tasks:", err)
	}

	loaded, err := service.LoadTasks(filename)
	if err != nil {
		t.Fatal("failed to load tasks:", err)
	}
	if len(loaded) != 1 || loaded[0].Description != "Test" {
		t.Error("loaded tasks do not match saved tasks")
	}
}
