package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func isValidBSTHelper(root *TreeNode,min int, max int) bool{
	if root == nil{
		return true
	}
	if(root.Val >= max || root.Val <= min){
		return false
	}
	return isValidBSTHelper(root.Left,min,root.Val) && isValidBSTHelper(root.Right,root.Val,max)
}
func isValidBST(root *TreeNode) bool {
	return (isValidBSTHelper(root,INT_MIN,INT_MAX))
}