package add

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
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

func loadData() []task.Task {
	file, err := os.ReadFile("file.json")
	//TODO: write switch statement
	switch {
	case errors.Is(err, os.ErrNotExist):
	if _, createErr := os.Create("file.json"); createErr != nil {
		log.Fatal(createErr)
	}
	case err != nil:
		log.Fatal(err)
	default:
		// Do nothing
	}
	
	data := []task.Task{}
	json.Unmarshal(file, &data)
	return data
}

func addTask(taskName string) {
	currentTime := time.Now().Format(time.DateTime)
	addedTask := &task.Task {
		Title: taskName,
		Status: status.ToDo.String(),
		AddedDate: currentTime,
	}

	data := loadData()
	data = append(data, *addedTask)

	dataBinary, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	writeFile(dataBinary)	
	fmt.Printf("Task %s added successfully!\n", taskName)
}