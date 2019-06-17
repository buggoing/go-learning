package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func NewRSAPrivateKey(bits int) (privateKey *rsa.PrivateKey, err error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func EncryptRSA(data string, pubkey rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")

	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &pubkey, []byte(data), label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return "Error from encryption"
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptRSA(data string, privKey *rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(data)
	label := []byte("OAEP Encrypted")

	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, privKey, ct, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return "Error from Decryption"
	}

	return string(plaintext)
}

func SignRSA(plaintext string, privKey *rsa.PrivateKey) string {

	rng := rand.Reader
	hashed := sha256.Sum256([]byte(plaintext))
	signature, err := rsa.SignPKCS1v15(rng, privKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return "Error from signing"
	}
	return base64.StdEncoding.EncodeToString(signature)
}

func VerifyRSA(signature string, plaintext string, pubkey rsa.PublicKey) bool {
	sig, _ := base64.StdEncoding.DecodeString(signature)
	hashed := sha256.Sum256([]byte(plaintext))
	err := rsa.VerifyPKCS1v15(&pubkey, crypto.SHA256, hashed[:], sig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return false
	}
	return true
}
func main() {
	privateKey, err := NewRSAPrivateKey(2048)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.PublicKey
	fmt.Println("privateKey: ", *privateKey)
	fmt.Println("publicKey: ", publicKey)

	message := "hello, 世界"
	encryptedData := EncryptRSA(message, publicKey)
	fmt.Println("encrypt: ", encryptedData)

	decryptedData := DecryptRSA(encryptedData, privateKey)
	fmt.Println("decrypt: ", decryptedData)

	singedData := SignRSA(message, privateKey)
	fmt.Println("signed: ", singedData)

	isVerified := VerifyRSA(singedData, message, publicKey)
	fmt.Println("isVerified: ", isVerified)

}
