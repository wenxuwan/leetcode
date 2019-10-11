/*
和62题思路基本一样，唯一的不同就是需要判断自己的值是0还是
1。
*/
package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var path [][]int
	for i := 0; i < len(obstacleGrid);i++{
		path = append(path,make([]int, len(obstacleGrid[0])))
	}
	for i := 0; i < len(obstacleGrid); i++{
		for j := 0;j < len(obstacleGrid[0]);j++{
			if i != 0 && j != 0 && obstacleGrid[i][j] == 0{
				path[i][j] = path[i - 1][j] + path[i][j - 1]
			}else{
				if i == 0 && j == 0{
					if obstacleGrid[i][j] == 0{
						path[i][j] = 1
					}
				}else{
					if j == 0{
						if obstacleGrid[i][j] == 0{
							path[i][j] = path[i - 1][j]
						}
					}else{
						if obstacleGrid[i][j] == 0{
							path[i][j] = path[i][j - 1]
						}
					}
				}
			}
		}
	}
	return path[len(obstacleGrid) - 1][len(obstacleGrid[0]) -1]
}

func main(){

}
