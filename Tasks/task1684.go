package main

func countConsistentStrings(allowed string, words []string) int {
	hash := make(map[rune]int)
	for _, i := range allowed {
		hash[i] = 1
	}
	count := 0
	for _, word := range words {
		flag := true
		for _, ch := range word {
			if _, ok := hash[ch]; !ok {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}
	return count
}

func main() {

}
