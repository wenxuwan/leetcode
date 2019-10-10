/*
递归可以解决，停止条件有两个
1.如果左括号和右括号剩余的个数都为0证明我们匹配完了就return
2.第二种就是左括号的数目大于右括号的数目，这种情况下就证明出错了
所以直接return，其它时候随机取一个。
 */
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
