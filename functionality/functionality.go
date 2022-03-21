package functionality

import (
	"fmt"
	"os"
)

// CreateDirPath Create directory paths
//
//  @return1 (err): error variable
func CreateDirPath() (err error) {
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

// CreateSamplesFiles Create sample files to use
//
//  @return1 (err): error variable
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

// CreateFile Create a new file with content
//  @param1 (filename): name of file
//  @param2 (content): content of file
//
//  @return1 (err): error variable
func CreateFile(filename string, content string) (err error) {
	err = os.WriteFile(fmt.Sprintf("./files/%s", filename), []byte(content), 0600)
	if err != nil {
		return
	}

	return
}
