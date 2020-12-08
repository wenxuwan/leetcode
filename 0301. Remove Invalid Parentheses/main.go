package main

import (
	"fmt"
)

func isValid(s string) bool {
	left := 0
	for _, v := range []byte(s) {
		if v == '(' {
			left++
		} else if v == ')' {
			left--
		}
		if left < 0 {
			return false
		}
	}
	return left == 0
}

func getStringList(s string) []string {
	l := len(s)
	result := make([]string, 0)
	resultMap := make(map[string]bool)
	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == ')' {
			tmp := s[0:i] + s[i+1:l]
			if _, ok := resultMap[tmp]; !ok {
				resultMap[tmp] = true
				result = append(result, tmp)
			}
		}
	}
	return result
}

func removeInvalidParentheses(s string) []string {
	result := make([]string, 0)
	resMap := make(map[string]bool)
	if isValid(s) == true {
		result = append(result, s)
		return result
	}
	tmpList := getStringList(s)
	for len(tmpList) != 0 {
		nextList := make([]string, 0)
		success := false
		for _, value := range tmpList {
			if isValid(value) {
				if _, ok := resMap[value]; !ok {
					resMap[value] = true
					result = append(result, value)
					success = true
				}
			}
			if success != true {
				nextList = append(nextList, getStringList(value)...)
			}
		}
		if success {
			return result
		}
		tmpList = nextList
	}

	return result
}

func main() {
	fmt.Println(removeInvalidParentheses(")()(e()))))))(((("))
	fmt.Println(getStringList("((())"))

}
