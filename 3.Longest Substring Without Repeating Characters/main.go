package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	currentLen := 0
	currentIndex := 0
	j := 0
	for i := 0;i < len(s);i++{
		for j = currentIndex; j < i; j++ {
			if s[j] == s[i]{
				currentIndex = j + 1
				currentLen = i - j
				if currentIndex == i{
					currentLen = 1
				}
				break
			}
		}
		if j == i{
			currentLen += 1
		}
		if currentLen > maxLen{
			maxLen = currentLen
		}
	}
	return maxLen
}

func main(){
	//s := "abcabcbb"
	s1 := "ckilbkd"
	//len := lengthOfLongestSubstring(s)
	len1 := lengthOfLongestSubstring(s1)
	//len2 := lengthOfLongestSubstring("pwwkew")
	//len3 := lengthOfLongestSubstring("a")
	//len4 := lengthOfLongestSubstring("dvdf")
	//fmt.Println(len,len1,len2,len3,len4)
	fmt.Println(len1)
}
