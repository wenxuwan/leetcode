package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0{
		return  nil
	}
	root := new(TreeNode)
	middle := len(nums) / 2
	root.Val = nums[middle]
	leftArry := nums[0:middle]
	rightArry := nums[middle+1:len(nums)]
	root.Left = sortedArrayToBST(leftArry)
	root.Right = sortedArrayToBST(rightArry)
	return root
}