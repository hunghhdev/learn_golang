package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 0, x/2+1
	for left <= right {
		mid := (left + right) / 2
		midSquare := mid * mid
		if midSquare == x {
			return mid
		}
		if midSquare < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if right*right == x {
			return right
		}
	}
	return right
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(25))
	fmt.Println(mySqrt(321))
	fmt.Println(mySqrt(7234))
	fmt.Println(mySqrt(63451))
}
