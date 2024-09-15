package main

import "fmt"

type Set []int

func (s *Set) add(value int) {
	for _, i := range *s {
		if i == value {
			return
		}
	}
	*s = append(*s, value)
}

func createTree(tree *map[int]map[int]Set, edgeList [][]int) {
	for _, val := range edgeList {
		if v, ok := (*tree)[val[0]]; ok {
			if v1, ok1 := v[val[1]]; ok1 {
				v1.add(val[2])
				(*tree)[val[0]][val[1]] = v1
			} else {
				(*tree)[val[0]][val[1]] = Set{val[2]}
			}
		} else {
			(*tree)[val[0]] = map[int]Set{val[1]: {val[2]}}
		}

		if v, ok := (*tree)[val[1]]; ok {
			if v1, ok1 := v[val[0]]; ok1 {
				v1.add(val[2])
				(*tree)[val[1]][val[0]] = v1
			} else {
				(*tree)[val[1]][val[0]] = Set{val[2]}
			}
		} else {
			(*tree)[val[1]] = map[int]Set{val[0]: {val[2]}}
		}
	}
}

func (s *Set) minSet() int {
	res := (*s)[0]
	for i := 1; i < len(*s); i++ {
		if (*s)[i] < res {
			res = (*s)[i]
		}
	}
	return res
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	n = len(queries)
	result := make([]bool, n)
	tree := make(map[int]map[int]Set)
	createTree(&tree, edgeList)

	for i, value := range queries {
		if d := tree[value[0]][value[1]]; d.minSet() < value[2] {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	fmt.Println(tree)
	return result
}

func main() {
	fmt.Println(distanceLimitedPathsExist(3, [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}, [][]int{{0, 1, 2}, {0, 2, 5}}))
}
