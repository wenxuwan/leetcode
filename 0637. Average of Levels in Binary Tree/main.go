package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	q := list.New()
	if root == nil {
		return result
	}
	q.PushFront(root)
	for q.Len() > 0 {
		currentLenth := q.Len()
		var tmpResult float64
		var sum int
		for i := 0; i < currentLenth; i++ {
			node := q.Back()
			n, _ := node.Value.(*TreeNode)
			sum = sum + n.Val
			if n.Left != nil {
				q.PushFront(n.Left)
			}
			if n.Right != nil {
				q.PushFront(n.Right)
			}
			q.Remove(node)
		}
		tmpResult = float64(sum) / float64(currentLenth)
		result = append(result, tmpResult)
	}
	return result
}

func main() {

}
