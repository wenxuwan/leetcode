/***
该题的主要解法就是两个链表，小于x值得放到一个链表，大于x的放到一个链表
最后两个链表合并。
**/
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	list1 := new(ListNode)
	flag1 := list1
	list2 := new(ListNode)
	flag2 := list2
	p := head
	for p != nil{
		if p.Val < x{
			flag1.Next = p
			flag1 = p
		}else{
			flag2.Next = p
			flag2 = p
		}
		p = p.Next
	}
	flag1.Next = list2.Next
	flag2.Next = nil
	return list1.Next
}
func createLisNode() *ListNode{
	head := new(ListNode)
	flag := head
	for i := 0; i < 6;i++{
		tmp := new(ListNode)
		tmp.Val = i
		tmp.Next = nil
		flag.Next = tmp
		flag = tmp
	}
	flag = head.Next
	flag.Val = 1
	flag = head.Next
	flag.Val = 4
	flag = head.Next
	flag.Val = 3
	flag = head.Next
	flag.Val = 2
	flag = head.Next
	flag.Val = 5
	flag = head.Next
	flag.Val = 2
	return  head.Next
}
func main(){
	p := createLisNode()
	s := partition(p,3)
	for s != nil{
		fmt.Println(s.Val)
		s = s.Next
	}
}
