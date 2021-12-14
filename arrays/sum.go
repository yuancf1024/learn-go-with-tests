package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)

	var sums []int

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		// 使用 append 函数，它能为切片追加一个新值。
		sums = append(sums, Sum(numbers))
	}

	return sums
}