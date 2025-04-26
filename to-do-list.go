package main

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
)

type Task struct {
	Name string
	isCompleted bool
}

const greeting string = `
To-Do CLI
-----------
1. Add a new task
2. List all tasks
3. Mark a task as completed
4. Delete a task
5. Exit

Enter your choice (1-5):
`

func main() {
	// All Tasks
	var tasks []Task

	// We Will Store the User Response in This
	var response uint8

	// Printing the Greeting
	fmt.Println(greeting)
	fmt.Scanln(&response)

	for response != 5 {
		switch response {
		case 1:
			
		}
	}
}

func add(new_task Task) {}

func remove(task Task) {}

func list_tasks() {}

func mark_completed(task_index uint8) {}

func save(tasks []Task, file *os.File) error {
	if file == nil {
		var err error
		file, err = check_file()
		if err != nil { return err }
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func load(to []Task) error {
	check_file()
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &to)
	if err != nil {
		return err
	}

	return nil
}

func check_file() (*os.File, error) {
	var file *os.File
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var has_file bool = false
	for _, entry := range entries {
		if strings.Contains(entry.Name(), "tasks.json") == true {
			has_file = true
		}
	}
	if has_file == false {
		file, err = os.Create("tasks.json")
		if err != nil { return nil, err }
	}
	return file, nil
}
