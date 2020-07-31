package main

import "fmt"

func waysToChange(n int) int {
	dp := make([]int, n+1, n+1)
	coins := []int{1, 5, 10, 25}
	for i := 0; i <= n; i++ {
		dp[i] = 1
	}
	for j := 1; j < 4; j++ {
		for p := 1; p <= n; p++ {
			if p >= coins[j] {
				dp[p] = (dp[p] + dp[p-coins[j]]) % 1000000007
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(waysToChange(929782))
}
