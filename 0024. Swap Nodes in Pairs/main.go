/*
该题目是反转相邻的两个节点，最主要的是不允许改变值，
而是直接改变
指针。
|1|2|3|4|5|6|7|
|h|t|h|t|h|t|h|

每次传入head，我们让temp指向head.next。然后下次的
head变成temp.next，继续。然后返回本次的temp到上次
head.next。

*/
package main

type ListNode struct {
	Val int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	if head.Next == nil{
		return head
	}
	temp := head.Next
	head.Next = swapPairs(temp.Next)
	temp.Next = head
	return temp
}
func main(){
}
