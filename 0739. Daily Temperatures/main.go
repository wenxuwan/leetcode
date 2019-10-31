/*
从后面往前遍历，这样的话可以减少重复的遍历，
eg: {2,3,4,5,7,6,4,2,8}
从后面开始往前遍历
8的时候对应的是0
然后到2的时候是1
然后到4的时候是2
到6的时候发现4比它小，然后可以直接dp[6（数字4的下标）] + 5(数字6的下表)
这样就不需要再和2去比较了。可以减少一定的循环次数。


*/
package main

import "fmt"

func dailyTemperatures(T []int) []int {
	lenth := len(T)
	dp := make([]int, lenth, lenth)
	for i := lenth - 1; i >=0;i--{
		for j := i + 1; j < lenth;{
			if T[j] > T[i]{
				dp[i] = j - i
				break
			}else{
				if dp[j] == 0{
					dp[i] = 0
					break
				}else{
					j = j + dp[j]
				}
			}
		}
	}
	return dp
}

func main(){
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
