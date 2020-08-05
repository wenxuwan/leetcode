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

func helper(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		return helper(head.Next, root.Left) || helper(head.Next, root.Right)
	}
	return false
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	if helper(head, root) {
		return true
	}
	if isSubPath(head, root.Left) {
		return true
	}
	if isSubPath(head, root.Right) {
		return true
	}
	return false
}

func main() {

}
