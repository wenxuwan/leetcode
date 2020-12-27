package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	l := make([]*Node, 0, 0)
	l = append(l, root)
	for len(l) > 0 {
		lenth := len(l)
		for i := 0; i < lenth; i++ {
			if i == lenth-1 {
				l[i].Next = nil
			} else {
				l[i].Next = l[i+1]
			}
			if l[i].Left != nil {
				l = append(l, l[i].Left)
			}
			if l[i].Right != nil {
				l = append(l, l[i].Right)
			}
		}
		l = l[lenth:]
	}
	return root
}

func main() {

}
