/**
贪心算法，只看局部最优解，能挣钱就卖。不能挣钱就不卖，买卖次数无限制
*/
package main

func maxProfit(prices []int) int {
	maxprofit := 0
	value := 0
	for i := 1; i < len(prices);i++{
		value = prices[i] - prices[i - 1]
		if value > 0{
			maxprofit = maxprofit + value
		}
	}
	return maxprofit
}

func main(){

}

