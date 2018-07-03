/***
该题最笨的解法，就是遍历一下链表，得到链表的长度，然后再遍历到len-N的地方
把该元素去掉。
第二种做法就是两个指针，一前以后，中间距离n个node。当第二个指针的Next等于nil
的时候前面的指针的位置的Next就是要去掉的node


**/
package main

import "fmt"

type ListNode struct {
      Val int
      Next *ListNode
 }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := new(ListNode)
	fast := new(ListNode)
	slow = head
	fast = head
	for i:=0;i < n;i++{
		fast = fast.Next
	}
	if fast == nil{
		head = head.Next
		return head
	}
	for fast.Next != nil{
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func createList(head *ListNode, listlen int){
	head.Next = nil
	head.Val = 0
	p := head
	for  i:= 0;i < listlen;i++{
		node := new(ListNode)
		node.Val = i
		node.Next = nil
		p.Next = node
		p = node
	}
}
func main(){
	var head * ListNode
	head = new(ListNode)
	createList(head,1)
	removeNthFromEnd(head,1)
	for p := head;p !=nil;p=p.Next{
		fmt.Print(p.Val)
	}
}
