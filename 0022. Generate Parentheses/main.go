package main

import "fmt"

func selectSub(substring string, result *[]string ,left int ,right int){
	if left == 0 && right ==0{
		*result = append(*result, substring)
		return
	}
	if left > right{
		return
	}
	if left > 0{
		selectSub(substring + "(", result ,left - 1 ,right)
	}
	if right > 0{
		selectSub(substring + ")", result , left ,right - 1)
	}
}
func generateParenthesis(n int) []string{
	result := []string{""}
	selectSub("",&result,n,n)
	return result[1:]
}

func main(){
	fmt.Println(generateParenthesis(0))
}
