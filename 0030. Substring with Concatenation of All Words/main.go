/*
1.遍历s,从s中找到以words里面的单词开头的地方，保存到tmpWords
2.将words保存在hash表里
3.遍历得到的tmpWords，然后和hash表里的比对
*/
package main

import "reflect"

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0{
		return nil
	}
	wordLenth := len(words[0])
	tmpWords := make([]string,0)
	wordsHash := make(map[string]int)
	tmpResult := make([]int,0)
	result := make([]int,0)
	for _, value := range words{
		if wordsHash[value] != 0{
			wordsHash[value]++
			continue
		}
		for i := 0;i <= len(s) - (len(words) * wordLenth);i++{
			if s[i:i + wordLenth] == value{
				tmpWords = append(tmpWords,s[i:i + len(words) * wordLenth])
				tmpResult = append(tmpResult, i)
			}
		}
		wordsHash[value]++
	}
	//fmt.Println(tmpResult)
	//fmt.Println(tmpWords)
	for i := 0;i < len(tmpWords);i++{
		tmpHash := make(map[string]int)
		for j := 0; j <= len(tmpWords[i]) - wordLenth;j = j + wordLenth{
			s := tmpWords[i][j:j + wordLenth]
			tmpHash[s]++
		}
		//fmt.Println(tmpHash,wordsHash)
		/*reflect包里面的反射并没有比自己的慢呀
		for key, value := range(tmpHash){
			if _,ok := wordsHash[key];ok{
				if wordsHash[key] != value{
					flag = false
					break
				}
			}else{
				flag = false
				break
			}
		}
		if flag{
			result = append(result,tmpResult[i])
		}*/
		if reflect.DeepEqual(tmpHash, wordsHash){
			result = append(result,tmpResult[i])
		}
	}
	//fmt.Println(result)
	return result
}

func main(){
	words := []string{"fooo","barr","wing","ding","wing"}
	s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	findSubstring(s, words)
}
