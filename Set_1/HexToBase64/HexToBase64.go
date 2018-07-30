package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func hexToBase64(input string) (string, error) {
	result, err := hex.DecodeString(input)

	if err != nil {
		return "", err
	}

	log.Printf("%s", result) // printing as a string, the the raw byte value of the result
	return base64.StdEncoding.EncodeToString(result), nil
}

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	result, err := hexToBase64(input)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("output: ", result)
}
