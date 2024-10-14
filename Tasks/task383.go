package main

func canConstruct(ransomNote string, magazine string) bool {
	count := make([]int, 26)
	for _, c := range magazine {
		count[c-'a']++
	}
	for _, c := range ransomNote {
		if count[c-'a'] == 0 {
			return false
		}
		count[c-'a']--
	}
	return true
}

func main() {

}
