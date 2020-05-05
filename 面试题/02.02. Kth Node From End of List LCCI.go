package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	slow := head
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil{
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}

func main() {

}
