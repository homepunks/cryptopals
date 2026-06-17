/* Break repeating-key XOR */

package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"os"
)

func main() {
	bytes64, err := os.ReadFile("./6.txt")
	if err != nil {
		panic(fmt.Errorf("[ERROR] Could not read the text file: %v\n", err))
	}

	cipher := make([]byte, base64.StdEncoding.DecodedLen(len(bytes64)))
	n, err := base64.StdEncoding.Decode(cipher, bytes64)
	if err != nil {
		panic(fmt.Errorf("[ERROR] i'm not handling this bro"))
	}

	// reslice to get rid of trailing zeros
	cipher = cipher[:n]

	var keySize int
	diffNormalizedLowest := math.MaxFloat64
	for i := 2; i <= 40; i++ {
		var totalDist int
		pairs := 0
		for j := 0; j+2*i <= len(cipher); j += i {
			totalDist += hammingDist(cipher[j:j+i], cipher[j+i:j+2*i])
			pairs++
		}
		diffNormalized := float64(totalDist) / float64(i) / float64(pairs)
		if diffNormalized < diffNormalizedLowest {
			diffNormalizedLowest = diffNormalized
			keySize = i
		}
	}

	blocks := make([][]byte, keySize)
	for i := 0; i < len(cipher); i++ {
		idx := i % keySize
		blocks[idx] = append(blocks[idx], cipher[i])
	}

	keyFull := make([]byte, keySize)
	for i := 0; i < keySize; i++ {
		keyByte, _, _ := solveSingleByteXor(blocks[i])
		keyFull[i] = keyByte
	}

	decrypted := make([]byte, len(cipher))
	for i := 0; i < len(cipher); i++ {
		decrypted[i] = cipher[i] ^ keyFull[i%keySize]
	}

	fmt.Println(string(decrypted))
}

func hammingDist(a, b []byte) int {
	if len(a) != len(b) {
		return -1
	}

	aBits := bytesToBits(a)
	bBits := bytesToBits(b)

	var dist int
	for i := 0; i < len(aBits); i++ {
		if aBits[i] != bBits[i] {
			dist++
		}
	}

	return dist
}

func bytesToBits(b []byte) string {
	var bitString string
	for i := 0; i < len(b); i++ {
		bitString += fmt.Sprintf("%08b", b[i])
	}

	return bitString
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

		if ch == '.' || ch == ',' {
			score += 1
		}

		if (ch < 32 || ch > 126) && ch != '\t' && ch != '\n' {
			score -= 10
		}
	}

	return score
}
