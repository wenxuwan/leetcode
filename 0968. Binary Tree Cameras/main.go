package main

var cnt int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode) int {
	if root == nil {
		return 2
	}
	left := helper(root.Left)
	right := helper(root.Right)
	if left+right >= 5 {
		return 2
	}
	if left == 2 && right == 2 {
		return 1
	} else {
		cnt++
		return 3
	}
}
func minCameraCover(root *TreeNode) int {
	cnt = 0
	if helper(root) == 1 {
		cnt++
	}
	return cnt
}

func main() {

}
