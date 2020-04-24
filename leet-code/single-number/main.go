package main

import "fmt"

func main() {
	arrays := [...]int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arrays[:]))
}

func singleNumber(nums []int) int {
	var result = 0
	for i := 0; i < len(nums); i++ {
		if onlyNumber(nums, nums[i]) {
			return nums[i]
		}
	}
	return result
}

func onlyNumber(nums []int, num int) bool {
	var count = 0
	for i := 0; i < len(nums); i++ {
		if num == nums[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
