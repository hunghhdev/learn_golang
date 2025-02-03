package main

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	numA, _ := new(big.Int).SetString(a, 2)
	numB, _ := new(big.Int).SetString(b, 2)
	sum := new(big.Int).Add(numA, numB)
	return sum.Text(2)
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1010", "1010"))
	fmt.Println(addBinary("1111", "1111"))
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}
