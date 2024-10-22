package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, str := range strs {
		bytesStr := []byte(str)
		sort.Slice(bytesStr, func(i, j int) bool {
			return bytesStr[i] < bytesStr[j]
		})
		hash[string(bytesStr)] = append(hash[string(bytesStr)], str)
	}
	res := make([][]string, 0)
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
