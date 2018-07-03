package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for index, x := range nums {
		for j := index + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				result = append(result, index)
				result = append(result, j)
				break
			}
		}
		if len(result) != 0 {
			break
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	value1 := 6
	slice := twoSum(arr, value1)
	fmt.Println(slice)
}
