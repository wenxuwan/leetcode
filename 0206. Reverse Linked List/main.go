package main
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return  head
	}
	p := head
	newList := new(ListNode)
	newList = nil
	for p != nil{
		q := p.Next
		p.Next = newList
		newList = p
		p = q
	}
	return newList
}
