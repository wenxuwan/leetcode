package main

import "fmt"

func dfs(current int, result * [][]int, tmp []int, value []int, flag []bool){
	if current == len(value){
		s := make([]int, len(value))
		copy(s, tmp)
		*result = append(*result, s)
		return
	}
	for i := 0; i < len(value);i++{
		if flag[i] == false{
			flag[i] = true
			tmp[current] = value[i]
			dfs(current + 1,result,tmp,value,flag)
			flag[i] = false
		}
	}
}
func permute(nums []int) [][]int {
	var result [][]int
	flag := make([]bool,len(nums))
	tmp := make([]int, len(nums))
	dfs(0, &result, tmp, nums, flag)
	return result
}

func main(){
	a := []int{1,2,3}
	fmt.Println(permute(a))
}