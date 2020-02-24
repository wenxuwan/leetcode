package main

import (
	"fmt"
	"math"
)

/*
和打家劫舍题目是一个套路，把数组转换成每个值对应的value。这样就和打家劫舍一模一样了。
不能拿相邻的数字。
*/

func deleteAndEarn(nums []int) int {
	numsMap := make(map[int]int,len(nums))
	maxVlaue := 0
	for _, value := range nums{
		numsMap[value] = numsMap[value]+ value
		if value >= maxVlaue{
			maxVlaue = value
		}
	}
	if maxVlaue == 0{
		return 0
	}
	dp := make([]int,maxVlaue + 1,maxVlaue + 1)
	dp[0] = 0
	dp[1] = numsMap[1]
	for i := 2;i <= maxVlaue;i++{
		dp[i] = int(math.Max(float64(dp[i - 1]),float64(dp[i-2] + numsMap[i])))

	}
	return dp[maxVlaue]
}

func main(){
	nums := []int{}
	fmt.Println(deleteAndEarn(nums))
}