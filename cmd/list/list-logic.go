package list

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker/internal/task"
)

var taskTemplate = 
`
Number: #%v
Task: %s 
Description: %s
Status: %s
Creation date: %s
Update date: %s
Tags: %s
`

func listTask() {
	file, err := os.ReadFile("file.json")
	if err == os.ErrNotExist {
		fmt.Print("There is no tasks yet!")
	}

	data := []task.Task{}
	json.Unmarshal(file, &data)
	if len(data) == 0 {
		fmt.Printf("There are no tasks! You may go chill... \n")
	} else {
		fmt.Printf("To do:\n")
		for num, task := range data {
			fmt.Printf(taskTemplate,
				num + 1, 
				task.Title, 
				task.Description,
				task.Status,
				task.AddedDate,
				task.UpdatedDate,
				task.Tags,
			)
		}
	}
	
}