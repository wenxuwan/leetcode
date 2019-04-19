/*看到题目第一反应就是中序遍历一下二叉树，求出中序的序列，然后就可以用
k做下标求出值了。
 */
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
 }

 var result []int

 func treeMiddleSearch(root *TreeNode){
 	if root != nil {
		treeMiddleSearch(root.Left)
		result = append(result, root.Val)
		treeMiddleSearch(root.Right)
	}
 }
func kthSmallest(root *TreeNode, k int) int {
	result = result[0:0]
	treeMiddleSearch(root)
	fmt.Println(result)
	return result[k - 1]
}

func main(){

}