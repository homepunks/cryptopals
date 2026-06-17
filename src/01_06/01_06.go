/* Break repeating-key XOR */

package main

import (
	"fmt"
	"os"
)

const (
	FOO = "this is a test"
	BAR = "wokka wokka!!!"
)

func main() {
	f, err := os.Open("./6.txt")
	if err != nil {
		fmt.Errorf("[ERROR] Could not open the text file: %v\n", err)
		return
	}
	defer f.Close()

	fmt.Println(hammingDist(FOO, BAR))
}

func hammingDist(a, b string) int {
	if len(a) != len(b) {
		return -1
	}

	a = stringBits(a)
	b = stringBits(b)

	var dist int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist
}

func stringBits(s string) string {
	var bitString string
	for i := 0; i < len(s); i++ {
		bitString += fmt.Sprintf("%08b", s[i])
	}

	return bitString
}
