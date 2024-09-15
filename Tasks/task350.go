package main

func countNumbers(nums []int) map[int]int {
	hash := make(map[int]int)
	for _, i := range nums {
		hash[i]++
	}
	return hash
}

func result(hash map[int]int, nums []int) []int {
	res := make([]int, 0)
	for _, i := range nums {
		if val, ok := hash[i]; ok && val > 0 {
			hash[i]--
			res = append(res, i)
		}
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	n, m := len(nums1), len(nums2)
	if n < m {
		hash := countNumbers(nums1)
		return result(hash, nums2)
	}
	hash := countNumbers(nums2)
	return result(hash, nums1)
}

func main() {

}
