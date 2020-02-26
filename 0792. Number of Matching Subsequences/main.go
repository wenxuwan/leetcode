/*
这个题目就是个392版本的多输入变形，本质上没什么区别，但看有人先遍历words存到字典，的确会有提高时间
只能说测试用例中有挺多重复的words。

然后就是len函数耗时也蛮长的，尽量前面用一个变量记录长度，不是每次都len(s)或者len(S)
*/
package main

import "fmt"


func numMatchingSubseq(S string, words []string) int {
	result := 0
	i := 0
	j := 0
	lenS := len(S)
	wordsMap := make(map[string]int,len(words))
	for _, w := range words{
		wordsMap[w] += 1
	}
	for s, count := range wordsMap{
		i = 0
		j = 0
		l := len(s)
		for ;i < lenS && j < l;i++{
			if S[i] == s[j]{
				j++
			}
		}
		if j == l{
			fmt.Println(s)
			result += count
		}
	}
	fmt.Println(result)
	return result
}

func main(){
	S := "abcde"
	words := []string{"a", "bb", "acd", "ace"}
	numMatchingSubseq(S, words)
}
