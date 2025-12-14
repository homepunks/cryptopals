/* Imlement repeating-key XOR */

package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var key string = "ICE"
	var input string = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	
	ciphertext := repeatedKeyXor([]byte(input), []byte(key))
	ciphertextHex := hex.EncodeToString(ciphertext)
	fmt.Println(ciphertextHex)
}

func repeatedKeyXor(plaintext []byte, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i % len(key)]
	}

	return ciphertext
}
