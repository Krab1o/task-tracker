package listlogic

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	jsonstruct "task-tracker/internal/json_struct"
)

var taskHeader =
"Number\tTask\tDescription\t\tStatus\tCreationDate\t\tUpdateDate\t\tTags"

var taskTemplate = 
"#%v\t%s\t%s\t%s\t%s\t%s\t%s\n"

func ListTask(statuses []string) {
	file, err := os.ReadFile("file.json")
	if err == os.ErrNotExist {
		fmt.Print("There is no tasks yet!")
	}

	data := jsonstruct.Data{}
	json.Unmarshal(file, &data)

	if len(data.Tasks) == 0 {
		fmt.Printf("There are no tasks! You may go chill... \n")
	} else {
		fmt.Println(taskHeader)
		for _, task := range data.Tasks {
			if (slices.Contains(statuses, task.Status)) {
				fmt.Printf(taskTemplate,
					task.ID, 
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
}