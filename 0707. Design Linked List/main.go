package main

import "fmt"

type MyLinkedList struct {
	Val int
	Next *MyLinkedList
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := MyLinkedList{}
	head.Next = nil
	return head
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	this = this.Next
	if index < 0{
		return -1
	}
	for i := 0; i < index && this != nil;i++{
		//fmt.Println(this.Val)
		this = this.Next

	}
	if this == nil{
		return -1
	}
	return this.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	node := new(MyLinkedList)
	node.Val = val
	node.Next = this.Next
	this.Next = node
	//fmt.Println("hahahah",this.Val)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	node := new(MyLinkedList)
	node.Val = val
	node.Next = nil
	for this.Next != nil{
		this = this.Next
	}
	this.Next = node
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int){
	preNode := new(MyLinkedList)
	node := new(MyLinkedList)
	i := 0
	for i = 0; i < index + 1 && this != nil;i++{
		preNode = this
		this = this.Next
	}
	if this == nil && i != index + 1{
		return
	}
	//fmt.Println(i, node.Val)
	node.Val = val
	node.Next = preNode.Next
	preNode.Next = node
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	preNode := new(MyLinkedList)
	for i := 0;i <= index && this != nil; i++{
		preNode = this
		this = this.Next
	}
	if this != nil{
		preNode.Next = this.Next
	}
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main(){
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1))
	obj.AddAtIndex(1,2)
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(0))
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	//fmt.Println(obj.Next.Val)
	//obj.AddAtIndex(1,2)
	//fmt.Println(obj.Val)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(0))
	//obj.Get(2)
}
