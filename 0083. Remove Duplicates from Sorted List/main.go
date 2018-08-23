package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return  head
	}
	preNode := new(ListNode)
	currenNode := new(ListNode)
	preNode = head
	currenNode = head.Next
	for currenNode != nil{
		if currenNode.Val != preNode.Val{
			preNode.Next = currenNode
			preNode = currenNode
		}
		currenNode = currenNode.Next
	}
	preNode.Next = nil
	return  head
}

func createLisNode() *ListNode{
	head := new(ListNode)
	flag := head
	for i := 0; i < 5;i++{
		tmp := new(ListNode)
		tmp.Val = i
		tmp.Next = nil
		flag.Next = tmp
		flag = tmp
	}
	return  head.Next
}

func main(){
	s := createLisNode()
	p := deleteDuplicates(s)
	for p != nil{
		fmt.Println(p.Val)
		p = p.Next
	}
}
