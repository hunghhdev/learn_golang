package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	var result []bool
	for _, v := range candies {
		if v >= max {
			max = v
		}
	}
	for _, v := range candies {
		result = append(result, (v+extraCandies) >= max)
	}
	return result
}
func main() {
	candies := [...]int{4, 2, 1, 1, 2}
	fmt.Println(kidsWithCandies(candies[:], 1))
}
