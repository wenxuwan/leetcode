package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetricHelper(leftChild *TreeNode,rightChild *TreeNode) bool {
	if leftChild == nil && rightChild != nil{
		return false
	}
	if leftChild != nil && rightChild == nil{
		return false
	}
	if leftChild == nil && rightChild == nil{
		return  true
	}
	if leftChild.Val != rightChild.Val{
		return false
	}
	return  isSymmetricHelper(leftChild.Right,rightChild.Left) && isSymmetricHelper(leftChild.Left,rightChild.Right)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return isSymmetricHelper(root.Left,root.Right)
}