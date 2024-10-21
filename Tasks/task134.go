package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	currentGain := 0
	totalGain := 0
	answer := 0

	for i := 0; i < len(gas); i++ {
		currentGain += gas[i] - cost[i]
		totalGain += gas[i] - cost[i]

		if currentGain < 0 {
			answer = i + 1
			currentGain = 0
		}
	}

	if totalGain < 0 {
		return -1
	} else {
		return answer
	}
}

func main() {
	fmt.Println(canCompleteCircuit([]int{2}, []int{2})) // 3
}
