# task-tracker-cli
A simple command-line task tracker written in **Go**.  
It allows users to add, update, delete tasks, change task status, and list tasks using only native Go libraries and positional CLI arguments.
Tasks are stored in a local `tasks.json` file in the current directory.

Roadmap Project Challenge: https://roadmap.sh/projects/task-tracker

## ✅ Features

- Add, update, and delete tasks  
- Mark tasks as **todo**, **in-progress**, or **done**  
- List all tasks or filter by status  
- Automatically creates `tasks.json` if it doesn’t exist  
- Built without external libraries  
- Graceful error handling
## ⚙️ Installation

1. **Clone the repository**:

```bash
git clone https://github.com/yourusername/task-tracker-cli.git
cd task-tracker-cli
```

2. **Build the binary**:

```bash
go build -o task-cli
```

3. **Run the app**:

```bash
./task-cli <command> [arguments]
```
or

```bash
go run ./cmd/main.go <command> [arguments]
```

## 📘 Usage Guide

### ➕ Add a task

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### ✏️ Update a task

```bash
task-cli update <id> "New task title"
task-cli update 1 "Buy groceries and cook dinner"
```

### ❌ Delete a task

```bash
task-cli delete <id>
task-cli delete 1
```

### 🚧 Mark as in progress

```bash
task-cli mark-in-progress <id>
```

### ✅ Mark as done

```bash
task-cli mark-done <id>
```

### 📋 List all tasks

```bash
task-cli list
```

### 📂 List tasks by status

```bash
task-cli list todo
task-cli list in-progress
task-cli list done
```

## 🖥️ Example Output

```bash
$ task-cli list

[1] Desc: "SomeTask1" | Status: (todo) | CreatedAt: 2025-06-24 12:50:08; UpdatedAt: 2025-06-24 12:50:08
[2] Desc: "SomeTask2" | Status: (todo) | CreatedAt: 2025-06-24 13:48:31; UpdatedAt: 2025-06-24 13:48:52
[3] Desc: "SomeTask3" | Status: (todo) | CreatedAt: 2025-06-24 13:48:36; UpdatedAt: 2025-06-24 13:48:36
```

## 💾 Task Storage Format (`tasks.json`)

All tasks are stored in JSON format:

```json
[
 {
  "task_id": 1,
  "description": "SomeTask1",
  "status": "todo",
  "created_at": "2025-06-24T12:50:08.8991133+03:00",
  "updated_at": "2025-06-24T12:50:08.8991133+03:00"
 },
 {
  "task_id": 2,
  "description": "SomeTask2",
  "status": "todo",
  "created_at": "2025-06-24T13:48:31.0745442+03:00",
  "updated_at": "2025-06-24T13:48:52.6237612+03:00"
 },
 {
  "task_id": 3,
  "description": "SomeTask3",
  "status": "todo",
  "created_at": "2025-06-24T13:48:36.6884142+03:00",
  "updated_at": "2025-06-24T13:48:36.6884142+03:00"
 }
]
```

## 🧰 Requirements

- Go 1.16 or higher

## 🚀 Potential Improvements

- Add support for due dates and priorities  
- Export/import tasks  
- Add search or filter by keyword  
- Add colored output by status  
- Unit tests  

## 📄 License

This project is open-source and available under the (**MIT License**)[https://github.com/abneed/task-tracker-cli/blob/master/LICENSE].

