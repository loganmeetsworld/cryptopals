package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func fixedXOR(x, y string) []byte {
	decodedX := hexDecode(x)
	decodedY := hexDecode(y)

	if len(decodedX) != len(decodedY) {
		panic("xor: mismatched lengths") // built-in to the language
	}

	result := make([]byte, len(decodedX))
	for index := range decodedX {
		result[index] = decodedX[index] ^ decodedY[index]
	}
	log.Printf("%s", decodedX) // printing as a string, the the raw byte value of x
	log.Printf("%s", decodedY) // printing as a string, the the raw byte value of y
	return result
}

func hexDecode(s string) []byte {
	result, err := hex.DecodeString(s)
	if err != nil {
		panic("error!")
	}

	return result
}

func main() {
	x := "1c0111001f010100061a024b53535009181c"
	y := "686974207468652062756c6c277320657965"

	result := fixedXOR(x, y)
	fmt.Println("output: ", result)
}
