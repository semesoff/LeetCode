package main

import "fmt"

func minOperations(logs []string) int {
	depth := 0
	for _, log := range logs {
		if log == "../" {
			if depth > 0 {
				depth--
			}
		} else if log == "./" {
			continue
		} else {
			depth++
		}
	}
	return depth
}

func main() {
	values := []string{"d1/", "../", "../", "../"}
	result := minOperations(values)
	fmt.Println(result)
}
