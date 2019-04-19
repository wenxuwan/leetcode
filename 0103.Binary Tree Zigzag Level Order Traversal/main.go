package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int,0)
	if root == nil{
		return result
	}
	queue := make([]*TreeNode,0)
	l := make([]int,0,1)
	l =append(l,root.Val)
	result = append(result,l)
	head := 0
	i := 0
	end  := 0
	queue = append(queue,root)
	end++
	rowBegin := head
	rowEnd := end
	for head < len(queue) {
		rowList := make([]int,0)
		for rowBegin < rowEnd{
			if queue[rowBegin] == nil{
				fmt.Println(rowEnd,rowBegin,len(queue))
			}
			if queue[rowBegin].Left != nil{
				queue = append(queue,queue[rowBegin].Left)
				end++
			}
			if queue[rowBegin].Right != nil{
				queue = append(queue,queue[rowBegin].Right)
				end++
			}
			rowList = append(rowList,queue[rowBegin].Val)
			rowBegin++
		}
		i++
		if i % 2 == 0{
			rowList = reverse(rowList)
		}
		if head != 0{
			result = append(result,rowList)
		}
		rowEnd = end
		head = rowBegin
		fmt.Println(head,end,rowBegin,rowEnd,result)
	}
	return  result
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
	//fmt.Println(preorder,inorder)
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
	fmt.Println(zigzagLevelOrder(buildTree(pre,mid)))
}