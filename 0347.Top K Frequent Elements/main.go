package main

import (
	"sort"
	"fmt"
)

type NewMap []MapStruct

func (a NewMap) Len() int {
	return len(a)
}

func (c NewMap) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c NewMap) Less(i, j int) bool {
	return c[i].value > c[j].value
}
type MapStruct struct {
	key int
	value int
}

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

func main(){
	s  := []int{1,1,1,2,2,3}
	k := 2
	fmt.Println(topKFrequent(s,k))
}