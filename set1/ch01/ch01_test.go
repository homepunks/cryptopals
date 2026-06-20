/* Convert hex to base 64*/

package ch01

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestChallenge01(t *testing.T) {
	var input string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hex, err := hex.DecodeString(input)
	if err != nil {
		fmt.Errorf("[ERROR] Could not decode hex from string %s: %v\n", input, err)
	}

	res := hexToBase64(hex)

	fmt.Println(string(res))
}

func hexToBase64(hex []byte) []byte {
	res := make([]byte, base64.StdEncoding.EncodedLen(len(hex)))
	base64.StdEncoding.Encode(res, hex)

	fmt.Println(base64.StdEncoding.EncodeToString(hex))
	return res
}
