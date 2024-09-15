package main

import "fmt"

func emailIsNormal(email string) string {
	res := ""
	localName := true
	for i := 0; i < len(email); i++ {
		if email[i] == '+' && localName {
			for email[i] != '@' {
				i++
			}
			localName = false
		} else if email[i] == '.' && localName {
			continue
		} else if email[i] == '@' {
			localName = false
		}
		res += string(email[i])
	}
	return res
}

func numUniqueEmails(emails []string) int {
	hash := make(map[string]int)
	for _, email := range emails {
		hash[emailIsNormal(email)]++
	}
	return len(hash)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}))
}
