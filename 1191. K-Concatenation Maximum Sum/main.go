package main

import (
	"fmt"
	"math"
)

func helper(arr []int, len int) (int, int, int) {
	tmpMax := 0
	totleMax := math.MinInt64
	oneArryMax := math.MinInt64
	totleSum := 0
	begin, end := 0, 0
	tmpBegin := 0

	for index, value := range arr {
		if index == len {
			oneArryMax = totleSum
		}
		totleSum += value
		tmpMax += value
		if tmpMax < 0 {
			tmpBegin = index + 1
			tmpMax = 0
		} else {
			if tmpMax >= totleMax {
				begin = tmpBegin
				totleMax = tmpMax
				end = index
			}
		}
	}
	if totleMax < 0 {
		return 0, 0, totleSum
	}
	fmt.Printf("totleMax is %+v, start at begin %+v, end %+v \n", totleMax, begin, end)
	return totleMax, oneArryMax, totleSum
}
func kConcatenationMaxSum(arr []int, k int) int {
	lenth := len(arr)
	tmpArr := append(arr, arr...)
	mod := int(math.Pow10(9) + 7)
	totleMax, oneArryMax, totleSum := helper(tmpArr, lenth)
	if totleSum >= 0 {
		if totleMax > oneArryMax {
			return (oneArryMax*(k-2) + totleMax) % mod
		}
		return (oneArryMax*(k-2) + oneArryMax) % mod
	} else {
		if totleMax > oneArryMax {
			return totleMax % mod
		}
		return oneArryMax % mod
	}
}

func main() {
	arr := []int{-5, -7, -9993, -10000}
	helper(arr, 4)
}
