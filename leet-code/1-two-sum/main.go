package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			return []int{val, i}
		}
		m[target-nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := [...]int{2, 7, 11, 15}
	fmt.Println(twoSum(nums[:], 9))
	// nums := [...]int{3, 2, 4}
	// fmt.Println(twoSum(nums[:], 6))
	// nums := [...]int{3, 2, 3}
	// fmt.Println(twoSum(nums[:], 6))
}
