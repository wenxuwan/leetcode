package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	result  := [3]int{-1000000000000000,-1000000000000000,-1000000000000000}
	for _ , value := range nums{
		if value == result[0] || value == result[1] || value == result[2]{
			continue
		}
		if value > result[0]{
			result[2] = result[1]
			result[1] = result[0]
			result[0] = value
			continue
		}
		if value > result[1]{
			result[2] = result[1]
			result[1] = value
			continue
		}
		if value > result[2]{
			result[2] = value
			continue
		}
	}
	if result[2] == -1000000000000000{
		return result[0]
	}
	return result[2]
}

func main(){
	test := []int{1,2}
	fmt.Print(thirdMax(test))
}
