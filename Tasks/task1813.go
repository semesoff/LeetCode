package main

import (
	"fmt"
	"slices"
	"strings"
)

func countingSimilarWords(a, b []string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			count++
		} else {
			break
		}
	}
	return count
}

func isSimilar(a, b []string) bool {
	if len(a) > len(b) {
		a, b = b, a
	}
	prefix := countingSimilarWords(a, b)
	slices.Reverse(a)
	slices.Reverse(b)
	suffix := countingSimilarWords(a, b)
	return len(a)-(prefix+suffix) <= 0
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if sentence1 == sentence2 {
		return true
	}
	array1 := strings.Split(sentence1, " ")
	array2 := strings.Split(sentence2, " ")
	return isSimilar(array1, array2)
}

func main() {
	fmt.Println(areSentencesSimilar("a", "a aa a"))
}
