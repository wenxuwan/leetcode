package main

import "fmt"

func isValid(s string) bool {
	var maptable map[rune]int = map[rune]int{'(': 1, '[': 2, '{': 3, '}': 4, ']': 5, ')': 6}
	stack := make([]rune, len(s))
	var index int = 0
	for _, charvalue := range s {
		if maptable[charvalue] < 4 {
			stack[index] = charvalue
			index++
		} else {
			if index == 0 {
				return false
			}
			if maptable[rune(charvalue)]+maptable[stack[index-1]] != 7 {
				return false
			} else {
				index--
			}

		}
	}
	if index == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var str string = "]"
	fmt.Println(isValid(str))
}
