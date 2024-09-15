package main

func checkSubarraySum(nums []int, k int) bool {
	remainderMap := map[int]int{0: -1}
	cumulativeSum := 0

	for i, num := range nums {
		cumulativeSum += num
		remainder := cumulativeSum % k

		if _, ok := remainderMap[remainder]; ok {
			if i-remainderMap[remainder] > 1 {
				return true
			}
		} else {
			remainderMap[remainder] = i
		}
	}

	return false
}
