package main

import "fmt"

//题目真的挺牛逼的这个解法，可以看一下官方的第三种解法
func checkSubarraySum(nums []int, k int) bool {
	re := make(map[int]int, len(nums))
	rem := 0
	re[0] = -1 //这个是为了0,0这种情况的出现
	for index, value := range nums {
		rem = rem + value
		if k != 0 {
			rem = rem % k
		}
		if value, ok := re[rem]; ok {
			if index-value > 1 {
				return true
			}
		} else {
			re[rem] = index
		}
	}
	return false
}

func main() {
	data := []int{0, 1, 0}
	fmt.Println(checkSubarraySum(data, 0))
}
