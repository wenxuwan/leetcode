/*
直接把两个链表分开，最后合并到一起就行
*/
package main

type ListNode struct {
	Val int
	Next *ListNode
 }
func oddEvenList(head *ListNode) *ListNode {
	newOushuHead := new(ListNode)
	//newOushuHead = nil
	list1 := newOushuHead
	newJishuHead := new(ListNode)
	//newJishushuHead = nil
	list2 := newJishuHead
	currentNode := new(ListNode)
	currentNode = head
	index := 0
	for currentNode != nil{
		if index % 2 == 0{
			list1.Next = currentNode
			list1 = currentNode
			//list1.Next = nil
		}else{
			list2.Next = currentNode
			list2 = currentNode
			//list2.Next = nil
		}
		currentNode = currentNode.Next
		index++
	}
	list1.Next = newJishuHead.Next
	list2.Next = nil
	return newOushuHead.Next
}

func main(){

}