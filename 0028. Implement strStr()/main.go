package main

import "fmt"

func kmp(subString string)[]int{
	p := 0
	k := -1
	nextArr := make([]int,len(subString) + 1,len(subString) + 1)
	nextArr[0] = -1
	for p < len(subString){
		if k == -1 || subString[p] == subString[k]{
			k++
			p++
			nextArr[p] = k
		}else {
			k = nextArr[k]
		}
	}
	return  nextArr
}
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0{
		return 0
	}
	if len(haystack) == 0{
		return -1
	}
	if len(needle) == 0{
		return  0
	}
	next := kmp(needle)
	fmt.Println(next)
	i := -1
	j := -1
	for i < len(haystack) && j < len(needle){
		if j == -1 || haystack[i] == needle[j]{
			i++
			j++
		}else{
			j = next[j]
		}
	}
	if j != len(needle){
		return  -1
	}else{
		return  i - j
	}
}

func main(){
	str1 := "mississippi"
	str2 := "issip"
	fmt.Println(strStr(str1,str2))
}