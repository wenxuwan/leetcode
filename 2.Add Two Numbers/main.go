package main

import "fmt"

type ListNode struct {
	Val int
	Next  *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := 0
	jinwei := 0
	l1flag := l1
	l2flag := l2
	listHead := new(ListNode)
	currentNode := listHead
	for l1flag != nil || l2flag !=nil{
		if l1flag == nil{
			sum = l2flag.Val + jinwei
		}
		if l2flag == nil{
			sum = l1flag.Val + jinwei
		}
		if l1flag != nil && l2flag != nil{
			sum = l1flag.Val + l2flag.Val + jinwei
		}
		jinwei = 0
		if sum >= 10{
			jinwei = 1
			sum = sum % 10
		}
		if l1flag != nil{
			l1flag = l1flag.Next
		}
		if l2flag != nil{
			l2flag = l2flag.Next
		}
		newLinkList := new(ListNode)
		newLinkList.Val = sum
		newLinkList.Next = nil
		currentNode.Next = newLinkList
		currentNode = newLinkList
	}

	if jinwei != 0{
		newLinkList := new(ListNode)
		newLinkList.Val = jinwei
		newLinkList.Next = nil
		currentNode.Next = newLinkList
	}
	return listHead.Next
}

func main(){
	l1 := ListNode{2,nil}
	l2 := ListNode{8,nil}
	l3 := addTwoNumbers(&l1,&l2)
	for l3 != nil{
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}


