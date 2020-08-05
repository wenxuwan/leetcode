package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//找到中间位置，反转后面的，然后再遍历一次
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	newList := new(ListNode)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		tmp := newList.Next
		newList.Next = slow
		tmp2 := slow.Next
		slow.Next = tmp
		slow = tmp2
	}
	newList = newList.Next
	for newList != nil {
		if head.Val != newList.Val {
			return false
		}
		newList = newList.Next
		head = head.Next
	}
	return true
}

func main() {

}
