package main

func Count[T comparable](slice []T, value T) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

func minSwaps(nums []int) int {
	n := len(nums)
	// Количество единиц в массиве
	countOnes := Count(nums, 1)
	if countOnes == 0 {
		return 0
	}

	// Дублируем массив для удобства циклической обработки
	nums = append(nums, nums...)

	// Считаем количество нулей в начальном окне
	countZeroes := Count(nums[:countOnes], 0)
	minZeroes := countZeroes

	// Используем метод скользящего окна для подсчета нулей
	for i := 1; i <= n; i++ {
		if nums[i-1] == 0 {
			countZeroes--
		}
		if nums[i+countOnes-1] == 0 {
			countZeroes++
		}
		if countZeroes < minZeroes {
			minZeroes = countZeroes
		}
	}

	return minZeroes
}
