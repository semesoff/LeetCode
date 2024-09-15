package main

import (
	"fmt"
)

package main

import (
	"fmt"
)

func longestSubarray(nums []int) int {
	var data [2]int
	data[0] = nums[0]
	data[1] = 1

	maxNumber := nums[0]
	size := 1

	for _, num := range nums[1:] {
		if num > maxNumber {
			maxNumber = num
			size = 1
		} else if num == maxNumber {
			size++
		} else {
			size = 0
		}
		
		if maxNumber > data[0] {
			data[0] = maxNumber
			data[1] = size
		} else if maxNumber == data[0] {
			if size > data[1] {
				data[1] = size
			}
		}
	}

	return data[1]
}



func main() {
	fmt.Println(longestSubarray([]int{100, 5, 5, 105, 105, 0, 0, 105, 105, 105, 0, 105}))
}
