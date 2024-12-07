package add

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type task struct {
	Title string		`json:"stringValue"`
	DateAdded time.Time
}

func writeFile(data []byte) {
	permissions := 0644
	err := os.WriteFile("file.json", data, os.FileMode(permissions))
	if err != nil {
		fmt.Print("Error occured")
	}
}

func addTask(taskName string) {
	addedTask := &task{
		Title: taskName,
		DateAdded: time.Now(),
	}
	data, _ := json.Marshal(addedTask)
	writeFile(data)
	fmt.Printf("Task %s added successfully!\n", taskName)
}