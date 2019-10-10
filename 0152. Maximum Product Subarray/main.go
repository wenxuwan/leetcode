/*
这个问题是个动态规划的问题。
用两个变量max和min分别来表示每一步的时候的最大最小值。
可以得到以下方程：
max = MAX(nums[i], nums[i] * max)
min = MIN(nums[i], nums[i] * max)
这里比较关键的点在于当出现负数的时候怎么处理。我们都知道
出现负数的时候最大值会变为最小值，最小值会变为最大值。所以
在nums[i] < 0 的时候我们要做的就是对调两个变量的值。
max, min := min, max
*/
package main

func maxProduct(nums []int) int {
	maxResult := nums[0]
	max := 1
	min := 1
	for i := 0;i < len(nums);i++{
		if nums[i] < 0{
			max,min = min, max
		}
		if max * nums[i] < nums[i]{
			max = nums[i]
		}else{
			max = max * nums[i]
		}
		if min * nums[i] > nums[i]{
			min = nums[i]
		}else{
			min = min * nums[i]
		}
		if max >  maxResult{
			maxResult = max
		}
	}
	return  maxResult
}

func main(){


}
