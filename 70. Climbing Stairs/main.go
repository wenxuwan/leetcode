package main

import "fmt"

/* 动态规划问题，假设我们要走到第一层台阶，那么我们只有一种走法，如果要到
第二个台阶那么我们有（1,1）（2）两种走法。如果要走到第三层那？因为步子的大
小为1和2，那么登上第三个台阶的走法就是从第一层直接走两步上来，或者从第二层
走一步上来。

这个地方就是问题的关键。f(n) = f(n-1) +ｆ（ｎ－２）　（ｎ　＞　２）这就是我们的
状态迁移公式，ｎ　＝０　和ｎ　＝　１的时候就不需要这样了。

*/
func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1{
		return 1
	}
	if n == 2{
		return 2
	}
	a := 1
	b := 2
	temp := 0
	for i:= 3 ; i <= n;i++{
		temp = a + b
		a,b = b,temp
	}
	return  temp
}

func main(){
	fmt.Print(climbStairs(10))
}
