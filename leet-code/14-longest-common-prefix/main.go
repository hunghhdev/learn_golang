package main

import (
	"fmt"
)

func main() {
	arrays := [...]string{"dog", "racecar", "car"}
	// arrays := [...]string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(arrays[:]))
}

func longestCommonPrefix(strs []string) string {
	str := strs[0]
	for i := 1; i < len(strs); i++ {
		str = compare(str, strs[i])
		if len(str) == 0 {
			return ""
		}
	}
	return str
}

func compare(str1 string, str2 string) string {
	length := len(str1)
	if length > len(str2) {
		length = len(str2)
	}
	str := ""
	for i := 0; i < length; i++ {
		if str1[i] == str2[i] {
			str = str1[0 : i+1]
			continue
		}
		return str
	}
	return str
}
