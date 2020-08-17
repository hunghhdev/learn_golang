package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				result++
			}
		}
	}
	return result
}
func main() {
	nums := [...]int{1, 2, 3}
	fmt.Println(numIdenticalPairs(nums[:]))
}
