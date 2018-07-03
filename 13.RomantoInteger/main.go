package main

import "fmt"

func romanToInt(s string) int {
	romanmap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result int = 0
	for index, value := range s {
		if index == 0 || romanmap[value] <= romanmap[rune(s[index-1])] {
			result += romanmap[value]
		} else {
			result = result - 2*romanmap[rune(s[index-1])] + romanmap[value]
		}
	}
	return result
}
func main() {
	var s string = "VIII"
	fmt.Println(romanToInt(s))
}
