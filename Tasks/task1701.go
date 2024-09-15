package main

func averageWaitingTime(customers [][]int) float64 {
	currentTime, average, n := 0, 0, 0

	for _, customer := range customers {
		arrival := customer[0]
		time := customer[1]
		n++
		if currentTime < arrival {
			currentTime = arrival
		}
		currentTime += time
		average += currentTime - arrival
	}
	return float64(average) / float64(n)
}
