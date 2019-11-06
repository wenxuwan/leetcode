package main

import "fmt"

func reverseVowels(s string) string {
	begin := 0
	end := len(s) - 1
	left := 0
	right := 0
	var result []byte = []byte(s)
	for begin < end{
		if s[begin] == 'a' || s[begin] == 'e' || s[begin] == 'i' || s[begin] == 'o' || s[begin] == 'u' || s[begin] == 'A' || s[begin] == 'E' || s[begin] == 'I' || s[begin] == 'O' || s[begin] == 'U' {
			left = 1
		}else{
			left = 0
			begin++
		}
		if s[end] == 'a' || s[end] == 'e' || s[end] == 'i' || s[end] == 'o' || s[end] == 'u' || s[end] == 'A' || s[end] == 'E' || s[end] == 'I' || s[end] == 'O' || s[end] == 'U' {
			right = 1
		}else{
			right = 0
			end--
		}
		if begin < end && right == 1 && left == 1{
			result[begin], result[end] = s[end],s[begin]
			end--
			begin++
			right = 0
			left = 0
		}
	}
	return string(result)
}

func main(){
	fmt.Println(reverseVowels("leetcode"))
}
