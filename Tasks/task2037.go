package main

import (
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	sum := 0

	for i := range seats {
		sum += int(math.Abs(float64(students[i] - seats[i])))
	}

	return sum
}
