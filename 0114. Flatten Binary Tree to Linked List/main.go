package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	helper(root.Right)
	tmpRight := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmpRight
}

func flatten(root *TreeNode) {
	helper(root)
}

func main() {

}
