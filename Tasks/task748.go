package main

import (
	"fmt"
	"maps"
	"math"
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	hashLicense := make(map[rune]int)
	for _, i := range licensePlate {
		if unicode.IsLetter(i) {
			hashLicense[unicode.ToLower(i)]++
		}
	}
	shortestWord := ""
	size := math.MaxInt32
	for _, word := range words {
		hashWord := maps.Clone(hashLicense)
		for _, char := range word {
			if _, ok := hashWord[char]; ok {
				hashWord[char]--
				if hashWord[char] == 0 {
					delete(hashWord, char)
				}
			}
		}
		if len(hashWord) == 0 {
			if t := len(word); t < size {
				shortestWord = word
				size = t
			}
		}
	}
	return shortestWord
}

func main() {
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
}
