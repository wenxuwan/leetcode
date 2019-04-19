package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	dp := make([]int,len(nums),len(nums))
	dp[0] = nums[0]
	if nums[0] >= nums[1]{
		dp[1] = nums[0]
	}else{
		dp[1] = nums[1]
	}
	for i := 2;i < len(nums);i++{
		if nums[i] + dp[i - 2] > dp[i - 1]{
			dp[i] = nums[i] + dp[i - 2]
		}else {
			dp[i] = dp[i - 1]
		}
	}
	return dp[len(nums) - 1]
}
func main(){
	arr := []int{1,2,3,1}
	fmt.Println(rob(arr))
}