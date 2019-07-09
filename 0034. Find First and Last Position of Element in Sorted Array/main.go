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
