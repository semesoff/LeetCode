package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	Val  rune
	Next [26]*Trie
}

func longestCommonPrefix(strs []string) string {
	trie := &Trie{}
	countPrefix := len(strs[0])
	previousPrefix := strs[0]
	for i := 0; i < len(strs); i++ {
		word := strs[i]
		node := trie
		isAdd := true
		prefix := strings.Builder{}
		count := 0
		for _, char := range word {
			if node.Next[char-'a'] == nil {
				node.Next[char-'a'] = &Trie{char, [26]*Trie{}}
				node = node.Next[char-'a']
				isAdd = false
			} else if isAdd {
				node = node.Next[char-'a']
				if count == countPrefix {
					break
				}
				prefix.WriteRune(char)
				count++
			}
		}

		if i != 0 && count < countPrefix {
			countPrefix = count
			previousPrefix = prefix.String()
		} else if i != 0 && previousPrefix != prefix.String() {
			return ""
		}
	}
	return previousPrefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
