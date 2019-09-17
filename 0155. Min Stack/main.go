package main

import (
	"fmt"
)

type MinStack struct {
	val int
	next *MinStack
	small *MinStack
	big *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var m MinStack
	//m.big = new(MinStack)
	//m.small = new(MinStack)
	return m
}


func (this *MinStack) Push(x int)  {
	node := new(MinStack)
	node.val = x
	node.next = this.next
	this.next = node

	tmp := this.big
	if tmp == nil{
		this.big = node
		node.small = this
		node.big = nil
		return
	}
	for tmp != nil{
		if tmp.val > x{
			tmp.small.big = node
			node.small = tmp.small
			node.big = tmp
			tmp.small = node
			return
		}else{
			if tmp.big == nil{
				//fmt.Println("heheh",x,tmp.val)
				tmp.big = node
				node.small = tmp
				node.big = nil
				return
			}
		}
		tmp = tmp.big
	}
}


func (this *MinStack) Pop()  {
	this.next.small.big = this.next.big
	if this.next.big != nil{
		this.next.big.small = this.next.small
	}
	this.next = this.next.next
}


func (this *MinStack) Top() int {
	return this.next.val
}


func (this *MinStack) GetMin() int {
	return this.big.val
}

func (this *MinStack) GetAll() {
	tmp := this.big
	for tmp != nil{
		fmt.Println("---",tmp.val)
		tmp = tmp.big
	}
}
func main(){
	var m MinStack
	m.Push(2)
	m.Push(0)
	m.Push(3)
	m.Push(0)
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.GetMin())
	m.GetAll()
	m.Pop()

	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.GetMin())
}