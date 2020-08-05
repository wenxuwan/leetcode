package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	middle := len(nums) / 2
	root.Val = nums[middle]
	leftArry := nums[0:middle]
	rightArry := nums[middle+1 : len(nums)]
	root.Left = helper(leftArry)
	root.Right = helper(rightArry)
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	array := make([]int, 0, 0)
	node := head
	for node != nil {
		array = append(array, node.Val)
		node = node.Next
	}
	return helper(array)
}

func main() {

}
