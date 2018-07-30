package main

import (
	"encoding/hex"
	"fmt"
)

func fixedXOR(x string, y string) (string, error) {
	result, err := hex.DecodeString(x)

	return result, nil
}

func main() {
	x := "1c0111001f010100061a024b53535009181c"
	y := "686974207468652062756c6c277320657965"

	result, err := fixedXOR(x, y)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("output: ", result)
}
