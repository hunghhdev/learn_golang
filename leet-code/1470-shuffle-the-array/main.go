package main

import "fmt"

func shuffle(nums []int, n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[n+i])
	}
	return result
}
func main() {
	nums := [...]int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(shuffle(nums[:], 4))
}
