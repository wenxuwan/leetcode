package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totleLen := len(nums1) + len(nums2)
	result := make([]int,0,totleLen/2)
	step := 0
	step1 := 0
	step2 := 0
	var re float64
	for step <= totleLen /2{
		if step1 == len(nums1){
			result = append(result,nums2[step2])
			step2++
			step++
			continue
		}
		if step2 == len(nums2){
			result = append(result,nums1[step1])
			step1++
			step++
			continue
		}
		if nums1[step1] < nums2[step2]{
			result = append(result,nums1[step1])
			step1++
		}else {
			result = append(result,nums2[step2])
			step2++
		}
		step++
	}
	fmt.Println(result)
	if totleLen % 2 == 0{
		re = float64(result[len(result) - 1] + result[len(result) - 2]) / 2
	}else{
		re = float64(result[len(result) - 1])
	}
	return re
}

func main(){
	nums1 := []int{1, 2}
	nums2 := []int{3}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
}
