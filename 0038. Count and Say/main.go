package main

import "fmt"

func helper(s string) string {
	begin := 0
	end := 0
	var result []byte
	for end = 1; end < len(s); end++ {
		if s[begin] != s[end] {
			result = append(result, byte(end-begin+'0'))
			result = append(result, s[begin])
			begin = end
		}
	}
	if begin < end {
		result = append(result, byte(end-begin+'0'))
		result = append(result, s[end-1])
	}
	return string(result)
}
func countAndSay(n int) string {
	var result string = "1"
	for i := 1; i < n; i++ {
		//fmt.Println(result)
		result = helper(result)
	}
	return result
}

func main() {
	fmt.Println(countAndSay(5))
}
