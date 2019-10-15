package main

import "fmt"

func findDuplicate(nums []int) int {
	result := make(map[int]int)
	for _, value := range nums{
		result[value] = result[value] + 1
	}
	for key,value := range result{
		if value > 1{
			return key
		}
	}
	return -1
}
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast{
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	begin := 0
	for begin != slow{
		begin = nums[begin]
		slow = nums[slow]
	}
	return begin
}

func main(){

}
