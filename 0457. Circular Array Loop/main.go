/*
求环就想到用快慢指针就可以了，如果从某个节点开始，后面构不成环，那么遍历过的
节点就都不可能形成环了，可以跳过。
*/

package main

import "fmt"

func helper(nums []int, begin int)int{
	next := ((nums[begin] + begin) % len(nums) + len(nums)) % len(nums)//让负数不要溢出
	return next
}

func circularArrayLoop(nums []int) bool {
	lenth := len(nums)
	slow := 0
	fast := 0
	for i := 0; i < lenth;i++{
		if nums[i] == 0{
			continue
		}
		slow = i
		fast = helper(nums, slow)
		val := nums[slow]
		for val * nums[fast] > 0 && val * nums[helper(nums,fast)] > 0{
			if fast == slow{
				if slow == helper(nums, slow){
					break
				}
				return true
			}
			slow = helper(nums, slow) //慢指针走一次
			fast = helper(nums,helper(nums, fast)) //快指针走两步
		}
		slow = i;
		for val * nums[slow] > 0 {
			nums[slow] = 0;
			slow =  helper(nums,slow);
		}

	}
	return false
}

func main(){
	testData := []int{2,-1,1,2,2}
	fmt.Println(circularArrayLoop(testData))
}
