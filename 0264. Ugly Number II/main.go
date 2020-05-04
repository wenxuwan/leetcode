package main

import (
	"fmt"
	"math"
)

var result []int
func init(){
	result = append(result, 1)
	i2, i3, i5 := 0,0,0
	for i := 1; i < 1690 ;i++{
		tmp := math.Min(float64(result[i2] * 2), float64(result[i3] * 3))
		res := math.Min(tmp, float64(result[i5] * 5))
		result = append(result, int(res))
		if res == float64(result[i2] * 2){
			i2++
		}
		if res == float64(result[i3] * 3){
			i3++
		}
		if res == float64(result[i5] * 5){
			i5++
		}
	}
}

func nthUglyNumber(n int) int {
	return result[n - 1]
}

func main(){
	fmt.Println(nthUglyNumber(378))
}