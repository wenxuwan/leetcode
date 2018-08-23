package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return  head
	}
	p := new(ListNode)
	p = head
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

func isPalindrome(head *ListNode) bool {
	if head == nil{
		return true
	}
	slow := head
	fast := head.Next
	for fast != nil && slow != nil{
		slow = slow.Next
		fast = fast.Next
		if fast != nil{
			fast = fast.Next
		}
	}
	newList := reverseList(slow.Next)
	slow.Next = nil
	for newList != nil && head != nil{
		if newList.Val != head.Val{
			return  false
		}else{
			head = head.Next
			newList = newList.Next
		}
	}
	if head != nil{
		if head.Next == nil{
			return true
		}
		if head.Next.Val == head.Val{
			return true
		}else{
			return false
		}
	}
	return true
}

func createLisNode() *ListNode{
	head := new(ListNode)
	flag := head
	for i := 0; i < 1;i++{
		tmp := new(ListNode)
		tmp.Val = i % 3
		tmp.Next = nil
		flag.Next = tmp
		flag = tmp
	}
	return  head.Next
}

func main(){
	list := createLisNode()
	p := list
	for p != nil{
		fmt.Println(p.Val)
		p = p.Next
	}
	fmt.Println(isPalindrome(list))
}