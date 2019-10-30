/*和438一样的解法，具体参照438*/
package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	appearCount := 0
	left := 0
	right := 0
	var result string
	maxLen := math.MaxInt32
	for _, value := range t{
		m[byte(value)]++
	}
	for right < len(s){
		if _,ok := m[byte(s[right])];ok{
			m[byte(s[right])]--
			if m[byte(s[right])] >= 0{
				appearCount++
			}
		}
		if appearCount == len(t){
			for appearCount == len(t) {
				if _,ok := m[byte(s[left])];ok{
					if m[byte(s[left])] < 0{
						m[byte(s[left])]++
						left++
					}else{
						fmt.Println(m)
						m[byte(s[left])]++
						appearCount--
						break
					}
				}else{
					left++
				}
			}
			fmt.Println(right,left,maxLen)
			if maxLen > right - left{
				result = s[left:right + 1]
				maxLen = right - left
			}
			left++
		}
		right++
	}
	return result
}

func main(){
	S := "ADOBECODEBANC"
	T := "ABC"
	fmt.Println(minWindow(S,T))
}
