/*
该题目主要是还是要回溯，因为可以无限制的使用数字，所以每次找下一个的
时候index都要从此次的index开始。
*/
package main

import (
	"sort"
	"fmt"
)

func findSubArray(target int , start_index int, candidate []int, subarray []int, result *[][]int){
	if target < 0 {
		return
	}
	if target == 0{
		*result = append(*result, subarray)
		return
	}
	for i := start_index; i < len(candidate) && candidate[i] <= target;i++{
		tmp := make([]int, len(subarray))
		copy(tmp,subarray)
		tmp = append(tmp,candidate[i])
		findSubArray(target - candidate[i], i, candidate ,tmp,result)
	}
}
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	subarray := []int{}
	findSubArray(target,0, candidates, subarray, &result)
	return result
}

func main(){
	candidate := []int{7,3,2}
	fmt.Println(combinationSum(candidate,18))
}
