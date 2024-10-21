package main

import "fmt"

func isAnagram(s string, t string) bool {
	hash := make(map[rune]int)
	for _, v := range s {
		hash[v]++
	}
	for _, v := range t {
		if _, ok := hash[v]; ok {
			hash[v]--
			if hash[v] == 0 {
				delete(hash, v)
			}
		} else {
			return false
		}
	}
	return len(hash) == 0
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
}
