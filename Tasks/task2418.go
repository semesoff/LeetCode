package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name   string
	Height int
}

type ByHeight []Person

func (a ByHeight) Len() int           { return len(a) }
func (a ByHeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHeight) Less(i, j int) bool { return a[i].Height > a[j].Height }

func sortPeople(names []string, heights []int) []string {
	persons := make([]Person, len(heights))
	for i, name := range names {
		persons[i] = Person{name, heights[i]}
	}
	sort.Sort(ByHeight(persons))
	
	for i, person := range persons {
		names[i] = person.Name
		heights[i] = person.Height
	}

	return names
}

func main() {
	fmt.Println(sortPeople([]string{"Kat", "Bow"}, []int{4, 3}))
}
