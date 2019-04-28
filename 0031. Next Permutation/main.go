package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int)  {
	lenth := len(nums)
	flag := 0
	j := 0
	for i := lenth - 2; i >= 0;i--{
		if nums[i] < nums[i + 1]{
			for j = i + 1;j < lenth;j++{
				if nums[j] > nums[i]{
					continue
				}else{
					break
				}
			}
			nums[j - 1],nums[i] = nums[i], nums[j - 1]
			flag = i + 1
			break
		}
	}
	sort.Ints(nums[flag:])
}

func main(){
	s := []int{1}
	nextPermutation(s)
	fmt.Println(s)
}
