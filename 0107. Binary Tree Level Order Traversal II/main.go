package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	q := list.New()
	if root == nil {
		return result
	}
	q.PushFront(root)
	for q.Len() > 0 {
		currentLenth := q.Len()
		tmpResul := make([]int, 0, 0)
		for i := 0; i < currentLenth; i++ {
			node := q.Back()
			n, _ := node.Value.(*TreeNode)
			tmpResul = append(tmpResul, n.Val)
			if n.Left != nil {
				q.PushFront(n.Left)
			}
			if n.Right != nil {
				q.PushFront(n.Right)
			}
			q.Remove(node)
		}
		result = append(result, tmpResul)
	}
	l := len(result) - 1
	for i := 0; i <= l/2; i++ {
		result[i], result[l-i] = result[l-i], result[i]
	}
	return result
}

func main() {

}
