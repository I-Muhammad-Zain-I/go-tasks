package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
*
1. Add a task by reading from user input prompt
2. View a tasks/Show all tasks
3. mark task as completed
==================================================
*/

const (
	Pending   = "pending"
	Completed = "completed"
)

type Task struct {
	Content   string `json:"content"`
	DateAdded string `json:"dateAdded"`
	Status    string `json:"status"`
}

func getCurrentDate() string {
	var currentDate string = time.Now().Format("02-01-2006")
	return currentDate
}

func createNewTask(content string) Task {
	currentDate := getCurrentDate()
	newTask := Task{Content: content, DateAdded: currentDate, Status: Pending}
	return newTask
}

func printAllTask(taskList []Task) {
	fmt.Print("\n")
	for i, task := range taskList {
		fmt.Printf("%d.| %s | %s | %s | \n", i+1, task.Content, task.Status, task.DateAdded)
	}
	fmt.Print("\n")
}

func handleListTask(taskList []Task) AppState {
	printAllTask(taskList)
	return MainMenu
}

func completeTask(taskList []Task, index int) {
	taskList[index].Status = Completed
}

func readInt(reader *bufio.Reader) int {
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		fmt.Println("Error: An error has occured during input parsing")
	}

	intInput, _ := strconv.Atoi(input)
	return intInput
}

type AppState int

const (
	MainMenu AppState = iota
	AddTask
	ListTask
	CompleteTask
	RecoverBackup
	Exit
)

func handleMainMenu(reader *bufio.Reader, taskCount int) AppState {
	fmt.Println("Welcome to CLI Task Manager...")
	fmt.Println("================================")

	fmt.Printf("Menu:\n 1. Add a new Task\n 2. Show All Task(%d)\n 3. Mark a task as complete\n 4. Exit\n", taskCount)
	fmt.Print("Choice: ")

	choice := readInt(reader)

	switch choice {
	case 1:
		return AddTask
	case 2:
		return ListTask
	case 3:
		return CompleteTask
	case 4:
		return Exit
	default:
		return MainMenu
	}
}
func checkIfFileAlreadyExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func handleAddTask(reader *bufio.Reader, taskList []Task) (AppState, []Task) {
	fmt.Print("\nEnter new Task: ")
	newTask, _ := reader.ReadString('\n')
	newTask = strings.TrimSpace(newTask)
	taskList = append(taskList, createNewTask(newTask))
	fmt.Println("New Task Added Successfully.")
	saveTask(taskList)
	return MainMenu, taskList
}

func handleCompleteTask(reader *bufio.Reader, taskList []Task) AppState {

	if len(taskList) == 0 {
		fmt.Println("No Task are present Currently")
		return MainMenu
	}

	fmt.Println("Enter the number of task you want to mark complete [-1 to go back]")
	printAllTask(taskList)

	fmt.Print("Choice: ")
	choice := readInt(reader)

	if choice == -1 {
		fmt.Println("Go Back")
		return MainMenu
	}

	if choice > len(taskList) || choice == 0 || choice < -1 {
		fmt.Println("Error: Incorrect Choice entered")
		return CompleteTask
	}

	completeTask(taskList, choice-1)
	saveTask(taskList)

	return MainMenu
}

func initializeEmptyTasks() []Task {
	taskList := make([]Task, 0, 10)
	return taskList
}

func loadTask() []Task {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println("Error: an error occured during loading tasks.")
		return initializeEmptyTasks()
	}
	defer jsonFile.Close()

	byteValues, _ := io.ReadAll(jsonFile)

	var taskList []Task

	err = json.Unmarshal(byteValues, &taskList)

	if err != nil {
		fmt.Println("An error occured during un marshal process")
		jsonFile.Close()
		backupFileName := fmt.Sprintf("data-backup-%s.json", time.Now().Format("2006-01-02_15-04-05"))
		fmt.Println(backupFileName)
		renameErr := os.Rename("data.json", backupFileName)
		if renameErr != nil {
			fmt.Println(renameErr)
		}
		return initializeEmptyTasks()
	}
	return taskList
}

func saveTask(taskList []Task) {

	byteValues, err := json.Marshal(taskList)

	if err != nil {
		fmt.Println("Error: an error occured during saving Tasks")
		return
	}

	err = os.WriteFile("data-temp.json", byteValues, 0644)

	if err != nil {
		fmt.Println("Error: an error occured during saving Tasks")
		return
	}

	err = os.Rename("data-temp.json", "data.json")

	if err != nil {
		fmt.Println("Error: an error occured during saving Tasks")
		return
	}

	fmt.Println("Tasks Saved Successfully")

}

func main() {
	state := MainMenu

	taskList := loadTask()

	reader := bufio.NewReader(os.Stdin)

	for state != Exit {
		switch state {
		case MainMenu:
			state = handleMainMenu(reader, len(taskList))
		case AddTask:
			state, taskList = handleAddTask(reader, taskList)
		case ListTask:
			state = handleListTask(taskList)
		case CompleteTask:
			state = handleCompleteTask(reader, taskList)
		case RecoverBackup:
			state = RecoverBackup
		}
	}
}
