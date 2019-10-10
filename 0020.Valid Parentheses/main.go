/*
看匹配不匹配，就是要看在出现右括号的时候，前面的左括号能不能匹配，如果能匹配
就把左括号弹出，如果不能匹配就直接失败。
*/

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
