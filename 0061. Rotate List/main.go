/***
该题目思路比较简单，主要是弄清楚旋转之后的头从哪开始，然后把新的头前面的元素的
Next指向Nil，把原先的尾指向head。

两种情况：
1.步数(k)小于链表的长度(len)：
	1.1.新的head的坐标：
		len - k
2.步数(k)大于等于链表的长度（len）:
	2.1: k % len == 0 , return head directly
	2.2: k % len != 0, len - (k % len)
剩下的就容易解决了。
***/

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0{
		return head
	}
	listLen := 0
	p := head
	lastNode := new(ListNode)
	newList := new(ListNode)
	for p != nil{
		p = p.Next
		listLen += 1
	}
	p = head
	lastIndex := 0
	if k % listLen == 0{
		return head
	}
	if k < listLen{
		lastIndex = listLen - k
	}else{
		lastIndex = listLen - (k % listLen)
	}
	for i := 1; i < listLen;i++{
		if i == lastIndex{
			newList = p.Next
			lastNode = p
		}
		p = p.Next
	}
	p.Next = head
	lastNode.Next = nil
	return newList
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
	he := rotateRight(s,2)
	for he != nil{
		fmt.Println(he.Val)
		he = he.Next
	}
}