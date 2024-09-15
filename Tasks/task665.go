package main

func checkPossibility(nums []int) bool {
	n := len(nums)
	isChanged := false
	for j := 0; j < 2; j++ {
		for i := 0; i < n-1; i++ {
			if nums[i] > nums[i+1] {
				if isChanged {
					return false
				}
				if i > 0 && nums[i-1] > nums[i+1] {
					nums[i+1] = nums[i]
				} else {
					nums[i] = nums[i+1]
				}

				isChanged = true
			}
		}
	}
	return true
}

func main() {

}
