/*
可以想象从第一个元素遍历，然后该点的次数等于左边孩子的排列个数 * 右边孩子的排列个数。
从而可得表达式：
f(m) = f（m - 1） * f(n - m)
*/
package main

import "fmt"

func numTrees(n int) int {
	result := make([]int,n + 1,n + 1)
	result[1] = 0
	result[0] = 1
	for i := 1;i <= n;i++{
		for j := 1; j <= i;j++{
			result[i] += result[j - 1] * result[i - j]
		}
	}
	fmt.Println(result)
	return result[n]
}

func main(){
	fmt.Println(numTrees(3))
}
