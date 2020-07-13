package main

import "fmt"

//看了一下题解。。。。
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums))
	currentQueue := make([]int, 0, len(nums))
	for i, value := range nums {
		l := len(currentQueue)
		for l != 0 && value >= nums[currentQueue[l-1]] {
			currentQueue = currentQueue[:l-1]
			l--
		}
		currentQueue = append(currentQueue, i)
		if currentQueue[0] == i-k {
			currentQueue = currentQueue[1:]
		}
		if i+1-k >= 0 {
			result = append(result, nums[currentQueue[0]])
		}
	}
	return result
}

func main() {
	data := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(data, k))
}
