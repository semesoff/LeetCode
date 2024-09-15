package main

import "strconv"

func countSeniors(details []string) int {
	var total int

	for _, i := range details {
		if val, _ := strconv.Atoi(i[11:13]); val > 60 {
			total++
		}
	}
	return total
}

func main() {

}
