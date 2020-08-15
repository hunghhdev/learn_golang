package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}
	result, lastDigit := 0, 0
	arr := strings.Split(s, "")
	for _, v := range arr {
		tmp := romanMap[string(v)]
		if lastDigit < tmp {
			tmp -= 2 * lastDigit

		}
		lastDigit = tmp
		result += lastDigit
	}
	return result
}

func main() {
	fmt.Println(romanToInt("CD"))
}
