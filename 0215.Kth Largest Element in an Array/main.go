package main

import (
	"fmt"
	"container/heap"
)

/*
冒泡一下
*/
func findKthLargest(nums []int, k int) int {
	flag := true
	for i := 0; i < k;i++{
		for j := 0; j < len(nums) - i - 1;j++{
			if nums[j] > nums[j + 1]{
				flag = false
				nums[j],nums[j + 1] = nums[j + 1],nums[j]
			}
		}
		if flag == true{
			return nums[len(nums) - k]
		}
	}
	fmt.Println(nums)
	return nums[len(nums) - k]
}

type maxHeap struct {
	data []int
}

func (h *maxHeap) Len() int {
	return len(h.data)
}

func (h *maxHeap) Less(i, j int) bool {
	return h.data[i] > h.data[j]
}

func (h *maxHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *maxHeap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	top := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return top
}

func findKthLargest(nums []int, k int) int {
	var res int
	h := new(maxHeap)
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
	}

	for i := 0; i < k; i++ {
		res = heap.Pop(h).(int)
	}

	return res
}


func main(){

}