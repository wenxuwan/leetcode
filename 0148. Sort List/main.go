/*
归并排序是一个稳定算法，时间复杂度是nlogn。一开始想到的是二分法，所以自然也就想到了
归并排序。归并排序的主要算法流程：
1.拆分
2.合并

拆分是指拆分成一个个有序的数组，然后一个个合并有序的数组

比如 3，4，2，1

首先拆分成[3],[4],[2],[1]然后[3],[4]合并成[3,4]，[2],[1]合并成[1,2]然后
再合并[3,4],[1,2]。可以得到[1,2,3,4].这就是归并排序的核心原理。

*/

package main

type ListNode struct {
	Val int
	Next *ListNode
}
func merge(l1 *ListNode,l2 *ListNode) *ListNode{
	list := new(ListNode)
	tmp := list
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			tmp.Next = l1
			l1 = l1.Next
		}else{
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil{
		tmp.Next = l2
	}else{
		tmp.Next = l1
	}
	return list.Next
}
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	slow := head
	fast := slow.Next

	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil

	l1 := sortList(head)
	l2 := sortList(fast)

	return merge(l1, l2)
}

func main(){

}
