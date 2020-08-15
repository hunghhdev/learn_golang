package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	// for i := 0; i < len(nums); i++ {
	// 	for j := len(nums) - 1; j >= 0; j-- {
	// 		if i == j {
	// 			continue
	// 		}
	// 		if target == (nums[j] + nums[i]) {
	// 			result = append(result, i, j)
	// 			return result
	// 		}
	// 	}
	// }

	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			result = append(result, val, i)
			return result
		}
		m[target-nums[i]] = i
	}
	return result
}

func main() {
	// nums := [...]int{2, 7, 11, 15}
	// fmt.Println(twoSum(nums[:], 9))
	nums := [...]int{3, 2, 4}
	fmt.Println(twoSum(nums[:], 6))
	// nums := [...]int{3, 2, 3}
	// fmt.Println(twoSum(nums[:], 6))
}
