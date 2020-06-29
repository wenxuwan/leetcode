package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func help(root *TreeNode, sum, tmp int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum == tmp+root.Val {
			return true
		}
		return false
	}
	res1 := help(root.Left, sum, tmp+root.Val)
	if res1 == true {
		return true
	}
	res2 := help(root.Right, sum, tmp+root.Val)
	if res2 == true {
		return true
	}
	return false
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return help(root, sum, 0)
}

func main() {

}
