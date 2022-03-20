package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	priASN1 := x509.MarshalPKCS1PrivateKey(privateKey)

	priBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: priASN1,
	})

	err = os.WriteFile("./keys/key.priv", priBytes, 0600)
	if err != nil {
		log.Fatal(err)
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	err = os.WriteFile("./keys/key.pub", pubBytes, 0600)
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := x509.ParsePKCS1PublicKey(pubBytes)
	if err != nil {
		log.Fatal(err)
	}

	encrypt, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte("hola? estas?"), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(encrypt))
}
