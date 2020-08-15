package main

import "fmt"

func isPalindrome(x int) bool {
	return x == reverse(x)
}

func reverse(x int) int {
	var result int
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	return result
}

func main() {
	fmt.Println(isPalindrome(-121))
}
