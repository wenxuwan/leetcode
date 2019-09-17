/*

*/
package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	result := make([]bool,len(s) + 1,len(s) + 1)
	//fmt.Println(result)
	result[0] = true
	for i := 1;i < len(s) + 1;i++ {
		for j := 0; j < i; j++ {
			if result[j] == true {
				tmp := s[j:i]
				//fmt.Println(j, i, tmp)
				for _, word := range wordDict{
					if len(word) == len(tmp){
						count := 0
						for index, value := range word{
							if int32(tmp[index]) != value{
								break
							}
							count++
						}
						if count == len(word){
							result[i] = true
							break
						}
					}
				}
			}
		}
	}
	return result[len(s)]
}

func main(){
	s := "catssanddog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
