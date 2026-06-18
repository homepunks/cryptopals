/* AES in ECB mode */

package main

import (
	"fmt"
	"os"
	"encoding/base64"
	"crypto/aes"
)

const (
	KEY = "YELLOW SUBMARINE"
	FILE = "./7.txt"
	CHUNKSIZE = 128 / 8
)

func main() {
	KEYBYTES := []byte(KEY)

	bytes64, err := os.ReadFile(FILE)
	if err != nil {
		fmt.Printf("[ERROR] Could not read file contents of %v\n", FILE)
		return
	}

	ciphertext := make([]byte, base64.StdEncoding.DecodedLen(len(bytes64)))
	n, err := base64.StdEncoding.Decode(ciphertext, bytes64)
	if err != nil {
		fmt.Printf("[ERROR] Could not decode binary content of %v: %v\n", FILE, err)
		return
	}
	ciphertext = ciphertext[:n]

	block, err := aes.NewCipher(KEYBYTES)
	if err != nil {
		fmt.Printf("[ERROR] Failed to create cipher: %v\n", err)
		return
	}

	var decrypted string
	plaintext := make([]byte, aes.BlockSize) // just found it
	chunks := make([][]byte, len(ciphertext)/CHUNKSIZE)
	for i := 0; i < len(chunks); i++ {
		chunks[i] = ciphertext[16*i:16*(i+1)]
		block.Decrypt(plaintext, chunks[i])
		decrypted += string(plaintext)
	}

	fmt.Println(decrypted)
}
