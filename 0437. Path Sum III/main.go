package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, sum int, result *int) {
	if root == nil {
		return
	}
	if sum-root.Val == 0 {
		*result++
	}
	helper(root.Left, sum-root.Val, result)
	helper(root.Right, sum-root.Val, result)
}

func pathSumHelper(root *TreeNode, sum int, result *int) int {
	if root == nil {
		return *result
	}
	helper(root, sum, result)
	pathSumHelper(root.Left, sum, result)
	pathSumHelper(root.Right, sum, result)
	return *result
}
func pathSum(root *TreeNode, sum int) int {
	result := 0
	pathSumHelper(root, sum, &result)
	return result
}

func main() {

}
