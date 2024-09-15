package main

import "sort"

//
//import "sort"
//
//type NumFreq struct {
//	Num  int
//	Freq int
//}
//
//type NumFreqs []NumFreq
//
//func (nf NumFreqs) Len() int {
//	return len(nf)
//}
//
//func (nf NumFreqs) Swap(i, j int) {
//	nf[i], nf[j] = nf[j], nf[i]
//}
//
//func (nf NumFreqs) Less(i, j int) bool {
//	// Сначала сортируем по частоте
//	if nf[i].Freq != nf[j].Freq {
//		return nf[i].Freq < nf[j].Freq
//	}
//	// Если частоты равны, сортируем по убыванию значения
//	return nf[i].Num > nf[j].Num
//}
//
//func frequencySort(nums []int) []int {
//	freq := make(map[int]int)
//	for _, num := range nums {
//		freq[num]++
//	}
//	var numFreqs NumFreqs
//	for num, freq := range freq {
//		numFreqs = append(numFreqs, NumFreq{num, freq})
//	}
//	sort.Sort(numFreqs)
//
//	var result []int
//	for _, num := range numFreqs {
//		for i := 0; i < num.Freq; i++ {
//			result = append(result, num.Num)
//		}
//	}
//	return result
//}

func frequencySort(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sort.SliceStable(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}

func main() {

}
