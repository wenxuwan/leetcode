/*
题目最简单的做法还是先排序，两个指针遍历数组，这样时间复杂度o(2n)，我看很多
人用map来存，可以达到n的时间复杂度，但空间就不是o(1)了
*/
package main

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, k int) int {
	result := 0
	begin := 0
	end := 1
	sort.Ints(nums)
	lastValue := -65535
	for end < len(nums){
		for begin < end {
			if end >= len(nums){
				return result
			}
			if nums[end] - nums[begin] == k{
				if lastValue != nums[end]{
					lastValue = nums[end]
					result++
				}
				end++
				begin++
			}else{
				if nums[end] - nums[begin] > k{
					begin++
				}else{
						end++
				}
			}
		}
		if begin == end{
			end++
		}
	}
	fmt.Println(nums,result)
	return result
}


func main(){
	data := []int{-1,-2,-3}
	k := 2
	findPairs(data, k)
	//1,1,3,4,5
}
