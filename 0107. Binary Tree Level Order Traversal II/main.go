package main

import (
	"fmt"
	"sort"
)

func fun1(nums1[]int, nums2 []int)float32{
	nums1 = append(nums1,nums2...)
	sort.Ints(nums1)
	var result float32
	if len(nums1) % 2 == 0{
		result = float32((nums1[len(nums1) / 2] + nums1[len(nums1) / 2 -1])) / 2
	}else{
		result = float32(nums1[len(nums1) / 2])
	}

	return result
}

func main(){
	num1 := []int{1, 2}
	num2 := []int{3,4}
	fmt.Println(fun1(num1,num2))
}
