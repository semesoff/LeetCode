package main

import "fmt"

func findTheLongestSubstring(s string) int {
	vowelMask := map[rune]int{
		'a': 1 << 0, // 00001
		'e': 1 << 1, // 00010
		'i': 1 << 2, // 00100
		'o': 1 << 3, // 01000
		'u': 1 << 4, // 10000
	}

	mask := 0
	maskMap := map[int]int{0: -1}
	maxLength := 0

	for i, ch := range s {
		if val, ok := vowelMask[ch]; ok {
			mask ^= val
		}
		if pos, ok := maskMap[mask]; ok {
			maxLength = max(maxLength, i-pos)
		} else {
			maskMap[mask] = i
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
}
