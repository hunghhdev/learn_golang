package main

import "fmt"

func isHappy(n int) bool{
	var x, y int = n, n
	for true{
		x = numSquareSum(x)
		y = numSquareSum(numSquareSum(y))
		if x != y{
			continue
		} else {
			break
		}
	}
	return x==1
}

func numSquareSum(n int) int{
    squareSum := 0; 
    for n != 0{
		squareSum += (n % 10) * (n % 10); 
        n = n / 10;
	} 
    return squareSum; 
}

func main()  {
	n := 19
	if isHappy(n) {
		fmt.Printf("%d is happy number", n)
	}
}
