package main

import (
	"fmt"
	"strconv"
	"strings"
)

var d = map[int]string{
	0:          "Zero",
	1:          "One",
	2:          "Two",
	3:          "Three",
	4:          "Four",
	5:          "Five",
	6:          "Six",
	7:          "Seven",
	8:          "Eight",
	9:          "Nine",
	10:         "Ten",
	11:         "Eleven",
	12:         "Twelve",
	13:         "Thirteen",
	14:         "Fourteen",
	15:         "Fifteen",
	16:         "Sixteen",
	17:         "Seventeen",
	18:         "Eighteen",
	19:         "Nineteen",
	20:         "Twenty",
	30:         "Thirty",
	40:         "Forty",
	50:         "Fifty",
	60:         "Sixty",
	70:         "Seventy",
	80:         "Eighty",
	90:         "Ninety",
	100:        "Hundred",
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}

type Stack []string

func (s *Stack) add(value string) {
	*s = append(*s, value)
}

func (s *Stack) pop() string {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func ReverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func (s *Stack) lenStack() int {
	return len(*s)
}

func (s *Stack) popLeft() string {
	item := (*s)[0]
	*s = (*s)[1:]
	return item
}

func transformationThreeDigitNumber(num string) Stack {
	stack := Stack{}
	res := ""
	countZeroes := 0
	for i, j := range num {
		if j == '0' {
			countZeroes++
			continue
		}
		if countZeroes == 0 {
			key, _ := strconv.Atoi(string(j))
			res = strings.Join([]string{d[key], res}, "")
			stack.add(d[key])
		} else if countZeroes == 1 {
			if j == '1' {
				key, _ := strconv.Atoi(strings.Join([]string{string(j), string(num[i-1])}, ""))
				res = d[key]
				if !stack.isEmpty() {
					stack.pop()
				}
				stack.add(d[key])
			} else {
				key, _ := strconv.Atoi(strings.Join([]string{string(j), "0"}, ""))
				res = strings.Join([]string{d[key], res}, "")
				stack.add(d[key])
			}
		} else {
			numInt, _ := strconv.Atoi(strings.Join([]string{"1", strings.Repeat("0", countZeroes)}, ""))
			key, _ := strconv.Atoi(string(j))
			if _, err := d[numInt]; err {
				res = strings.Join([]string{d[key], d[numInt], res}, "")
				stack.add(d[numInt])
				stack.add(d[key])
			} else {
				cur := numInt
				multip := 1
				for cur != 1000 && cur != 1000000 && cur != 1000000000 {
					cur /= 10
					multip *= 10
				}
				res = strings.Join([]string{d[multip], d[cur], res}, "")
				stack.add(d[cur])
				stack.add(d[multip])
			}
		}
		countZeroes++
	}
	return stack
}

func getKey(q int, res *string) {
	if q == 1 {
		*res = d[1000] + " " + *res
	} else if q == 2 {
		*res = d[1000000] + " " + *res
	} else if q == 3 {
		*res = d[1e9] + " " + *res
	}
}

func numberToWords(num int) string {
	n := len(strconv.Itoa(num))
	cur := ""
	strNum := ReverseString(strconv.Itoa(num))
	q := -1
	res := ""
	stack := Stack{}
	if n < 4 {
		stack = transformationThreeDigitNumber(strNum)
		for !stack.isEmpty() {
			cur += stack.pop() + " "
		}
		res = cur
	} else {
		for i, j := range strNum {
			cur += string(j)
			if (i+1)%3 == 0 {
				q += 1

				stack = transformationThreeDigitNumber(cur)
				if !stack.isEmpty() {
					getKey(q, &res)
				}
				cur = ""
				for !stack.isEmpty() {
					cur += stack.pop() + " "
				}
				res = cur + res
				cur = ""
			}
		}

		if len(cur) > 0 {
			q += 1
			if len(cur) >= 2 && cur[len(cur)-1] > '1' {
				getKey(q, &res)
				p := 0
				for _, j := range cur {
					if p == 0 {
						if j != '0' {
							k, _ := strconv.Atoi(string(j))
							res = d[k] + " " + res
						}
					} else {
						k, _ := strconv.Atoi(string(j) + "0")
						res = d[k] + " " + res
					}
					p += 1
				}
			} else {
				k, _ := strconv.Atoi(ReverseString(cur))
				getKey(q, &res)
				res = d[k] + " " + res
			}

		}
	}
	res = strings.TrimSpace(res)
	if len(res) == 0 {
		return "Zero"
	}
	return res
}

func main() {
	fmt.Println(numberToWords(3000000))
}

// "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
