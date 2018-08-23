package main

import (
	"fmt"
)


func firstUniqChar(s string) int {
	runes := []rune(s)
	var countMap map[rune]int
	countMap = make(map[rune]int)
	for i := 0;i < len(runes);i++{
		countMap[runes[i]] = countMap[runes[i]] + 1
	}
	for i := 0;i < len(runes);i++{
		if countMap[runes[i]] == 1{
			return  i
		}
	}
	return  -1
}

func main(){
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}