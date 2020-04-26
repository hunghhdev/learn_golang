package main

import "fmt"

func main() {
	nums := [...]int{1, 2, 3, 4}
	fmt.Println(smallerNumbersThanCurrent(nums[:]))
}

func smallerNumbersThanCurrent(nums []int) []int {
	var arr []int
	for _, num := range nums {
		arr = append(arr, countNumberSmall(nums, num))
	}
	return arr
}

func countNumberSmall(nums []int, num int) int {
	var count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < num {
			count++
		}
	}
	return count
}
