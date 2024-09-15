package main

import (
	"fmt"
	"sort"
	"strings"
)

func maxLength(data map[string]int) int {
	var m int
	for _, i := range data {
		if i > m {
			m = i
		}
	}
	return m
}

func getNames(data map[string]int, length int) []string {
	names := make([]string, 0)
	for i, j := range data {
		if j == length {
			names = append(names, i)
		}
	}
	return names
}

func largestWordCount(messages []string, senders []string) string {
	data := make(map[string]int)
	for idx, sender := range senders {
		length := len(strings.Split(messages[idx], " "))
		if _, ok := data[sender]; ok {
			data[sender] += length
		} else {
			data[sender] = length
		}
	}

	names := getNames(data, maxLength(data))
	sort.Strings(names)
	return names[len(names)-1]
}

func main() {
	fmt.Println(largestWordCount([]string{"How is leetcode for everyone", "Leetcode is useful for practice"}, []string{"Bob", "Charlie"}))
}
