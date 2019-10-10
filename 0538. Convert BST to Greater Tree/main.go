/*
二叉树的东西基本都是递归解决，先把公式套出来：
对于任何一个节点来说：

root.val = root.val + 右节点的值

如果是左孩子需要把父节点传下去
root.val = root.val + 右节点的值 + 父节点的值(已改变过的)
如果是右孩子则需要把父节点之前的值加上
root.val = root.val + 右节点的值 + 父亲节点的父亲节点的值（已经改变过）


*/
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func f(root *TreeNode, value *int){
	if root == nil{
		return
	}
	*value += root.Val
	f(root.Left, value)
	f(root.Right, value)
}
func f2(root *TreeNode, rootValue int) {
	var value int
	if root == nil{
		return
	}
	f(root.Right, &value)
	root.Val = root.Val + value + rootValue
	f2(root.Right, rootValue)
	f2(root.Left, root.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	var value int
	if root == nil{
		return nil
	}
	f(root.Right, &value)
	fmt.Println(root.Val, value)
	root.Val = root.Val + value

	f2(root.Left, root.Val)
	f2(root.Right, 0)
	return root
}
func main(){

}
