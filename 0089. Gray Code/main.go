package main

import "fmt"

var visited map[int]bool
var result []int

func helper(current, n int) bool {
	result = append(result, current)
	if len(result) == (1 << n) {
		return true
	}
	visited[current] = true
	var i int
	for i = 0; i < n; i++ {
		next := current ^ (1 << i)
		if _, ok := visited[next]; !ok && helper(next, n) {
			return true

		}
	}
	result = result[:len(result)-1]
	delete(visited, current)
	return false
}

func grayCode(n int) []int {
	visited = make(map[int]bool)
	result = make([]int, 0)
	helper(0, n)
	return result
}

func main() {
	fmt.Println(grayCode(2))
}
