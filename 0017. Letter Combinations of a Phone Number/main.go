package main

import "fmt"

func letterCombinations(digits string) []string {
	m := map[int32]string{50:"abc",51:"def",52:"ghi",53:"jkl",54:"mno",55:"pqrs",56:"tuv",57:"wxyz"}
	var result []string
	lenth := len(result)
	flag := 0
	for index, val := range digits{
		//fmt.Println(result)
		lenth = len(result)
		if index == 0{
			for _, val := range m[val]{
				result = append(result, string(val))
			}
			continue
		}
		for flag < lenth{
			for _, val2 := range m[val]{
				tmp := result[flag] +  string(val2)
				result = append(result, tmp)
			}
			flag++
		}
	}
	//fmt.Println(flag)
	return result[flag:]
}
func main(){
	fmt.Println(letterCombinations("23"))
}
