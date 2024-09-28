package main

import "fmt"

type MyCalendarTwo struct {
	dates [][2]int
	hash  map[[2]int]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{dates: [][2]int{}, hash: map[[2]int]int{}}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	borders := make([][2]int, 0)
	for _, date := range this.dates {
		if date[0] < end && start < date[1] {
			borders = append(borders, date)
		}
	}
	n := len(borders)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if borders[i][0] < borders[j][1] && borders[j][0] < borders[i][1] {
				return false
			}
		}
	}
	this.dates = append(this.dates, [2]int{start, end})
	return true
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Book(10, 20))
	fmt.Println(obj.Book(50, 60))
	fmt.Println(obj.Book(10, 40))
	fmt.Println(obj.Book(5, 15))
	fmt.Println(obj.Book(5, 10))
	fmt.Println(obj.Book(25, 55))
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
