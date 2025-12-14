/* Single-byte XOR cipher */

package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var inputHex string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	data, err := hex.DecodeString(inputHex)
	if err != nil {
		fmt.Printf("[ERROR] Could not decode %s: %v\n", inputHex, err)
		return
	}

	key, plaintext := solveSingleByteXor(data)
	fmt.Printf("Most likely to be the key: %c (ASCII %d)\n", key, key)
	fmt.Printf("Corresponding plaintext: %s\n", plaintext)
	
}

func solveSingleByteXor(ciphertext []byte) (byte, string) {
	var bestKey byte
	var bestScore int
	var bestPlaintext []byte
	
	for key := 0; key < 256; key++ {
		plaintext := make([]byte, len(ciphertext))

		for i := 0; i < len(ciphertext); i++ {
			plaintext[i] = ciphertext[i] ^ byte(key)
		}

		score := cryptoScore(plaintext)
		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestPlaintext = plaintext
		}
	}

	return bestKey, string(bestPlaintext)
}

func cryptoScore(text []byte) int {
	score := 0
	for _, ch := range text {
		// all scores are arbitrary
		if ch >= 'A' && ch <= 'Z' {
			score += 1
		}
		
		if ch >= 'a' && ch <= 'z' {
			score += 2
		}

		if ch == ' ' {
			score += 3
		}

		if ch == '.' && ch == ',' {
			score += 1
		}
		
		if (ch < 32 || ch > 126) && ch != '\t' && ch != '\n' {
			score -= 10
		}
	}
	
	return score
}
