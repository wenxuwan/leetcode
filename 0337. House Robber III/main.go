/*
动态规划问题，需要找出动态方程。
和前面的打家劫舍1基本类型差不多，唯一的改变是数据结构从
数组变为了树。

对于动态规划的题目，我们要站在"任何一个节点"上来求出动态方程。
对于任一个节点来说他只有两种选择，偷或者不偷。所以我们可以用
一个长度为2的数组表示该节点偷或者不偷，假设0代表不偷，1代表
的就是偷。

那么该节点不偷的时候，孩子节点偷不偷都行，该节点最大值为左右孩子
的最大值：

p[0] = MAX(left[0],left[1]) + MAX(right[0],right[1])

如果该节点偷，那么所得就是两个孩子都不偷：
p[1] = left[0] + right[0] + p.val


*/
package main

import (
	"math"
)

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func helper(root *TreeNode)[2]int{
	var result [2]int
	if root == nil{
		return result
	}
	left := helper(root.Left)
	right := helper(root.Right)
	result[0] = int(math.Max(float64(right[0]),float64(right[1])) + math.Max(float64(left[0]),float64(left[1])))
	result[1] = right[0] + left[0] + root.Val
	return result
}
func rob(root *TreeNode) int {
	result := helper(root)
	if result[0] > result[1]{
		return result[0]
	}else{
		return result[1]
	}
}

func main(){

}
