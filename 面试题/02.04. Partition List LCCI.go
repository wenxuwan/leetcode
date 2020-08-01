package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	slow := head
	fast := head
	for slow != nil {
		if slow.Val >= x {
			fast = slow.Next
			break
		}
		slow = slow.Next
	}
	if slow == nil {
		return head
	}
	for fast != nil && slow != nil {
		if fast.Val < slow.Val {
			fast.Val, slow.Val = slow.Val, fast.Val
			for slow.Val < x {
				slow = slow.Next
			}
		}
		fast = fast.Next
	}
	return head
}

func main() {

}
