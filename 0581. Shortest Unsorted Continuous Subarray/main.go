/*
不考虑复杂度的话，最简单的做法就是排序，然后从两头查找
找到变化过的第一个数据。
*/
package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	tmp := make([]int,len(nums),len(nums))
	copy(tmp, nums)
	sort.Ints(nums)
	start := -1
	end := -1
	for i := 0;i < len(nums);i++{
		if tmp[i] != nums[i]{
			if start == -1{
				start = i
				break
			}
		}
	}
	for i := len(nums) - 1;i > 0;i--{
		if tmp[i] != nums[i]{
			if end == -1{
				end = i
				break
			}
		}
	}
	fmt.Println(start, end)
	if start == -1{
		return 0
	}
	return  end - start + 1
}
func main(){

}
