package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func countOfAtoms(formula string) string {
	n := len(formula)
	stack := []map[string]int{{}}
	index := 0

	for index < n {
		if formula[index] == '(' {
			stack = append(stack, map[string]int{})
			index++
		} else if formula[index] == ')' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			index++
			start := index
			for index < n && unicode.IsDigit(rune(formula[index])) {
				index++
			}
			multiplier, err := strconv.Atoi(formula[start:index])
			if err != nil {
				multiplier = 1
			}
			for elem, count := range top {
				def := 0
				if stack[len(stack)-1][elem] > 0 {
					def = stack[len(stack)-1][elem]
				}
				stack[len(stack)-1][elem] = def + count*multiplier
			}
		} else {
			start := index
			index++
			for index < n && unicode.IsLower(rune(formula[index])) {
				index++
			}
			elem := formula[start:index]

			start = index
			for index < n && unicode.IsDigit(rune(formula[index])) {
				index++
			}
			count, err := strconv.Atoi(formula[start:index])
			if err != nil {
				count = 1
			}
			def := 0
			if stack[len(stack)-1][elem] > 0 {
				def = stack[len(stack)-1][elem]
			}
			stack[len(stack)-1][elem] = def + count
		}

	}

	counts := stack[len(stack)-1]
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var result string
	for _, elem := range keys {
		if counts[elem] > 1 {
			result += fmt.Sprintf("%s%d", elem, counts[elem])
		} else {
			result += fmt.Sprintf("%s", elem)
		}
	}
	return result
}

func main() {
	fmt.Println(countOfAtoms("Mg(OH)2"))
}
