package add

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	jsonstruct "task-tracker/internal/json_struct"
	"task-tracker/internal/status"
	"task-tracker/internal/task"
	"time"
)

func writeFile(data []byte) {
	permissions := 0644
	err := os.WriteFile("file.json", data, os.FileMode(permissions))
	if err != nil {
		panic(err)
	}
}

func addTask(taskName string) {
	file, err := os.ReadFile("file.json")

	switch {
	case errors.Is(err, os.ErrNotExist):
	if _, createErr := os.Create("file.json"); createErr != nil {
		log.Fatal(createErr)
	}
	case err != nil:
		log.Fatal(err)
	default:
		// Left blank intentionally
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)

	ID := data.ID + 1 
	
	currentTime := time.Now().Format(time.DateTime)
	addedTask := &task.Task {
		ID: 		ID,
		Title: 		taskName,
		Status: 	status.ToDo.String(),
		AddedDate: 	currentTime,
	}

	data.Tasks = append(data.Tasks, *addedTask)
	data.ID = ID

	dataBinary, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	writeFile(dataBinary)	
	fmt.Printf("Task %s added successfully!\n", taskName)
}