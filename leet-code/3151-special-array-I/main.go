package main

import "fmt"

func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	lastIsOdd := nums[0]%2 == 1
	for i := 1; i < len(nums); i++ {
		isOdd := nums[i]%2 == 1
		if isOdd == lastIsOdd {
			return false
		}
		lastIsOdd = isOdd
	}
	return true
}

func main() {
	// nums := [...]int{2, 7, 11, 15}
	// nums := [...]int{1}
	nums := [...]int{4, 3, 2, 1, 6}
	fmt.Println(isArraySpecial(nums[:]))
}
