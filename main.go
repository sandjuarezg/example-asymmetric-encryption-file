package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
				// filepath
				if strings.HasSuffix(file.Name(), ".txt") {
					fmt.Println(file.Name())
				}
			}

			fmt.Print("Enter name file: ")
			filename, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Fatal(err)
			}

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
				if strings.HasSuffix(file.Name(), ".encrypt") {
					fmt.Println(file.Name())
				}
			}

			fmt.Printf("Enter name file: ")
			filename, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Fatal(err)
			}

			err = cryptography.DecryptFile(string(filename))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()
			fmt.Println("- Decryption file created -")
			fmt.Println()

		}

	}
}
