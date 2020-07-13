package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Node struct {
	key   string
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
	if (*c)[i].value < (*c)[j].value {
		return true
	}
	if (*c)[i].value == (*c)[j].value {
		return (*c)[i].key > (*c)[j].key
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

func topKFrequent(words []string, k int) []string {
	Keymap := make(map[string]int)
	result := make([]string, 0, k)
	for _, value := range words {
		Keymap[value]++
	}
	smallHeap := &NewHeap{}
	for key, value := range Keymap {
		if smallHeap.Len() == k && value < smallHeap.Peek().value {
			continue
		}
		heap.Push(smallHeap, Node{key, value})
		if smallHeap.Len() > k {
			heap.Pop(smallHeap)
		}
	}
	sort.Sort(smallHeap)
	fmt.Println(smallHeap)
	for i := smallHeap.Len() - 1; i >= 0; i-- {
		result = append(result, (*smallHeap)[i].key)
	}
	return result
}

func main() {
	data := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(data, 4))
}
