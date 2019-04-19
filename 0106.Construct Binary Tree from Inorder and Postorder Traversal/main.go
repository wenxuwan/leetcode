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

func buildTree(inorder []int, postorder []int ) *TreeNode {
	fmt.Println(inorder,postorder)
	if len(postorder) == 0 || len(inorder) == 0{
		return nil
	}
	if len(postorder) == 1 && len(inorder) == 1{
		root := new(TreeNode)
		root.Val = postorder[0]
		return  root
	}
	root := new(TreeNode)
	root.Val = postorder[len(postorder) - 1]
	var nextLeftInorder []int
	var nextLeftPostorder []int
	var nextRightInorder []int
	var nextRightPostorder []int
	i := 0
	for i = 0;i < len(inorder);i++{
		if root.Val == inorder[i]{
			nextLeftInorder = inorder[0:i]
			nextRightInorder = inorder[i+1:len(inorder)]
			break
		}
	}
	nextLeftPostorder = postorder[0:len(nextLeftInorder)]
	nextRightPostorder = postorder[len(nextLeftInorder):len(postorder) - 1]
	root.Left = buildTree(nextLeftInorder,nextLeftPostorder)
	root.Right = buildTree(nextRightInorder,nextRightPostorder)
	return root
}

func buildTreeWithOptimize(inorder []int, postorder []int ) *TreeNode{
	if len(inorder) == 0{
		return nil
	}
	val,index := postorder[len(postorder) - 1],0
	for inorder[index] != val{
		index++
	}
	node := &TreeNode{val,nil,nil}
	node.Left = buildTreeWithOptimize(inorder[0:index],postorder[0:index])
	node.Right = buildTreeWithOptimize(inorder[index+1:],postorder[index:len(postorder) - 1])
	return node
}

func main(){
	post := []int{9,15,7,20,3}
	mid := []int{9,3,15,20,7}
	buildTree(mid,post)
}