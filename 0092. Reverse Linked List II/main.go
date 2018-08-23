/*
该题找到开始开始的m的前一个 prenode，然后从将m 到 n 的子链表倒置
然后将prenode的next指向倒置后的链表，倒置后的链表的next指向n后面的
节点
*/

package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	p := head
	sublink := new(ListNode)
	sublink = nil
	preLink := new(ListNode)
	preLink = nil
	i := 0
	for p != nil{
		i++
		if i == m - 1{
			preLink = p
		}
		if i >= m && i <= n{
			next := p.Next
			p.Next = sublink
			sublink = p
			p = next
			continue
		}
		if i > n{
			break
		}
		p = p.Next
	}
	if preLink != nil{
		preLink.Next.Next = p
		preLink.Next = sublink
	}else{
		head.Next = p
		head = sublink
	}
	return  head
}