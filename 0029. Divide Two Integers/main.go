package main

import "math"

func divide(dividend int, divisor int) int {
	sameSign := true
	var result int
	if dividend^divisor < 0 {
		sameSign = false
	}
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	for i := 31; i >= 0; i-- {
		if dividend>>i >= divisor {
			result = result + 1<<i
			dividend = dividend - divisor*1<<i
		}
	}
	if !sameSign {
		return -result
	}
	if result > 2147483647 {
		return 2147483647
	}
	return result
}

func main() {

}
