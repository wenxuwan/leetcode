/*
递归遍历二叉树的题目，求出每个节点的最长路径，然后不断
的递归比较，保留最长路径。
*/
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func f(root *TreeNode, max *int)int{
	if root == nil{
		return 0
	}
	l := f(root.Left, max)
	r := f(root.Right, max)
	if l + r > *max{
		*max = l + r
	}
	if l > r{
		return l + 1
	}else{
		return  r + 1
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	f(root, &max)
	return max
}

func main(){

}
