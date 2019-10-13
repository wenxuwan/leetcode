/*
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

因为行和列都是递增的，所以我们从第一行的最后面往前找小于target的列
找到后如果不想当则往下走，这样可以每次移动都抛弃某行或者某列，所以时间
复杂度是O(m + n)
*/

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0{
		return false
	}
	row := 0
	column := len(matrix[0]) - 1
	for row < len(matrix) && column >= 0{
		if matrix[row][column] == target{
			return true
		}
		if matrix[row][column] > target{
			column--
		}else{
			row++
		}
	}
	return false
}

func main(){
	test := [][]int{{-5}}
	fmt.Println(searchMatrix(test,0))
}
