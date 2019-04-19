package main

import "fmt"

type ListNode struct {
      Val int
      Next *ListNode
}

func reorderList(head *ListNode)  {
	slow := head
	fast := head
	preslow := head
	p := head
	if head == nil{
		return
	}
	for fast != nil && fast.Next != nil{
		preslow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast == head{
		return
	}
	preslow.Next = nil
	newHead := new(ListNode)
	newHead = nil
	for slow != nil{
		fmt.Println(slow.Val)
		tmp := slow.Next
		slow.Next = newHead
		newHead=slow
		slow = tmp
	}

	for p.Next != nil{
		//fmt.Println(p.Val,newHead.Val)
		tmp := p.Next
		tmp2 := newHead.Next
		p.Next = newHead
		p.Next.Next = tmp
		p = tmp
		newHead = tmp2
	}
	//fmt.Println(newHead.Val)
	if newHead != nil{
		p.Next = newHead
	}

}
