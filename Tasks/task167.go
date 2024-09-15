package main

func twoSum(nodes []int, k int) []int {
	left, right := 0, len(nodes)-1
	for left < right {
		if nodes[left]+nodes[right] == k {
			return []int{left + 1, right + 1}
		} else if nodes[left]+nodes[right] < k {
			left++
		} else {
			right--
		}

	}
	return []int{}
}

func main() {

}
