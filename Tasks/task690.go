package main

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	hash := make(map[int][2][]int)
	for _, i := range employees {
		hash[i.Id] = [2][]int{{i.Importance}, i.Subordinates}
	}
	var recursion func(idx int) int
	recursion = func(idx int) int {
		total := hash[idx][0][0]
		for _, i := range hash[idx][1] {
			total += recursion(i)
		}
		return total
	}
	return recursion(id)
}

func main() {
	t := make([]*Employee, 0)
	t = append(t, &Employee{1, 5, []int{2, 3}})
	t = append(t, &Employee{2, 3, []int{4}})
	t = append(t, &Employee{3, 4, []int{}})
	t = append(t, &Employee{4, 1, []int{}})
	fmt.Println(getImportance(t, 1))
}
