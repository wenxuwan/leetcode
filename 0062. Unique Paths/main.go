package main

import "fmt"

func uniquePaths(m int, n int) int {
	var path [][]int
	for i := 0; i < m;i++{
		path = append(path,make([]int, n))
	}

	for i := 0; i < m; i++{
		for j :=0;j < n;j++{
			if i != 0 && j != 0{
				path[i][j] = path[i - 1][j] + path[i][j - 1]
			}else{
				path[i][j] = 1
			}
		}
	}
	return path[m - 1][n -1]
}

func main(){
	m := 7
	n := 3
	fmt.Println(uniquePaths(m,n))
}
