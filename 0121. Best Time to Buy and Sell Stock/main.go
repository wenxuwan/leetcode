package main

//dp[i]代表[0,i]内交易股票的最大收益 minValue［0，i）之间价格的最小值
//状态转移方程 dp[i] = max(dp[i-1],prices[i]-minValue)

func maxProfit(prices []int) int {
	if len(prices) <=0{
		return 0
	}
	minValue := prices[0]
	dp := make([]int,len(prices),len(prices))
	for i := 1; i < len(prices);i++{
		if minValue > prices[i]{
			minValue = prices[i]
		}
		if dp[i-1] > (prices[i] - minValue){
			dp[i] = dp[i-1]
		}else{
			dp[i] = prices[i] - minValue
		}
	}
	return dp[len(prices) - 1]
}

func main(){

}