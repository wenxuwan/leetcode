package main

import (
	"fmt"
)

type Node struct{
	val int
	left_child *Node
	right_child *Node
}

type QueueNode struct {
	node *Node
	next *QueueNode
}
type Queue struct {
	head *QueueNode
	tail *QueueNode
	len int
}

func (this *Queue)Push(node *QueueNode){
	if this.head == nil && this.tail == nil{
		this.head = node
		this.tail = node
	}else{
		this.tail.next = node
		this.tail = node
	}
	this.len++
}
func (this *Queue)Pop()*QueueNode{
	if this.Len() == 0{
		return nil
	}
	//fmt.Println("Test",this.head.node.val)
	tmp := this.head
	this.head = this.head.next
	if this.head == nil{
		this.tail = nil
	}
	this.len--
	return tmp
}
func (this *Queue)Len()int{

	return this.len
}
func create_new_node(value int)* Node{
	new_node := new(Node)
	new_node.right_child = nil
	new_node.left_child = nil
	new_node.val = value
	return new_node
}

//中序遍历对于搜索树来说输出的就是有序的数据
func search_tree_ldr(root *Node){
	if root == nil{
		return
	}
	search_tree_ldr(root.left_child)
	fmt.Print(root.val, ",")
	search_tree_ldr(root.right_child)
}

//水平遍历树，对于水平遍历二叉树，需要利用的就是队列的特性，先进先出的特性
func search_tree_with_row(root *Node){
	if root == nil{
		return
	}
	queue := new(Queue)
	queue.head = nil
	queue.tail = nil
	node := QueueNode{root,nil}
	queue.Push(&node)
	for queue.Len() != 0{
		tmpNode := queue.Pop().node
		fmt.Print(tmpNode.val,",")
		if tmpNode.left_child != nil{
			tmp_node := QueueNode{tmpNode.left_child,nil}
			queue.Push(&tmp_node)
		}
		if tmpNode.right_child != nil{
			tmp_node := QueueNode{tmpNode.right_child,nil}
			queue.Push(&tmp_node)
		}
	}
}

//常见面试题目之水平遍历树，然后分层打印
func search_tree_with_row_and_print(root *Node){
	if root == nil{
		return
	}
	queue := new(Queue)
	queue.head = nil
	queue.tail = nil
	node := QueueNode{root,nil}
	queue.Push(&node)
	for queue.Len() != 0{
		len := queue.Len()
		for i := 0;i < len;i++{
			tmpNode := queue.Pop().node
			fmt.Print(tmpNode.val,",")
			if tmpNode.left_child != nil{
				tmp_node := QueueNode{tmpNode.left_child,nil}
				queue.Push(&tmp_node)
			}
			if tmpNode.right_child != nil{
				tmp_node := QueueNode{tmpNode.right_child,nil}
				queue.Push(&tmp_node)
			}
		}
		fmt.Println()
	}
}

//建立搜索二叉树
func buildSearchTree(root **Node, value int){
	tmp_node := *root
	pre_node := new(Node)
	pre_node = nil
	if *root == nil{
		*root = create_new_node(value)
		//fmt.Println(root)
	}else{
		for tmp_node != nil{
			pre_node = tmp_node
			if tmp_node.val > value{
				tmp_node = tmp_node.left_child
				if tmp_node == nil{
					pre_node.left_child = create_new_node(value)
				}
			}else{
				tmp_node = tmp_node.right_child
				if tmp_node == nil{
					pre_node.right_child = create_new_node(value)
				}
			}
		}
	}
}
func main(){
	val := []int{100,5,6,1,3,4}
	root := new(Node)
	root = nil
	for i := 0;i < len(val);i++{
		buildSearchTree(&root, val[i])
	}
	//fmt.Println(root.val)
	fmt.Println("中序遍历搜索二叉树")
	search_tree_ldr(root)
	fmt.Println()
	fmt.Println("水平遍历搜索二叉树:")
	search_tree_with_row(root)
	fmt.Println()
	fmt.Println("水平遍历搜索二叉树,分层打印:")
	search_tree_with_row_and_print(root)
}
