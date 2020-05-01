package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	arr := strings.Split(strconv.Itoa(n), "")
	sum := 0
	multiplication := 1
	for _, num := range arr {
		n, _ := strconv.Atoi(num)
		sum += n
		multiplication *= n
	}
	return multiplication - sum
}
