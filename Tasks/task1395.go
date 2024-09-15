package main

import "fmt"

// TLE
//func numTeams(rating []int) int {
//	var total int
//
//	for i := 0; i < len(rating)-2; i++ {
//		for j := i + 1; j < len(rating)-1; j++ {
//			for k := j + 1; k < len(rating); k++ {
//				if (rating[i] < rating[j] && rating[j] < rating[k]) ||
//					(rating[i] > rating[j] && rating[j] > rating[k]) {
//					total++
//				}
//			}
//		}
//	}
//	return total
//}

func numTeams(rating []int) int {
	n := len(rating)
	result := 0

	for j := 1; j < n; j++ {
		left_smaller, left_larger, right_smaller, right_larger := 0, 0, 0, 0

		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				left_smaller++
			} else {
				left_larger++
			}
		}

		for k := j + 1; k < n; k++ {
			if rating[k] < rating[j] {
				right_smaller++
			} else {
				right_larger++
			}
		}
		result += left_smaller*right_larger + left_larger*right_smaller
	}

	return result
}

func main() {
	fmt.Println(numTeams([]int{2, 5, 3, 4, 1}))
}
