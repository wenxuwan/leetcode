package main

import "fmt"

func removeElement(nums []int, val int) int {
	new_len := len(nums)
	right_index := len(nums) - 1
	left_index := 0
	for {
		for left_index <= len(nums) - 1 && nums[left_index] != val{
			left_index++
		}
		for right_index >= 0 && nums[right_index] == val{
			right_index--
			new_len = new_len - 1
		}
		if(left_index < right_index){
			nums[left_index],nums[right_index]= nums[right_index],nums[left_index]
			left_index = left_index + 1
			right_index = right_index - 1
			new_len = new_len - 1
		}else{
			break
		}
	}
	return new_len
}

func main(){
	array := [ ]int{3,4,4,3}
	fmt.Print(removeElement(array,4))
}
