package markdone

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	jsonstruct "task-tracker/internal/json_struct"
	"task-tracker/internal/status"
	"task-tracker/internal/task"
	"time"

	"github.com/guregu/null"
)

func writeFile(data []byte) {
	permissions := 0644
	err := os.WriteFile("file.json", data, os.FileMode(permissions))
	if err != nil {
		panic(err)
	}
}

func markDoneTask(ID uint32) {
	file, err := os.ReadFile("file.json")

	if err != nil {
		log.Fatal(err)
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)
	var doneTask task.Task

	for i := range data.Tasks {
		if (data.Tasks[i].ID == ID) {
			doneTask = data.Tasks[i]
			data.Tasks[i].Status = status.Done.String()
			data.Tasks[i].UpdatedDate = null.StringFrom(time.Now().Format(time.DateTime))
		}
	}

	dataBinary, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	//TODO: Refactor multiple function definition
	writeFile(dataBinary)	
	fmt.Printf("Task %s done successfully!\n", doneTask.Title)

}