/*
题目就是不断的反转左右子树，所以不管用什么递归遍历都可以实现。
*/
package main

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right,root.Left
	return root
}

func main(){

}
