package main

func countBits(num int) []int {
	result := make([]int, num+1)
	result[0] = 0
	for value := 1; value <= num; value++ {
		tmp := 0
		if value%2 == 0 {
			tmp = result[value/2]
		} else {
			tmp = result[value-1] + 1
		}
		result[value] = tmp
	}
	return result
}

func main() {

}
