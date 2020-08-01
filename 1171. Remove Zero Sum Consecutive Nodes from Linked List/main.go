package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//这个题有个题解第一是真的够骚
func removeZeroSumSublists(head *ListNode) *ListNode {
	result := new(ListNode)
	result.Next = head
	tmpMap := make(map[int]*ListNode)
	sum := 0
	for node := result; node != nil; node = node.Next {
		sum += node.Val
		tmpMap[sum] = node
	}
	sum = 0
	for node := result; node != nil; node = node.Next {
		sum += node.Val
		node.Next = tmpMap[sum].Next
	}
	return result.Next
}

func main() {

}
