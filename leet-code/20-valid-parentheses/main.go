package main

func main() {
	println(isValid("([)]"))
}

func isValid(s string) bool {
	var open1 byte = '('
	var close1 byte = ')'
	var open2 byte = '['
	var close2 byte = ']'
	var open3 byte = '{'
	var close3 byte = '}'
	var tmp []byte
	for i := 0; i < len(s); i++ {
		if len(tmp) != 0 {
			n := len(tmp) - 1
			if (tmp[n] == open1 && s[i] == close1) || (tmp[n] == open2 && s[i] == close2) || (tmp[n] == open3 && s[i] == close3) {
				tmp = tmp[:n]
				continue
			}
		}
		tmp = append(tmp, s[i])
	}
	return len(tmp) == 0
}
