package main

func find(parent []int, x int) int {
	if parent[x] == -1 {
		return x
	}
	return find(parent, parent[x])
}
func union(parent []int, x, y int) {
	p1 := find(parent, x)
	p2 := find(parent, y)
	if p1 != p2 {
		parent[p1] = p2
	}
	if parent[y] != -1 && parent[y] != p2 {
		parent[y] = p2
	}
}

func findCircleNum(M [][]int) int {
	result := make([]int, len(M))
	for index, _ := range result {
		result[index] = -1
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 && i != j {
				union(result, i, j)
			}
		}
	}
	count := 0
	for i := 0; i < len(result); i++ {
		if result[i] == -1 {
			count++
		}
	}
	return count
}

func main() {

}
