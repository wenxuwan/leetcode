package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func LenList(l *ListNode) int {
	result := 0
	for l != nil {
		result++
		l = l.Next
	}
	return result
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := LenList(l1)
	len2 := LenList(l2)
	l := len1
	fast := l1
	slow := l2
	t := len1 - len2
	result := new(ListNode)
	if len1 < len2 {
		fast = l2
		slow = l1
		t = len2 - len1
		l = len2
	}
	tmp := make([]int, l)
	index := 0
	for fast != nil {
		if index < t {
			tmp[index] = fast.Val
		} else {
			tmp[index] = fast.Val + slow.Val
			slow = slow.Next
		}
		fast = fast.Next
		index++
	}
	for i := l - 1; i > 0; i-- {
		value := tmp[i]
		tmp[i] = tmp[i] % 10
		tmp[i-1] += int(value / 10)
	}
	result.Val = tmp[0] / 10
	tmp[0] = tmp[0] % 10
	location := result
	for _, value := range tmp {
		tmpNode := new(ListNode)
		tmpNode.Val = value
		location.Next = tmpNode
		location = tmpNode
	}
	if result.Val != 0 {
		return result
	}
	return result.Next
}

func main() {

}
