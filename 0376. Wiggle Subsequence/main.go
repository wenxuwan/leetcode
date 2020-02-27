/*
想到的就是动态规划的路子，动态规划最关键的还是要站在一个节点上看符不符合条件。

*/

package main

func wiggleMaxLength(nums []int) int {
	l := len(nums)
	i := 0
	dp := make([]int,l,l)
	if l < 2{
		return l
	}
	dp[0] = 1
	ppValue := nums[0]
	pValue := 0
	for i = 1;i < l;i++{
		if nums[i] != nums[0]{
			pValue = nums[i]
			dp[i] = 2
			break
		}else{
			dp[i] = 1
		}
	}
	if i == l{
		return 1
	}
	i++
	for ; i < l;i++{
		r := (nums[i] - pValue) * (pValue - ppValue)
		if r != 0{
			ppValue = pValue
			pValue = nums[i]
		}
		if  r < 0{
			dp[i] = dp[i - 1] + 1
		}else{
			dp[i] = dp[i - 1]
		}
	}
	return dp[l - 1]
}


func main(){
	nums := []int{33,53,12,64,50,41,45,21,97,35,47,92,39,0,93,55,40,46,69,42,6,95,51,68,72,9,32,84,34,64,6,2,26,98,3,43,30,60,3,68,82,9,97,19,27,98,99,4,30,96,37,9,78,43,64,4,65,30,84,90,87,64,18,50,60,1,40,32,48,50,76,100,57,29,63,53,46,57,93,98,42,80,82,9,41,55,69,84,82,79,30,79,18,97,67,23,52,38,74,15}
	wiggleMaxLength(nums)
}
