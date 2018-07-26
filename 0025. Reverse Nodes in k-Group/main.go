/***
该题目其实可以理解为反转链表，唯一的不同就是一段段的反转，
所以自然而然的就可以想到递归了。每次都反转K长度的链表。然后
把每一段链接起来就可以了。

***/
package main

import "fmt"

type ListNode struct {
	Val int
    Next *ListNode
 }

 func reverseList(head *ListNode, k int) *ListNode{
 	if k <= 1 {
		return head
	}
	step := k
	p := head
	len := k
	for p != nil{
		p = p.Next
		len = len - 1
		if 0 == len{
			break
		}
	}
	tmp := new(ListNode)
	newList :=new(ListNode)
	newList = nil
	if len == 0{
		tmp = head
		for k > 0 && tmp != nil{
			p := tmp.Next
			tmp.Next = newList
			newList = tmp
			tmp = p
			k = k - 1
		}
		head.Next = reverseList(tmp,step)
	}else{
		newList = head
	}
	 return newList
 }
func reverseKGroup(head *ListNode, k int) *ListNode {
	newList := new(ListNode)
	newList = reverseList(head,k)
	return newList
}

func createLisNode() *ListNode{
	head := new(ListNode)
	flag := head
	for i := 0; i < 10;i++{
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
	afterReverse := reverseKGroup(s,3)
	for afterReverse != nil{
		fmt.Println(afterReverse.Val)
		afterReverse = afterReverse.Next
	}
}