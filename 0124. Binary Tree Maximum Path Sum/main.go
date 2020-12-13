package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxValue = math.MinInt32

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	tmpMax1 := helper(root.Left)
	tmpMax2 := helper(root.Right)
	max2 := root.Val + tmpMax1 + tmpMax2
	maxValue = int(math.Max(float64(maxValue), float64(max2)))
	var res int
	if tmpMax1 > tmpMax2 {
		res = root.Val + tmpMax1
	} else {
		res = root.Val + tmpMax2
	}
	if res > 0 {
		return res
	}
	return 0
}

func maxPathSum(root *TreeNode) int {
	maxValue = math.MinInt32
	helper(root)
	return maxValue
}

func main() {
	var root TreeNode
	root.Val = -3
	fmt.Println(maxPathSum(&root))
}
