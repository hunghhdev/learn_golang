package main

import "fmt"

func main() {
	nums := [...]int{1,2,3,4}
	fmt.Println(decompressRLElist(nums[:]))
}

func decompressRLElist(nums []int) []int {
	var s []int
    for i := 0;  i<len(nums); i=i+2 {
		fmt.Println(i)
		for j := 0; j<nums[i]; j++{
			fmt.Println(j)
			s = append(s, nums[i+1])
		}
	}
	return s
}
