package main

import "fmt"

func dfs(result* [][]int, n int, k int, current int, tmp []int, start int){
	if current == k{
		tmpSlice := make([]int,k)
		copy(tmpSlice, tmp)
		*result = append(*result, tmpSlice)
		return
	}
	for i := start ; i < n;i++{
		tmp[current] = i + 1
		dfs(result, n ,k, current + 1,tmp,i + 1)
	}
}
func combine(n int, k int) [][]int {
	var result [][]int
	tmp := make([]int,k)
	dfs(&result,n,k,0,tmp,0)
	return result
}

func main(){
	fmt.Println(combine(4,2))
}