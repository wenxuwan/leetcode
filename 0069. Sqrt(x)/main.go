package main

func mySqrt(x int) int {
	left, right := 0, x
	result := -1
	for left <= right {
		mid := (right + left) / 2
		if mid*mid > x {
			right = mid - 1
		} else {
			result = mid
			left = mid + 1
		}
	}
	return result
}

func main() {

}
