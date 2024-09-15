package main

import (
	"fmt"
	"strings"
)

func customSortString(order string, s string) string {
	hash := make(map[rune]int)
	res := strings.Builder{}
	for _, i := range s {
		hash[i]++
	}
	for _, i := range order {
		if _, ok := hash[i]; ok {
			w := strings.Repeat(string(i), hash[i])
			res.WriteString(w)
			delete(hash, i)
		}
	}
	for i, j := range hash {
		w := strings.Repeat(string(i), j)
		res.WriteString(w)
	}
	return res.String()
}

func main() {
	fmt.Println(customSortString("kqep", "pekeq"))
}
