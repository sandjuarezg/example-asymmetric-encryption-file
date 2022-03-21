package main

import (
	"bufio"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
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

			var (
				filename string
				hashAux  string
				hash     hash.Hash
				label    string
			)

			fmt.Println()
			fmt.Println("< Enter to skip >")
			fmt.Print("Enter name file: ")
			fmt.Scanln(&filename)
			fmt.Print("Enter hash (sha1, sha256, sha512): ")
			fmt.Scanln(&hashAux)
			fmt.Print("Enter label: ")
			fmt.Scanln(&label)

			switch hashAux {
			case "sha1":
				hash = sha1.New()
			case "sha256":
				hash = sha256.New()
			case "sha512":
				hash = sha512.New()
			default:

				log.Println("hash no found")
				continue

			}

			err = cryptography.EncryptFile(filename, hash, []byte(label))
			if err != nil {
				log.Println(err)
				continue
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

			var (
				filename string
				hashAux  string
				hash     hash.Hash
				label    string
			)

			fmt.Println()
			fmt.Println("< Enter to skip >")
			fmt.Print("Enter name file: ")
			fmt.Scanln(&filename)
			fmt.Print("Enter hash (sha1, sha256, sha512): ")
			fmt.Scanln(&hashAux)
			fmt.Print("Enter label: ")
			fmt.Scanln(&label)

			switch hashAux {
			case "sha1":
				hash = sha1.New()
			case "sha256":
				hash = sha256.New()
			case "sha512":
				hash = sha512.New()
			default:
				log.Println("hash no found")
				continue
			}

			err = cryptography.DecryptFile(filename, hash, []byte(label))
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			fmt.Println("- Decryption file created -")
			fmt.Println()

		case 3:

			var filename string
			fmt.Println()
			fmt.Print("Enter name file: ")
			fmt.Scanln(&filename)

			fmt.Print("Enter text: ")
			content, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Println(err)
				continue
			}

			err = functionality.CreateFile(filename, string(content))
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			fmt.Println("- File created -")
			fmt.Println()

		}

	}
}
