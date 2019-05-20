package main

import "fmt"

/*
从第一个节点开始，挨着记录每过一个加油站，剩余的汽油的数量，当汽油的数量小于
0的时候代表前面的汽油不够，从该节点重新开始。遍历完毕后，比较剩下的汽油的数量
和前面需要的汽油的数量。如果剩余的大于需要的汽油的数量，那么就代表可以走完。

完全可以理解为一个循环列表，只要从某个开始到自己结束所有的gas的总和大于cost的
就可以了。
*/
func canCompleteCircuit(gas []int, cost []int) int {
	preRouteGasNedd := 0
	startIndex := 0
	gasLeft := 0
	for i := 0; i < len(gas);i++{
		if gasLeft >= 0{
			gasLeft = gasLeft - cost[i] + gas[i]
		}else{
			preRouteGasNedd = preRouteGasNedd + gasLeft
			gasLeft = gas[i] - cost[i]
			startIndex = i
		}
	}
	//fmt.Println(gasLeft,preRouteGasNedd)
	if gasLeft < -preRouteGasNedd{
		return -1
	}
	return startIndex
}
func main(){
	gas  := []int{4,5,2,6,5,3}
	cost := []int{3,2,7,3,2,9}
	fmt.Println(canCompleteCircuit(gas,cost))
}
