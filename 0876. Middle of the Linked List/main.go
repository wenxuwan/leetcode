package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}
func middleNode(head *ListNode) *ListNode {
	fmt.Println(head.Val)
	fast := head
	slow := head
	for fast != nil && fast.Next != nil{
		fast = fast.Next
		if fast != nil{
			fast = fast.Next
		}
		slow = slow.Next
	}
	return slow
}