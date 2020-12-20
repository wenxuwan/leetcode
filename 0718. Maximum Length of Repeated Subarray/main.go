package main

func findLength(A []int, B []int) int {
	result := make([][]int, len(A))
	max := 0
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(B))
	}
	for i := 0; i < len(B); i++ {
		if A[0] == B[i] {
			result[0][i] = 1
		}
	}
	for i := 0; i < len(A); i++ {
		if B[0] == A[i] {
			result[i][0] = 1
		}
	}
	for i := 1; i < len(A); i++ {
		for j := 1; j < len(B); j++ {
			if A[i] == B[j] {
				result[i][j] = result[i-1][j-1] + 1
				if result[i][j] > max {
					max = result[i][j]
				}
			} else {
				result[i][j] = 0
			}
		}
	}
	return max
}

func main() {

}
