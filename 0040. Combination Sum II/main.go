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
		//fmt.Println(result)
		*result = append(*result, subarray)
		return
	}
	for i := start_index; i < len(candidate) && candidate[i] <= target;i++{
		tmp := make([]int, len(subarray))
		copy(tmp,subarray)
		tmp = append(tmp,candidate[i])
		findSubArray(target - candidate[i], i + 1, candidate ,tmp,result)
		next := i + 1
		for next < len(candidate) && candidate[next] == candidate[i]{
			next++
		}
		i = next - 1
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	subarray := []int{}
	findSubArray(target,0, candidates, subarray, &result)
	return result
}

func main(){
	candidate := []int{10,1,2,7,6,1,5}
	fmt.Println(combinationSum2(candidate,8))
}