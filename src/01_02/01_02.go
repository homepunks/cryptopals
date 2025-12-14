/* Fixed XOR */

package main

import (
	"encoding/hex"
	"fmt"
)


func main() {
	var input string = "1c0111001f010100061a024b53535009181c"
	var fixed string = "686974207468652062756c6c277320657965"
	
	inputHex, err := hex.DecodeString(input)
	if err != nil {
		fmt.Errorf("[ERROR] Could not convert %s to hex: %v\n", input, err)
		return
	}
	fixedHex, err := hex.DecodeString(fixed)
	if err != nil {
		fmt.Errorf("[ERROR] Could not convert %s to hex: %v\n", fixed, err)
		return
	}
	
	res := fixedXor(inputHex, fixedHex)
	fmt.Println(string(res))
}

func fixedXor(input []byte, fixed []byte) []byte {
	res := make([]byte, len(input))
	for i, _ := range input {
		res[i] = input[i] ^ fixed[i]
	}
	
	return res
}
