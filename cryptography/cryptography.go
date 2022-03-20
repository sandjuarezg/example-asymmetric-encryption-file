package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"os"
	"strings"
)

func existKeysFiles() (err error) {
	_, err = os.Stat("./keys/key.priv")
	if err != nil {
		return
	}

	return
}

func getPublicKey() (publicKey *rsa.PublicKey, err error) {
	content, err := os.ReadFile("./keys/key.pub")
	if err != nil {
		return
	}

	publicKey, err = x509.ParsePKCS1PublicKey(content)
	if err != nil {
		return
	}

	return
}

func getPrivateKey() (privateKey *rsa.PrivateKey, err error) {
	content, err := os.ReadFile("./keys/key.priv")
	if err != nil {
		return
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(content)
	if err != nil {
		return
	}

	return
}

func GenerateKeysFiles() (err error) {
	err = existKeysFiles()
	if !os.IsNotExist(err) {
		return
	}

	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	auxPrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	err = os.WriteFile("./keys/key.priv", auxPrivateKey, 0600)
	if err != nil {
		return
	}

	// public key
	publicKey := &privateKey.PublicKey

	auxPublicKey := x509.MarshalPKCS1PublicKey(publicKey)

	err = os.WriteFile("./keys/key.pub", auxPublicKey, 0600)
	if err != nil {
		return
	}

	return
}

func EncryptFile(filename string) (err error) {
	content, err := os.ReadFile(fmt.Sprintf("./files/%s", filename))
	if err != nil {
		return
	}

	publicKey, err := getPublicKey()
	if err != nil {
		return
	}

	encrypt, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, content, nil)
	if err != nil {
		return
	}

	err = os.WriteFile(fmt.Sprintf("./files/%s.encrypt", filename), encrypt, 0600)
	if err != nil {
		return
	}

	return
}

func DecryptFile(filename string) (err error) {
	content, err := os.ReadFile(fmt.Sprintf("./files/%s", filename))
	if err != nil {
		return
	}

	privateKey, err := getPrivateKey()
	if err != nil {
		return
	}

	decrypt, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, content, nil)
	if err != nil {
		return
	}

	err = os.WriteFile(fmt.Sprintf("./files/%s.decrypt", strings.TrimSuffix(filename, ".encrypt")), decrypt, 0600)
	if err != nil {
		return
	}

	return
}
