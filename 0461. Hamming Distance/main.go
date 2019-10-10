package main

func hammingDistance(x int, y int) int {
	result  := x ^ y
	count := 0
	for ; result > 0; result /= 2{
		if result % 2 == 1{
			count++
		}
	}
	return count
}

func main(){
	hammingDistance(1, 4)
}
