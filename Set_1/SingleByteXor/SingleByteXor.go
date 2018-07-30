package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"unicode/utf8"
)

func readEnglishDataFromFile() string {
	randomEnglishBytes, err := ioutil.ReadFile("random_english.txt")
	if err != nil {
		panic("Failed to open file")
	}
	randomEnglish := string(randomEnglishBytes)
	return randomEnglish
}

func gatherEnglishData(data string) map[rune]float64 {
	charMap := make(map[rune]float64)
	for _, char := range data {
		charMap[char]++
	}

	total := utf8.RuneCountInString(data)
	for char := range charMap {
		charMap[char] = charMap[char] / float64(total)
	}

	return charMap
}

func scoreEnglishData(data string, charMap map[rune]float64) float64 {
	var score float64
	for _, char := range data {
		score += charMap[char]
	}

	return score / float64(utf8.RuneCountInString(data))
}

func singleXor(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i, c := range input {
		result[i] = c ^ key
	}
	return result
}

func singleByteXor(encodedMessage []byte, charMap map[rune]float64) []byte {
	var result []byte
	var lastScore float64
	for key := 0; key < 256; key++ {
		output := singleXor(encodedMessage, byte(key))
		score := scoreEnglishData(string(output), charMap)
		if score > lastScore {
			result = output
			lastScore = score
		}
	}
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
	message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	randomEnglish := readEnglishDataFromFile()
	englishRuneMap := gatherEnglishData(randomEnglish)
	answer := singleByteXor(hexDecode(message), englishRuneMap)

	fmt.Println("output: ", string(answer))
}
