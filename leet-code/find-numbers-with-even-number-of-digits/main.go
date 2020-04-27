package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := [...]int{555, 901, 482, 1771}
	fmt.Println(findNumbers(nums[:]))
}

func findNumbers(nums []int) int {
	var result int
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			result++
		}
	}
	return result
}
