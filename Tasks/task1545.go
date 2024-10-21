package main

import (
	"fmt"
)

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	length := (1 << n) - 1 // Длина строки S_n (2^n - 1)
	mid := length/2 + 1    // Центральный элемент

	if k == mid {
		return '1' // Центральный элемент всегда '1'
	} else if k < mid {
		return findKthBit(n-1, k) // Ищем в левой части строки
	} else {
		// Ищем в правой части: отражённой, инвертированной версии S_(n-1)
		return invert(findKthBit(n-1, length-k+1))
	}
}

func invert(b byte) byte {
	// Инвертируем бит
	if b == '0' {
		return '1'
	}
	return '0'
}

func main() {
	fmt.Println(findKthBit(4, 11)) // 0
}
