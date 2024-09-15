package main

import "fmt"

func isUnique(nums []int, m int) (bool, int) {
	uniques := map[int]int{}
	sum := 0
	for _, j := range nums {
		uniques[j] = 1
		sum += j
	}
	return len(uniques) >= m, sum
}

func createMap(uniques *map[int]int, nums []int) {
	for _, j := range nums {
		(*uniques)[j] += 1
	}
}

func maxSum(nums []int, m int, k int) int64 {
	uniques := map[int]int{}
	createMap(&uniques, nums[:k])
	maximum := 0
	curSum := 0
	ok, sum := isUnique(nums[:k], m)
	if ok {
		maximum = max(maximum, sum)
	}
	curSum = sum - nums[0]

	uniques[nums[0]] -= 1
	if uniques[nums[0]] == 0 {
		delete(uniques, nums[0])
	}
	n := len(nums)

	for i := 1; i < n; i++ {
		if i+k > n {
			break
		}

		uniques[nums[i+k-1]] += 1
		curSum += nums[i+k-1]

		if len(uniques) >= m {
			maximum = max(maximum, curSum)
		}

		uniques[nums[i]] -= 1
		if uniques[nums[i]] == 0 {
			delete(uniques, nums[i])
		}
		curSum -= nums[i]
	}
	return int64(maximum)
}

func main() {
	fmt.Println(maxSum([]int{1, 1, 1, 3}, 2, 2))
}
