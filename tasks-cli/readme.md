# CLI Task Manager

A simple **command-line task manager** written in Go.
Supports creating, listing, completing tasks, and crash-safe JSON persistence. Created for purpose of learning Go

---

## Features

* Add tasks with automatic creation date
* List all tasks with status and date added
* Mark tasks as completed
* Persistent storage using JSON (`data.json`)
* Crash-safe, atomic writes (`data-temp.json → rename → data.json`)
* Automatic backup if task file is corrupt (`data-backup-<timestamp>.json`)
* Lightweight, single-user CLI

---

## Installation

1. **Clone the repository**:

```bash
git clone <repo-url>
cd <repo-directory>
```

2. **Build the CLI**:

```bash
go build -o taskcli main.go
```

3. **Run the CLI**:

```bash
./taskcli
```

> Ensure you have Go installed (version 1.20+)

---

## Usage

When you run the CLI, you’ll see a **main menu**:

```
Welcome to CLI Task Manager...
================================
Menu:
 1. Add a new Task
 2. Show All Task(<number-of-tasks>)
 3. Mark a task as complete
 4. Exit
Choice:
```

### Menu Options

1. **Add a new Task**

   * Enter your task description
   * Task is saved immediately with status `pending`

2. **Show All Tasks**

   * Lists all tasks with index, content, status, and date added

3. **Mark a task as complete**

   * Enter the task number to mark it completed
   * Tasks are immediately saved

4. **Exit**

   * Close the CLI

> The CLI also handles invalid inputs gracefully and allows you to go back from task completion using `-1`.

---

## Persistence

* Tasks are stored in `data.json` in the same directory.
* Each mutation (add/complete) triggers a **full save**.
* To prevent corruption:

  * Data is first written to `data-temp.json`
  * Then **atomically renamed** to `data.json`
* If `data.json` is corrupted on load:

  * The corrupt file is backed up automatically with a timestamp
  * The CLI starts with an empty task list

Example backup filename:

```
data-backup-2026-01-01_18-30-25.json
```

---

## File Structure

* `main.go` — main CLI logic and task functions
* `data.json` — persistent task storage (auto-generated)
* `data-temp.json` — temporary file for atomic writes
* `data-backup-<timestamp>.json` — automatic backups for corrupted files

---

## Notes

* Designed for **single-user, local use**
* Does **not support multi-user or concurrent access**
* Tasks are simple and stored in **JSON format** for easy editing or integration
* Fully written in **Go**, no external dependencies

---

## Example Session

```
Welcome to CLI Task Manager...
================================
Menu:
 1. Add a new Task
 2. Show All Task(0)
 3. Mark a task as complete
 4. Exit
Choice: 1

Enter new Task: Finish writing README
New Task Added Successfully.

Menu:
 1. Add a new Task
 2. Show All Task(1)
 3. Mark a task as complete
 4. Exit
Choice: 2

1.| Finish writing README | pending | 01-01-2026 |

Choice: 3
Enter the number of task you want to mark complete [-1 to go back]: 1
Tasks Saved Successfully

Menu:
 1. Add a new Task
 2. Show All Task(1)
 3. Mark a task as complete
 4. Exit
Choice: 2

1.| Finish writing README | completed | 01-01-2026 |
```