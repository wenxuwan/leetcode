/***
该题目也是比较简单的链表题目，可以三个指针分别指向前中后三个node然后比较
不一样的就可以放到新的链表之中，注意开头和结尾边界。

***/

package main

type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newList := new(ListNode)
	p := newList
	preNode := new(ListNode)
	preNode = nil
	currentNode := new(ListNode)
	nextNode := new(ListNode)
	currentNode = head
	for currentNode.Next != nil{
		nextNode = currentNode.Next
		if preNode == nil{
			if currentNode.Val != nextNode.Val{
				p.Next = currentNode
				p = currentNode
				p.Next = nil
			}
		}else {
			if currentNode.Val != nextNode.Val && currentNode.Val != preNode.Val{
				p.Next = currentNode
				p = currentNode
				p.Next = nil
			}
		}
		preNode =currentNode
		currentNode = nextNode
	}
	if  preNode.Val != currentNode.Val{
		p.Next = currentNode
		currentNode.Next = nil
	}
	return newList.Next
}
func createLisNode() *ListNode{
	head := new(ListNode)
	flag := head
	for i := 0; i < 1;i++{
		tmp := new(ListNode)
		tmp.Val = 2
		tmp.Next = nil
		flag.Next = tmp
		flag = tmp
	}
	for i := 0; i < 3;i++{
		tmp := new(ListNode)
		tmp.Val = 3
		tmp.Next = nil
		flag.Next = tmp
		flag = tmp
	}
	for i := 0; i < 1;i++{
		tmp := new(ListNode)
		tmp.Val = 5
		tmp.Next = nil
		flag.Next = tmp
		flag = tmp
	}
	return  head.Next
}

func main(){
	s := createLisNode()
	p := s
	for p != nil{
		//fmt.Println(p.Val)
		p = p.Next
	}
}