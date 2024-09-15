package main

import "fmt"

func distanceTraveled(mainTank int, additionalTank int) int {
	total := 0

	for mainTank > 0 {
		if mainTank < 5 {
			total += mainTank * 10
			break
		}
		total += 50
		mainTank -= 5

		if additionalTank > 0 {
			additionalTank--
			mainTank++
		}
	}
	return total
}

//func distanceTraveled(mainTank int, additionalTank int) int {
//	var path int
//	for mainTank > 0 {
//		mainTank -= 5
//		if mainTank >= 0 {
//			if additionalTank > 0 {
//				mainTank += 1
//				additionalTank--
//			}
//			path += 5 * 10
//		} else {
//			path += (5 + mainTank) * 10
//		}
//	}
//	return path
//}

func main() {
	fmt.Println(distanceTraveled(10, 3))
}
