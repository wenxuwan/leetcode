/*该题目和0021基本类似，0021是合并两个list，这个是合并
K个list，我们自然的可以想到用分治算法，也就是不断把我们
传入的list数组切分。然后每两个list合并一次，重复这种操作，
从而得到最终的结果。

1) 该问题的规模缩小到一定的程度就可以容易地解决
2) 该问题可以分解为若干个规模较小的相同问题，即该问题具有最优子结构性质。
3) 利用该问题分解出的子问题的解可以合并为该问题的解；
4) 该问题所分解出的各个子问题是相互独立的，即子问题之间不包含公共的子子问题。
*
*/
package main

type ListNode struct {
	Val int
	Next *ListNode
 }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var tmp *ListNode
	if l1.Val < l2.Val {
		tmp = l1
		tmp.Next = mergeTwoLists(l1.Next, l2)
		return tmp
	} else {
		tmp = l2
		tmp.Next = mergeTwoLists(l1, l2.Next)
		return tmp
	}

}

func partList(lists []*ListNode,start int, end int)*ListNode{
	if start == end{
		return lists[start]
	}
	if start < end{
		mid := (start + end) / 2
		l1 := partList(lists,start,mid)
		l2 := partList(lists,mid+1,end)
		return mergeTwoLists(l1,l2)
	}
	return nil
}
func mergeKLists(lists []*ListNode) *ListNode {
	return  partList(lists,0 , len(lists) - 1)
}

func main(){
	//ListNodes := make([]*ListNode,0,10)
}
