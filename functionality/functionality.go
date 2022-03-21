package functionality

import (
	"fmt"
	"os"
)

func PreparePathDir() (err error) {
	err = os.MkdirAll("./files/", 0700)
	if err != nil {
		return
	}

	err = os.MkdirAll("./keys/", 0700)
	if err != nil {
		return
	}

	return
}

func CreateSamplesFiles() (err error) {
	err = os.WriteFile("./files/message.txt", []byte("Do you want to know a secret? 7u7"), 0600)
	if err != nil {
		return
	}

	err = os.WriteFile("./files/hello.txt", []byte("Hellooooooow"), 0600)
	if err != nil {
		return
	}

	return
}

func CreateFile(filename string, content string) (err error) {
	err = os.WriteFile(fmt.Sprintf("./files/%s", filename), []byte(content), 0600)
	if err != nil {
		return
	}

	return
}
