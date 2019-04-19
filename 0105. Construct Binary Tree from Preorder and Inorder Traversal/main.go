/*
树的问题基本离不开递归，前中遍历
{3,9,20,15,7}
{9,3,15,20,7}

前序的第一个变量肯定是父节点，然后根据中序判断哪些是该节点
的左树和右树。在3之前的肯定是左孩子，在3之后的是右孩子。所以
对于3这个节点来说
9 是左孩子
20,15,7 前序
15,20,7 中序

然后继续对20 递归即可


*/
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
 }

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0{
		return nil
	}
	if len(preorder) == 1 && len(inorder) == 1{
		root := new(TreeNode)
		root.Val = preorder[0]
		return  root
	}
	fmt.Println(preorder,inorder)
	root := new(TreeNode)
	root.Val = preorder[0]
	var nextLeftInorder []int
	var nextLeftPreorder []int
	var nextRightInorder []int
	var nextRightPreorder []int
	i := 0
	for i = 0;i < len(inorder);i++{
		if preorder[0] == inorder[i]{
			nextLeftInorder = inorder[0:i]
			nextRightInorder = inorder[i+1:len(inorder)]
			break
		}
	}
	nextLeftPreorder = preorder[1:len(nextLeftInorder) + 1]
	nextRightPreorder = preorder[len(nextLeftInorder) + 1:len(preorder)]
	root.Left = buildTree(nextLeftPreorder,nextLeftInorder)
	root.Right = buildTree(nextRightPreorder,nextRightInorder)
	return root
}

func main(){
	pre := []int{3,9,20,15,7}
	mid := []int{9,3,15,20,7}
	buildTree(pre,mid)
}