package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var tmp *ListNode
	if l1.Val < l2.Val {
		tmp = l1
		tmp.Next = mergeTwoLists(l1.Next, l2)
		return tmp
	} else {
		tmp = l2
		tmp.Next = mergeTwoLists(l1, l2.Next)
		return tmp
	}

}
func main() {

}
