/*
这个题目其实用很多方法都可以做，但因为牵扯到了增加和删除操作，所以
链表对于这种动态的增加删除肯定是最快的。所以采用双向链表。

但如果只是采用链表，每次的查找还是避免不了的，n（1）很难实现，所以这时候
就需要用到字典。用字典实现查询。字典的key是我们传入的key，字典value是
一个node节点。
*/
package main

import "fmt"

type node struct {
	key int
	value int
	next *node
	pre *node
}
type LRUCache struct {
	capacity int
	head *node
	tail *node
	hashMap map[int]*node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.capacity = capacity
	l.head = &node{value:-1}
	l.tail = &node{value:-1}
	l.head.next = l.tail
	l.tail.pre = l.head
	l.hashMap = make(map[int]*node)
	return l
}


func (this *LRUCache) Get(key int) int {
	n := this.hashMap[key]
	if n == nil{
		return -1
	}
	//remove the node from list
	n.pre.next = n.next
	n.next.pre = n.pre
	//add the node to header
	n.next = this.head.next
	n.pre = this.head
	this.head.next.pre = n
	this.head.next = n
	return n.value
}


func (this *LRUCache) Put(key int, value int)  {
	if this.capacity < 1{
		return
	}
	n := this.hashMap[key]
	if n != nil{
		n.value =value
		//remove the node from list
		n.pre.next = n.next
		n.next.pre = n.pre
		//add the node to header
		n.next = this.head.next
		n.pre = this.head
		this.head.next.pre = n
		this.head.next = n
	}else{
		if this.capacity == len(this.hashMap){
			tmp := this.tail.pre
			this.tail.pre = tmp.pre
			tmp.pre.next = this.tail

			fmt.Println("remove old node",tmp.key,tmp.value)
			delete(this.hashMap,tmp.key)
		}
		n = &node{value: value,key: key}
		this.hashMap[key] = n
		//add the node to header
		n.next = this.head.next
		n.pre = this.head
		this.head.next.pre = n
		this.head.next = n
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main(){

}