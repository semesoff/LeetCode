package main

import (
	"fmt"
	"math/bits"
)

//func minBitFlips(start int, goal int) int {
//	strStart := strconv.FormatInt(int64(start), 2)
//	strGoal := strconv.FormatInt(int64(goal), 2)
//	count := 0
//	if len(strStart) < len(strGoal) {
//		strStart = strings.Repeat("0", len(strGoal)-len(strStart)) + strStart
//	} else {
//		strGoal = strings.Repeat("0", len(strStart)-len(strGoal)) + strGoal
//	}
//	for i := range strStart {
//		if strStart[i] != strGoal[i] {
//			count++
//		}
//	}
//	return count
//}

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}

func main() {
	fmt.Println(minBitFlips(10, 7))
}
