/* Detect AES in ECB mode */

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

const (
	FILE = "./8.txt"
)

func main() {
	f, err := os.Open(FILE)
	if err != nil {
		fmt.Printf("[ERROR] Could not open file %v: %v\n", FILE, f)
		return
	}
	defer f.Close()

	
	var lines [][]byte
	sc  := bufio.NewScanner(f)
	for sc.Scan() {
		ln := sc.Bytes()
		dst := make([]byte, hex.DecodedLen(len(ln)))
		n, err := hex.Decode(dst, ln)
		if err != nil {
			fmt.Printf("[ERROR] Could not decode %v: %v\n", FILE, err)
			return
		}
		
		lines = append(lines, dst[:n])
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("[ERROR] Scan failed: %v\n", err)
		return
	}

	var maxScore int
	var bestIdx int
	for idx, bytes := range lines {
		ctr := make(map[string]int)
		for i := 0; i+16 < len(bytes); i += 16 {
			ctr[string(bytes[i:i+16])]++
		}

		score := 0
		for _, c := range ctr {
			if c > score {
				score = c
			}
		}

		if score > maxScore {
			maxScore = score
			bestIdx = idx
		}
	}

	fmt.Printf("[INFO] ECB line: %v (block repeats %v times)\n", bestIdx, maxScore)
}
