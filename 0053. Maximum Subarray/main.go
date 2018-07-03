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
