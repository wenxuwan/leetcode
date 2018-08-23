package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	if root == nil{
		return  0
	}
	if maxDepth(root.Left) > maxDepth(root.Right){
		return 1 + maxDepth(root.Left)
	}else {
		return  1 + maxDepth(root.Right)
	}
}
