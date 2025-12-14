/* Detect single-byte XOR */

package main

import (
	"os"
	"fmt"
	"encoding/hex"
	"bufio"
)

func main() {
	f, err := os.Open("./4.txt")
	if err != nil {
		fmt.Errorf("[ERROR] Could not open the text file: %v\n", err)
		return
	}
	
	defer f.Close()

	var bestScore int
	var bestKey byte
	var plaintext string
	var lineNum int = 1
	var original string
	
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cipher, err := hex.DecodeString(line)
		if err != nil {
			fmt.Errorf("[ERROR] Could not decode %s: %v\n", line, err)
			return
		}

		key, plain, score := solveSingleByteXor(cipher)
		if score > bestScore {
			bestScore = score
			bestKey = key
			plaintext = plain
			original = line
		}

		lineNum++
	}

	fmt.Printf("Most likely to be the key: %c (ASCII %d)\n", bestKey, bestKey)
	fmt.Printf("Corresponding plaintext: %s", plaintext)
	fmt.Printf("From: %s (line %d)\n", original, lineNum)	
}

func solveSingleByteXor(ciphertext []byte) (byte, string, int) {
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

	return bestKey, string(bestPlaintext), bestScore
}

func cryptoScore(text []byte) int {
	score := 0
	for _, ch := range text {
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
