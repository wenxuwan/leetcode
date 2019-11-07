/*
窗口滑动

假设窗口内出现最多的是字母M
那么right - left + 1 -M > k的时候就证明现在已经到达最大处
然后left就要左移了。
*/
package main

import "fmt"

func characterReplacement(s string, k int) int{
	left := 0
	right := 0
	max := 0
	result := 0
	hash := make([]int,26)
	for right = 0;right < len(s);right++{
		hash[s[right] - 'A']++
		if hash[s[right] - 'A'] > max{
			max =  hash[s[right] - 'A']
		}
		if right - left + 1 - max <= k{
			if result < right - left + 1{
				result = right - left + 1
			}
		}else{
			hash[s[left] - 'A']--
			left++
		}
	}
	return result
}

func characterReplacement(s string, k int) int {
	mapHash := make(map[byte]int)
	maxLen := 0
	begin := 0
	end := 0
	var maxKey byte
	maxNumber := 0
	for end < len(s){
		mapHash[s[end]]++
		lenth := 0
		if mapHash[s[end]] >= maxNumber{
			maxKey = s[end]
			maxNumber = mapHash[s[end]]
		}
		for key, value := range mapHash{
			if key != maxKey{
				lenth += value
			}
		}
		//fmt.Println(begin, end, maxKey, maxNumber,lenth)
		if lenth == k + 1{
			fmt.Printf("%c,%c,%d,%d",s[begin],s[end],begin,end)
			fmt.Println()
			if maxLen < maxNumber + lenth - 1{
				fmt.Println(begin, end, maxKey - 'A', maxNumber,lenth)
				maxLen = maxNumber + lenth - 1
			}
			for s[begin] == maxKey{
				mapHash[s[begin]]--
				begin++
				maxNumber--
			}
			mapHash[s[begin]]--
			begin++
		}else{
			if maxLen < end - begin + 1{
				fmt.Println(begin,end)
				maxLen = end - begin + 1
			}
		}
		end++
	}
	return maxLen
}

func main(){
	s := "IMNJJTRMJEGMSOLSCCQICIHLQIOGBJAEHQOCRAJQMBIBATGLJDTBNCPIFRDLRIJHRABBJGQAOLIKRLHDRIGERENNMJSDSSMESSTR"
	//s := "ABBAAA"
	k := 2
	fmt.Println(characterReplacement(s,k))
}
