package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	var resultarr [2001]int
	var resultcurr [2001]int
	resultarr[1000-nums[0]] += 1
	resultarr[1000+nums[0]] += 1
	for index, vars := range nums {
		if index == 0 {
			continue
		}
		for i := 0; i < 2000; i++ {
			if resultarr[i] != 0 {
				resultcurr[i+vars] += resultarr[i]
				resultcurr[i-vars] += resultarr[i]
			}

		}
		for i := 0; i < 2000; i++ {
			resultarr[i] = resultcurr[i]
			resultcurr[i] = 0
		}

	}
	if S > 1000 {
		return 0
	}
	return resultarr[1000+S]
}
func main() {
	var testvalue []int = []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays(testvalue, 3))

}
