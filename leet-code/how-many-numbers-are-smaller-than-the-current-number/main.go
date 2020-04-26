package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [...]int{1, 2, 3, 4}
	fmt.Println(smallerNumbersThanCurrent(nums[:]))
}

func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	counts := make([]int, len(nums))
	for i, num := range nums {
		counts[i] = sort.Search(len(sortedNums), func(j int) bool {
			return num <= sortedNums[j]
		})
	}
	return counts
}
