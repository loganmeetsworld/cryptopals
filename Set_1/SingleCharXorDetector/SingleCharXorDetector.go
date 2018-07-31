package main

import (
	"encoding/hex"
	"io/ioutil"
	"strings"
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

func singleByteXor(encodedMessage []byte, charMap map[rune]float64) ([]byte, float64) {
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
	return result, lastScore
}

func hexDecode(s string) []byte {
	result, err := hex.DecodeString(s)
	if err != nil {
		panic("error!")
	}

	return result
}

func main() {
	randomEnglish := readEnglishDataFromFile()
	englishRuneMap := gatherEnglishData(randomEnglish)

	stringBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Failed to open file")
	}
	text := string(stringBytes)
	var lastScore float64
	var result []byte
	for _, line := range strings.Split(string(text), "\n") {
		line := hexDecode(line)
		output, score := singleByteXor(line, englishRuneMap)
		if score > lastScore {
			result = output
			lastScore = score
		}
	}

	println(string(result))
}
