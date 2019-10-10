/*
该算法是贪心算法。

从头遍历，对于下表为i的变量来说，如果前面的sum + nums[i]大于0证明
可以加该值，但需要和前面的sum比
result = MAX(sum, sum + nums[i])
如果小于0，则sum直接归零。
*/

package main

import "fmt"

func maxSubArray(nums []int) int {
	max_sum := nums[0]
	current_sum := 0
	for index := 0;index < len(nums);index++{
		current_sum = current_sum + nums[index]
		if (current_sum) >= 0{
			if current_sum > max_sum {
				max_sum = current_sum
			}
		}else{
			if current_sum > max_sum {
				max_sum = current_sum
			}
			current_sum = 0
		}
	}
	return max_sum
}

func main(){
	array := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Print(maxSubArray(array))
}
