package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
	issame := true
	if p == nil && q == nil{
		return true
	}
	if p == nil && q != nil{
		return false
	}
	if p != nil && q == nil{
		return false
	}
	if p.Val != q.Val{
		return false
	}
	issame = issame && isSameTree(p.Right,q.Right)
	issame = issame && isSameTree(p.Left,q.Left)
	return issame
	}

func main(){

}
