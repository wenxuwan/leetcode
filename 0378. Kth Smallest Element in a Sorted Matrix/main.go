package main

import "container/heap"

type BigHeap []int

func (c *BigHeap) Len() int {
	return len(*c)
}

func (c *BigHeap) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *BigHeap) Less(i, j int) bool {
	return (*c)[i] > (*c)[j]
}

func (c *BigHeap) Push(x interface{}) {
	*c = append(*c, x.(int))
}

func (c *BigHeap) Pop() interface{} {
	n := len(*c) - 1
	x := (*c)[n]
	*c = (*c)[:n]
	return x
}
func (c *BigHeap) Peek() int {
	return (*c)[0]
}

func kthSmallest(matrix [][]int, k int) int {
	row := len(matrix)
	column := len(matrix[0])
	bigHeap := &BigHeap{}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if bigHeap.Len() == k && matrix[i][j] > bigHeap.Peek() {
				continue
			}
			heap.Push(bigHeap, matrix[i][j])
			if bigHeap.Len() > k {
				heap.Pop(bigHeap)
			}
		}
	}
	return bigHeap.Peek()
}
