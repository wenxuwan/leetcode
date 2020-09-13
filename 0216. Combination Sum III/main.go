package main

import "fmt"

var tmpResult [][]int

func helper(k int, n int, value int, tmp []int, number int) {
	if k == 0 && n == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		tmpResult = append(tmpResult, t)
		return
	}
	if k <= 0 || n <= 0 {
		return
	}
	for i := value + 1; i <= 9; i++ {
		tmp[number] = i
		helper(k-1, n-i, i, tmp, number+1)
	}
}

func combinationSum3(k int, n int) [][]int {
	var result [][]int

	for i := 1; i <= 9; i++ {
		tmp := make([]int, k)
		tmp[0] = i
		helper(k-1, n-i, i, tmp, 1)
		result = tmpResult
	}
	defer func() {
		tmpResult = nil
	}()
	return result
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
