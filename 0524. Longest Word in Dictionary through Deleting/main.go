/*
排序做好了就容易做了，先把字符串排序，然后遍历，然后去重即可

*/

package main

import (
	"fmt"
	"sort"
)

type Words []string

func (w Words) Len()int{
	return len(w)
}

func (s Words)Less(i, j int)bool{
	if s[i]< s[j] {
		if s[i][0] == s[j][0]{
			if len(s[i]) >= len(s[j]){
				return true
			}else{
				return false
			}
		}else{
			return true
		}

	} else {
		if s[i][0] == s[j][0]{
			if len(s[i]) > len(s[j]){
				return true
			}else{
				return false
			}
		}else{
			return false
		}
	}
}

func (s Words) Swap(i, j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func findLongestWord(s string, d []string) string {
	words := Words{}
	key := 'A'
	maxlen := 0
	haveFound := false
	result := ""
	for _,value := range d{
		words = append(words, value)
	}
	sort.Sort(words)
	for _, value := range words{
		index := 0
		begin := 0
		end := len(s)
		if key == int32(value[0]) && haveFound{
			continue
		}else{
			haveFound = false
			for begin < end && index < len(value){
				if value[index] == s[begin]{
					index++
				}
				begin++
			}
			if index == len(value){
				if maxlen < len(value){
					maxlen = len(value)
					result = value
				}
				key = int32(value[0])
				haveFound = true
			}
		}
	}
	//fmt.Println(words)
	return result
}

func main(){
	s := "abpcplea"
	d := []string{"ale","apple","applea","plea"}
	fmt.Println(findLongestWord(s, d))
}
