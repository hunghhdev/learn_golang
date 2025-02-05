package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	x, y := 1, 2

	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}

	return y
}

func main() {
	fmt.Println(climbStairs(2))  // 2
	fmt.Println(climbStairs(3))  // 3
	fmt.Println(climbStairs(4))  // 5
	fmt.Println(climbStairs(5))  // 8
	fmt.Println(climbStairs(45)) // 8
}
