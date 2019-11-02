package main

import (
	"fmt"
)

func helper(s string, start, end int)int{
	count := 0
	for start >= 0 && end < len(s) && (s[start] == s[end]){
		count++
		start--
		end++
	}
	return count
}
func countSubstrings(s string) int {
	result := 0
	for i := 0; i < len(s);i++{
		result += helper(s,i,i)//奇数对称
		result += helper(s,i,i+1)//偶数对称
	}
	return result
}

func main(){
	testData := "dnncbwoneinoplypwgbwktmvkoimcooyiwirgbxlcttgteqthcvyoueyftiwgwwxvxvg"
	fmt.Println(countSubstrings(testData))
}
