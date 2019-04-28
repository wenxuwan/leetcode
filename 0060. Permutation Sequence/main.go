package main

import (
	"fmt"
	"strconv"
)

func dfs(current int, result * [][]string, tmp []string, value []int, flag []bool,index *int,k int){
	if current == len(value){
		//fmt.Println(tmp)
		*index = *index + 1
		if *index == k{
			s := make([]string, len(value))
			copy(s, tmp)
			*result = append(*result, s)
			return
		}
	}
	for i := 0; i < len(value);i++{
		if flag[i] == false{
			flag[i] = true
			tmp[current] = strconv.Itoa(value[i])
			//fmt.Println(tmp[current])
			if *index >= k{
				return
			}
			dfs(current + 1,result,tmp,value,flag, index, k)
			flag[i] = false
		}
	}
}
func permute(nums []int,k int) [][]string {
	var result [][]string
	index := 0
	flag := make([]bool,len(nums))
	tmp := make([]string, len(nums))
	dfs(0, &result, tmp, nums, flag, &index, k)
	return result
}

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++{
		nums[i] = i + 1
	}
	result := permute(nums,k)
	//fmt.Println(result)
	s := ""
	for _, c := range result[0]{
		s = s + c
	}
	return s
	//return string(result[k - 1])
}

func main(){
	//a := []int{1,2,3}
	fmt.Println(getPermutation(9,219601))
}