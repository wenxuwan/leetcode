package main

import (
	"container/heap"
	"fmt"
)

// 堆的实现！！要注意的top k大用最小堆，top k小 用最大堆
type Node struct {
	key   int
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
	return (*c)[i].value < (*c)[j].value
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

/*
func topKFrequent(nums []int, k int) []int {
	Keymap := make(map[int]int)
	result := make([]int,0,len(nums))
	var m NewMap
	for i := 0;i < len(nums);i++{
		Keymap[nums[i]]++
	}
	for key,value := range Keymap{
		m1 := MapStruct{key,value}
		m = append(m,m1)
	}
	sort.Sort(m)
	for index,value := range m{
		if index == k{
			break
		}
		result = append(result,value.key)
	}
	return result
}
*/

func topKFrequent(nums []int, k int) []int {
	Keymap := make(map[int]int)
	result := make([]int, 0, k)
	for _, value := range nums {
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
	for smallHeap.Len() > 0 {
		result = append(result, smallHeap.Pop().(Node).key)
	}
	return result
}

func main() {
	s := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(s, k))
}
