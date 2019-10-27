/*
该题目和题目279是类似的，解题思路可以看279
*/
package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {
	result := make([]int, amount + 1,  amount + 1)
	sort.Ints(coins)
	for i := 1;i <= amount;i++{
		result[i] = -1
		for j := 0; j < len(coins) && coins[j] <= i;j++{
			if result[i] > result[i - coins[j]] || result[i] == -1{
				if result[i - coins[j]] != -1{
					result[i] = result[i - coins[j]] + 1
				}

			}
		}
	}
	fmt.Println(result)
	return result[amount]
}

func main(){
	testData := []int{474,83,404,3}
	fmt.Println(coinChange(testData,264))
}
