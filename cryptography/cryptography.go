package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"hash"
	"os"
	"strings"
)

// existKeysFiles Check existence of key files
//
//  @return1 (exist): file existence
func existKeysFiles() (exist bool) {
	_, err := os.Stat("./keys/key.priv")
	if os.IsNotExist(err) {
		return
	}

	_, err = os.Stat("./keys/key.pub")
	if os.IsNotExist(err) {
		return
	}

	exist = true

	return
}

// getPublicKeyFromFile Get file public key
//
//  @return1 (publicKey): public key
//  @return2 (err): error variable
func getPublicKeyFromFile() (publicKey *rsa.PublicKey, err error) {
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

// getPrivateKeyFromFile Get file private key
//
//  @return1 (privateKey): private key
//  @return2 (err): error variable
func getPrivateKeyFromFile() (privateKey *rsa.PrivateKey, err error) {
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

// CreateKeysFiles Create files for keys
//
//  @return1 (err): error variable
func CreateKeysFiles() (err error) {
	exist := existKeysFiles()
	if exist {
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

// EncryptFile Encrypt file
//  @param1 (filename): name of file
//  @param2 (hash): hash function
//  @param3 (label): label to context to the message
//
//  @return1 (err): error variable
func EncryptFile(filename string, hash hash.Hash, label []byte) (err error) {
	content, err := os.ReadFile(fmt.Sprintf("./files/%s", filename))
	if err != nil {
		return
	}

	publicKey, err := getPublicKeyFromFile()
	if err != nil {
		return
	}

	encrypt, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, content, label)
	if err != nil {
		return
	}

	err = os.WriteFile(fmt.Sprintf("./files/%s.encrypt", filename), encrypt, 0600)
	if err != nil {
		return
	}

	return
}

// DecryptFile Decrypt file
//  @param1 (filename): name of file
//  @param2 (hash): hash function
//  @param3 (label): label to the message
//
//  @return1 (err): error variable
func DecryptFile(filename string, hash hash.Hash, label []byte) (err error) {
	content, err := os.ReadFile(fmt.Sprintf("./files/%s", filename))
	if err != nil {
		return
	}

	privateKey, err := getPrivateKeyFromFile()
	if err != nil {
		return
	}

	decrypt, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, content, label)
	if err != nil {
		if err == rsa.ErrDecryption {
			err = errors.New("data doesn't match to decrypt")
		}

		return
	}

	err = os.WriteFile(fmt.Sprintf("./files/%s.decrypt", strings.TrimSuffix(filename, ".encrypt")), decrypt, 0600)
	if err != nil {
		return
	}

	return
}
