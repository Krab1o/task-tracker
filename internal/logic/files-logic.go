package logic

import (
	"fmt"
	"os"
)

var directoryPath = os.Getenv("HOME") + "/.task-tracker"
var dataPath = directoryPath + "/data.json"
var directoryPermissions = 0777
var dataPermissions = 0644

func createDirectory() {
	folderErr := os.Mkdir(directoryPath, os.FileMode(directoryPermissions))
	if (folderErr != nil) {
		fmt.Print(folderErr.Error())
	}
}

func writeFile(data []byte) {
	err := os.WriteFile(dataPath, data, os.FileMode(dataPermissions))
	if err != nil {
		panic(err)
	}
}