package main

import (
	"math"
	"fmt"
)

func maxArea(height []int) int {
	lenth := len(height)
	i := 0
	j := lenth - 1
	result := 0
	for i < j{
		tmp := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
		result = int(math.Max(float64(result),float64(tmp)))
		if height[i] < height[j]{
			i++
		}else{
			j--
		}
	}
	return result
}

func main(){
	tmp := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(tmp))
}