/*
动态规划问题，因为不能一次性偷两个相邻的房子。用dp[]记录到
i房子的时候可以偷的最多的情况。那么方程就可以得到

dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])


*/
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