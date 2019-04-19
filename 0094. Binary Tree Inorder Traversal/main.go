package main

type TreeNode struct {
 	Val int
 	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var Airy []int
	if root == nil{
		return nil
	}
	Airy = append(Airy,inorderTraversal(root.Left)...)
	Airy = append(Airy,root.Val)
	Airy = append(Airy,inorderTraversal(root.Right)...)
	return  Airy
}


