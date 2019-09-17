/*
这个题目看到就是bfs的问题，从‘1’开始遍历所有可达的节点
然后将‘1’改成‘0’，继续
 */
package main

import "fmt"

func bfs(x int, y int, grid [][]byte){
	if x < 0 || x >= len(grid){
		return
	}
	if y < 0 || y >= len(grid[0]){
		return
	}
	if grid[x][y] == '1'{
		grid[x][y] = '0'
		bfs(x - 1, y, grid)
		bfs(x + 1, y, grid)
		bfs(x, y - 1, grid)
		bfs(x, y + 1,grid)
	}
}
func numIslands(grid [][]byte) int {
	landNumber := 0
	for i := 0;i < len(grid);i++{
		for j := 0; j < len(grid[0]);j++{
			if grid[i][j] == '1'{
				landNumber++
				bfs(i,j,grid)
			}
		}
	}
	return landNumber
}

func main(){
	value := [][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	fmt.Println(numIslands(value))
}
