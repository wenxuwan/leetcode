package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	begin := 0
	end := len(s) - 1
	for begin < end{
		runes[begin], runes[end] = runes[end], runes[begin]
		begin++
		end--
	}
	return string(runes)
}

func main(){
	testString := "hello"
	fmt.Println(reverseString(testString))
}