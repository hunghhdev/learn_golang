package main

import "fmt"

func plusOne(digits []int) []int {
	last := len(digits) - 1
	for i := last; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{8, 9}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{9}))
}
