/*
对于所有的卖股票的情况，对于某一天来说都有两种状态：
用个三维的数组来表示，i代表的是某天，0 - len(prices) - 1
k代表的是可以进行多少次交易，r代表的是是否持有股票今天。
比如dp[1][1][0]代表的就是在第二天的时候我们还可以进行一次交易，现在我们
手头没有股票。
	对于i我们就是遍历，对于r来说分为以下几种情况：
	r = 0代表手头没有股票，这时候有两种原因，一种是昨天就没有，还有就是今天卖了。
	dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
                   max(   选择 rest  ,           选择 sell      )

	r = 1代表手头有股票，这时候也有两种，一种是昨天就有，还有一种是今天刚买的。
    dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
              max(   选择 rest  ,           选择 buy         )

	具体可以参照下面这个思路
    链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-lab/

*/

package main

import "fmt"

func maxProfit(prices []int) int {
	buyTimes := 2
	var dp [100000][3][2]int
	if len(prices) == 0{
		return 0
	}
	for i := 0;i < len(prices);i++{
		for j :=  buyTimes; j > 0;j--{
			if i  == 0{
				dp[i][j][0] = 0;
				dp[i][j][1] = -prices[i];
				continue;
			}
			if dp[i-1][j][0] >  dp[i-1][j][1] + prices[i]{
				dp[i][j][0] = dp[i-1][j][0]
			}else{
				dp[i][j][0] = dp[i-1][j][1] + prices[i]
			}
			if dp[i-1][j][1] >  dp[i-1][j - 1][0] - prices[i]{
				dp[i][j][1] = dp[i-1][j][1]
			}else{
				dp[i][j][1] = dp[i-1][j - 1][0] - prices[i]
			}
		}
	}
	return dp[len(prices) - 1][buyTimes][0]
}

func main(){
	testData := []int{1,11,4,7,2}
	fmt.Println(maxProfit(testData))
}
