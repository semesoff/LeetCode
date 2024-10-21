package main

import "fmt"

func maxUniqueSplit(s string) int {
	n := len(s)
	hash := make(map[string]int)
	var backtracking func(string, int) int
	backtracking = func(str string, index int) int {
		if index == n {
			return len(hash)
		}
		res := 0
		for i := index; i < n; i++ {
			subStr := s[index : i+1]
			if _, ok := hash[subStr]; !ok {
				hash[subStr] = 1
				res = max(res, backtracking(subStr, i+1))
				delete(hash, subStr)
			}
		}
		return res
	}
	return backtracking("", 0)
}

func main() {
	fmt.Println(maxUniqueSplit("ababccc"))
}
