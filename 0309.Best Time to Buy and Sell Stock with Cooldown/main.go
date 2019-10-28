package main

import "fmt"

/*
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])

动态方程入上

如果在i天手头没股票，那么要么是前天已卖，或者今天卖了
如果在i天手头有股票。那么要么是上上天买了，要么是今天买了。
*/
func maxProfit(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	var dp [][]int
	for i := 0;i < len(prices);i++{
		for j := 0;j < 2;j++{
			dp = append(dp, make([]int,2))
		}
	}
	for i := 0;i < len(prices);i++{
		if i == 0{
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		if i == 1{
			if dp[0][0] > dp[0][1] + prices[1]{
				dp[i][0] = dp[0][0]
			}else{
				dp[i][0] = dp[0][1] + prices[1]
			}
			if -prices[1] > dp[0][1]{
				dp[i][1] = -prices[1]
			}else{
				dp[i][1] = dp[0][1]
			}
			continue
		}
		if dp[i - 1][0] > dp[i-1][1] + prices[i]{
			dp[i][0] = dp[i - 1][0]
		}else{
			dp[i][0] = dp[i-1][1] + prices[i]
		}
		if dp[i - 1][1] > dp[i-2][0] - prices[i]{
			dp[i][1] = dp[i - 1][1]
		}else{
			dp[i][1] = dp[i-2][0] - prices[i]
		}
	}
	//fmt.Println(dp[0],dp[1],dp[2])
	return dp[len(prices) - 1][0]
}

func main(){
	data := []int{1}
	fmt.Println(maxProfit(data))
}
