/*
第一反应是用一个哈希表记录一下，但需要On大小的空间复杂度
算法要求的是时间复杂度是n,空间是常数，所以只能在原地变化。

1.遍历数组，把小于0的和大于len(nums)的数都标为1，防止后面数组下标越界
2.如果遍历过程中每出现1，直接返回1
3.把现有的数据以值作为下标，把里面的值置为负数（注意如果已经是负数，那么代表重复出现，不需要再重置）
4.然后从头遍历，找到第一个值不为负数的，然后输出下标值。
5.如果没找到，打印出len(nums) + 1
*/
package main

import "fmt"

func firstMissingPositive(nums []int) int {
	have_number_one := false
	for i, value := range nums{
		if value == 1{
			have_number_one = true
		}
		if value <= 0 || value > len(nums){
			nums[i] = 1
		}
	}
	if have_number_one == false{
		return 1
	}
	for _, value := range nums{
		if value < 0{
			value = -value
		}
		if nums[value - 1] > 0{
			nums[value - 1] = -nums[value - 1]
		}
	}
	fmt.Println(nums)
	for i, value := range nums{
		if value > 0{
			return i + 1
		}
	}

	return len(nums) + 1
}
func main(){

}
