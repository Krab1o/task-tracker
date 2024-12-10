package filelogic

import "os"

func WriteFile(data []byte) {
	permissions := 0644
	err := os.WriteFile("file.json", data, os.FileMode(permissions))
	if err != nil {
		panic(err)
	}
}