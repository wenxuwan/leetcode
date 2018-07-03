package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	var max int = 0x7fffffff
	var min int = -2147483648
	if x < min {
		fmt.Println(max, min)
		return 0
	}
	str1 := strconv.Itoa(x)
	var negative int = 0
	var result int = 0
	if str1[0] == '-' {
		negative = 1
	}
	for i := len(str1) - 1; i >= negative; i-- {
		result = result*10 + int(str1[i]-'0')
	}
	if negative == 1 {
		result = 0 - result
	}
	if result < min || result > max {
		return 0
	}
	return result
}

func main() {
	a := 12345
	x := reverse(a)
	fmt.Println(x)
}
