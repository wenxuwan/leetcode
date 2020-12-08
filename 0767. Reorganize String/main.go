package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	key   byte
	value int
}

type NewHeap []Node

func (c *NewHeap) Len() int {
	return len(*c)
}

func (c *NewHeap) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *NewHeap) Less(i, j int) bool {
	if (*c)[i].value > (*c)[j].value {
		return true
	}
	return false
}

func (c *NewHeap) Push(x interface{}) {
	*c = append(*c, x.(Node))
}

func (c *NewHeap) Pop() interface{} {
	n := len(*c) - 1
	x := (*c)[n]
	*c = (*c)[:n]
	return x
}
func (c *NewHeap) Peek() Node {
	return (*c)[0]
}

func reorganizeString(S string) string {
	result := make([]byte, len(S), len(S))
	index := 0
	data := []byte(S)
	var array [26]int
	for _, value := range data {
		array[value-'a']++
	}
	bigHeap := &NewHeap{}
	for key, value := range array {
		if value == 0 {
			continue
		}
		heap.Push(bigHeap, Node{byte(key + 'a'), value})
	}
	var preNode Node
	for bigHeap.Len() != 0 {
		maxNode := heap.Pop(bigHeap).(Node)
		if maxNode == preNode {
			newMaxNode := heap.Pop(bigHeap).(Node)
			heap.Push(bigHeap, maxNode)
			maxNode = newMaxNode
		}
		if maxNode.value > (len(S)+1)/2 {
			return ""
		}
		result[index] = maxNode.key
		index++
		maxNode.value--
		if bigHeap.Len() == 0 {
			break
		}
		secondNode := heap.Pop(bigHeap).(Node)
		result[index] = secondNode.key
		index++
		secondNode.value--
		preNode = secondNode
		if maxNode.value != 0 {
			heap.Push(bigHeap, maxNode)
		}
		if secondNode.value != 0 {
			heap.Push(bigHeap, secondNode)
		}
	}
	return string(result)
}

func main() {
	S := "aaab"
	fmt.Println(reorganizeString(S))
}
