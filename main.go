package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sandjuarezg/example-asymmetric-encryption-file/cryptography"
	"github.com/sandjuarezg/example-asymmetric-encryption-file/functionality"
)

func main() {
	var flag bool
	var opc int

	err := functionality.PreparePathDir()
	if err != nil {
		log.Fatal(err)
	}

	err = functionality.CreateSamplesFiles()
	if err != nil {
		log.Fatal(err)
	}

	err = cryptography.GenerateKeysFiles()
	if err != nil {
		log.Fatal(err)
	}

	for !flag {

		opc = 0

		fmt.Println("- Welcome -")
		fmt.Println("0. Exit")
		fmt.Println("1. Encrypt")
		fmt.Println("2. Decrypt")
		fmt.Println("3. Create file")
		fmt.Scanln(&opc)

		switch opc {

		case 0:
			flag = true
		case 1:

			files, err := os.ReadDir("./files")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()

			for _, file := range files {
				if filepath.Ext(file.Name()) != ".encrypt" && filepath.Ext(file.Name()) != ".decrypt" {
					fmt.Println(file.Name())
				}
			}

			var filename string
			fmt.Print("Enter name file: ")
			fmt.Scan(&filename)

			err = cryptography.EncryptFile(string(filename))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()
			fmt.Println("- Encryption file created -")
			fmt.Println()

		case 2:

			files, err := os.ReadDir("./files")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()

			for _, file := range files {
				if filepath.Ext(file.Name()) == ".encrypt" {
					fmt.Println(file.Name())
				}
			}

			var filename string
			fmt.Print("Enter name file: ")
			fmt.Scan(&filename)

			err = cryptography.DecryptFile(string(filename))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()
			fmt.Println("- Decryption file created -")
			fmt.Println()

		case 3:

			fmt.Println()
			fmt.Print("Enter file name: ")

			var filename string
			fmt.Print("Enter name file: ")
			fmt.Scan(&filename)

			fmt.Print("Enter text: ")
			content, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Fatal(err)
			}

			err = functionality.CreateFile(filename, string(content))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()
			fmt.Println("- File created -")
			fmt.Println()

		}

	}
}
