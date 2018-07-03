package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var max int = 0x7fffffff
	m := x
	var result int = 0
	for x > 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result != m {
		return false
	}
	if result > max {
		return false
	}
	return true
}
func main() {
	x := -2147447412
	flag := isPalindrome(x)
	if flag {
		fmt.Println(x, "is Palindrome")
	}
}
