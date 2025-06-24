package service_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"task-traker-cli/internal/model"
	"task-traker-cli/internal/service"
	"testing"
	"time"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old

	return buf.String()
}

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

func TestListTasks(t *testing.T) {
	tasks := []model.Task{
		{TaskID: 1, Description: "Task1", Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{TaskID: 2, Description: "Task2", Status: "in-progress", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{TaskID: 3, Description: "Task3", Status: "done", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	t.Run("list all tasks", func(t *testing.T) {
		output := captureOutput(func() {
			service.ListTasks(tasks, "")
		})
		if !strings.Contains(output, "Task1") ||
			!strings.Contains(output, "Task2") ||
			!strings.Contains(output, "Task3") {
			t.Error("listTasks should output all tasks")
		}
	})

	t.Run("filter todo", func(t *testing.T) {
		output := captureOutput(func() {
			service.ListTasks(tasks, "todo")
		})
		if !strings.Contains(output, "Task1") ||
			strings.Contains(output, "Task2") ||
			strings.Contains(output, "Task3") {
			t.Error("listTasks should output only 'todo' tasks")
		}
	})

	t.Run("filter in-progress", func(t *testing.T) {
		output := captureOutput(func() {
			service.ListTasks(tasks, "in-progress")
		})
		if !strings.Contains(output, "Task2") ||
			strings.Contains(output, "Task1") ||
			strings.Contains(output, "Task3") {
			t.Error("listTasks should output only 'in-progress' tasks")
		}
	})

	t.Run("filter done", func(t *testing.T) {
		output := captureOutput(func() {
			service.ListTasks(tasks, "done")
		})
		if !strings.Contains(output, "Task3") ||
			strings.Contains(output, "Task2") ||
			strings.Contains(output, "Task1") {
			t.Error("listTasks should output only 'done' tasks")
		}
	})
}
