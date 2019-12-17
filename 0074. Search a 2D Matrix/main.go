/*
这个题目算法就是二分查找：

下面的时间复杂度应该是m + log n

如果直接看成m*n个长的有序数组，那么时间复杂度就是

logmn = logm + logn

所以下面的算法是可以再用二分法求到某行的

*/

package main

func helper(val []int, target int)bool{
	begin := 0
	end := len(val) - 1
	for begin <= end{
		middle := (begin + end) / 2
		if target > val[middle]{
			begin = middle + 1
		}
		if target < val[middle]{
			end = middle - 1
		}
		if target == val[middle]{
			return true
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	rowLen := len(matrix)
	if rowLen <= 0 || len(matrix[0]) == 0{
		return false
	}
	columnLen := len(matrix[0])
	for i := 0; i < rowLen;i++{//就是这里，其实不需要遍历
		if matrix[i][0] <= target &&  matrix[i][columnLen - 1] >= target{
			if helper(matrix[i], target){
				return true
			}else{
				return false
			}
		}
	}
	return false
}

func main(){

}
