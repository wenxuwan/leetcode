/*
这个题目的问题在于时间复杂度要求lgn

第一时间想到的是二分法，但传统二分法有个问题就是只能查到value的位置。
题目要求的是求出起始和结束的位置。所以做法就是在二分法找到的时候，一个往
左边查找，一个往右边查找。

所以这里就调用了两次的二分法。

*/
package main

import "fmt"

var left_flag int = -1
var right_flag int = -1

func f_left(nums []int, target int, start int, end int){
	for start <= end{
		midle := start + (end - start) / 2
		if nums[midle] == target{
			if right_flag == -1{
				right_flag = midle
			}
			left_flag = midle
			f_left(nums, target, start, midle - 1)
			break
		}else{
			if nums[midle] < target{
				start = midle + 1
			}else{
				end = midle - 1
			}
		}
	}
}

func f_right(nums []int, target int, start int, end int){
	for start <= end{
		midle := start + (end - start) / 2
		if nums[midle] == target{
			if left_flag == -1{
				left_flag = midle
			}
			right_flag = midle
			f_right(nums, target, midle + 1, end)
			break
		}else{
			if nums[midle] < target{
				start = midle + 1
			}else{
				end = midle - 1
			}
		}
	}
}

func searchRange(nums []int, target int) []int {
	left_flag = -1
	right_flag = -1
	result := make([]int, 2, 2)
	f_left(nums,target,0,len(nums) - 1)
	f_right(nums,target,0,len(nums) - 1)
	fmt.Println(left_flag,right_flag)
	result[0] = left_flag
	result[1] = right_flag
	return result
}

func main(){
	nums := []int{5,7,7,8,8,10}
	searchRange(nums,6)
}
