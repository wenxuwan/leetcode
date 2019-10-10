/*
树的问题还是要用递归来解决，首先拆分任务，对于每一个节点都有如下的情况：

1.节点位于自己两边，那么return自己
2.节点位于自己右边，return右边
3.节点位于自己左边，return左边
*/
package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || p.Val == root.Val || q.Val == root.Val {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p , q)
    if left != nil && right != nil{
        return root
    }
    if left != nil{
        return left
    }
    if right != nil{
        return  right
    }
    return  nil
}
func main(){

}
