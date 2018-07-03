/*简单的动态规划问题，求最后到达楼梯的最少体力
相当于我们有len(cost)个台阶，我们用dp数组来记录
到达每一层的cost。到达第m层台阶的方式只有两种从m-2
跳跃两层和从m-1跨越一层。
所以递推公式 dp[i] = cost[i] + min(dp[i-2],dp[i-1])
*/
package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	arry := make([]int,len(cost),len(cost))
	arry[0] = cost[0]
	arry[1] = cost[1]
	for i := 2; i < len(cost);i++{
		if arry[i - 2] < arry[i - 1]{
			arry[i] = arry[i - 2] + cost[i]
		}else{
			arry[i] = arry[i - 1] + cost[i]
		}
	}
	if arry[len(cost) - 1] < arry[len(cost) - 2]{
		return arry[len(cost) - 1]
	}else{
		return arry[len(cost)- 2]
	}
}

func main(){
	cost := []int{10,15,20}
	fmt.Print(minCostClimbingStairs(cost))
}