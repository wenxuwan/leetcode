/*
求连续的，子序列想到的是用双指针，比如
10，5，2，6序列

一开始begin指向10，然后遍历数组，如果自身比k小而且begin到end之间
所有数的乘积小于k，那么可以组成的序列就是
end-begi+1， 1代表的是自身。
当自身大于k的时候end - begin变为-1，然后从下一轮开始。
*/
package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	preValue := 1
	begin := 0
	result := 0
	if k  <= 1{
		return 0
	}
	for end, value := range nums{
		preValue *= value
		for preValue >= k{
			preValue /= nums[begin]
			begin++
		}
		result += end - begin + 1
	}

	return result
}

func main(){
	data := []int{1,1,1}
	fmt.Println(numSubarrayProductLessThanK(data, 1))
}
