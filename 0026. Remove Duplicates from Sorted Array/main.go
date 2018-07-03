package main

import "fmt"

func removeDuplicates(nums []int) int {
	dumvalue := 0
	for index, num := range nums {
		if index != 0 && num == nums[index-1] {
			dumvalue++
		} else {
			if dumvalue != 0 {
				nums[index-dumvalue] = num
			}

		}

	}
	return len(nums) - dumvalue
}

func main() {
	var s1 []int = []int{2, 4}
	len := removeDuplicates(s1)
	for i := 0; i < len; i++ {
		fmt.Println(s1[i])
	}

}
