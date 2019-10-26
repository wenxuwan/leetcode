/*
该题目一开始没想到用动态规划，想的是dfs，但会超时，
贪心法也不适合本题目，因为并不是每次都取最大的就可以
保证后面肯定是最大的。
所以就看了一下解答，用动态规划，动态规划的本质还是找出
转换方程。

首先定义一个数组result[n + 1]用来存储0-n的最优解，初始化
对应的i，最差的情况下，result只能由i个1组成。

然后就是转换公式。对于result[n]来说，其最小值肯定对应的
是result[n - j * j] + 1或者是本身的result[n]，所以转换
方程就变成了result[n] = min(result[n], result(n - j * j) + 1)

*/
package main

import "fmt"

func numSquares(n int) int {
	result := make([]int, n + 1,  n + 1)
	for i := 1;i <= n;i++{
		result[i] = i
		for j := 1; i - j * j >= 0;j++{
			if result[i] > result[i - j * j] + 1{
				result[i] = result[i - j * j] + 1
			}else{
				result[i] = result[i]
			}
		}
	}
	return result[n]
}

func main(){
	fmt.Println(numSquares(43))
}
