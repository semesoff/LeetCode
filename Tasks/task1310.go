package main

func xorQueries(arr []int, queries [][]int) []int {
	counts := make([]int, 0)
	for _, val := range queries {
		res := 0
		for i := val[0]; i <= val[1]; i++ {
			if i == val[0] {
				res = arr[i]
			} else {
				res ^= arr[i]
			}
		}
		counts = append(counts, res)
	}
	return counts
}

func main() {

}
