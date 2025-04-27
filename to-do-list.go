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

Enter your choice (1-5): `

// All Tasks
var tasks []Task

func main() {
	load(tasks)
	// We Will Store the User Response in This
	var response uint8 = get_response(greeting)

	for response != 5 {
		switch response {
		case 1:
			var title string = get_string("Name of the Task: ")
			add(Task{title, false})
			response = get_response(greeting)

		case 2:
			list_tasks()
			response = get_response(greeting)

		case 3:
			list_tasks()
			var index uint8 = get_response("Which task do you want to mark completed?")
			mark_completed(index)
			response = get_response(greeting)
			
		case 4:
			list_tasks()
			var index uint8 = get_response("Which task do you want to delete?")
			remove(index)
			response = get_response(greeting)
		}
	}
	file, err := check_file()
	if err != nil { fmt.Println("Could not save file") }
	save(tasks, file)
}

func add(new_task Task) {
	tasks = append(tasks, new_task)
}

func remove(index uint8) {
	new_array := tasks[0:index]
    length := uint8(len(tasks))
	for i := index + 1; i < length; i++ {
		new_array = append(new_array, tasks[i])
	}
	tasks = new_array
}

func list_tasks() {
	fmt.Println("Tasks: ")
	for i, task := range tasks {
		if task.isCompleted == true {
			fmt.Printf("[Completed] %v. %v", i, task.Name)
		} else {
			fmt.Printf("[Pending] %v. %v", i, task.Name)
		}
	}
	fmt.Println("")
}

func mark_completed(task_index uint8) {
	tasks[task_index].isCompleted = true
}

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

func get_response(prompt string) uint8 {
	var response uint8
	fmt.Printf("%v", prompt)
	fmt.Scan(&response)

	return response
}

func get_string(prompt string) string {
	var response string
	fmt.Printf("%v", prompt)
	fmt.Scanln(&response)

	return response
}
