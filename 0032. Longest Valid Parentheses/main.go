package main

import (
	"container/list"
	"fmt"
)

type ValueStu struct {
	index int
	value int32
}

func longestValidParentheses(s string) int {
	l := list.New()
	resultList := make([]int, len(s), len(s))
	for index, value := range s {
		v := &ValueStu{}
		v.value = value
		v.index = index
		if value == '(' {
			l.PushBack(v)
		} else {
			if l.Back() != nil {
				if l.Back().Value.(*ValueStu).value == '(' {
					va := l.Back().Value.(*ValueStu)
					resultList[index] = 1
					resultList[va.index] = 1
					l.Remove(l.Back())
				} else {
					l.PushBack(v)
				}
			} else {
				l.PushBack(v)
			}
		}
	}
	tmpMax := 0
	max := 0
	for _, value := range resultList {
		if value == 1 {
			tmpMax++
		} else {
			tmpMax = 0
		}
		if tmpMax > max {
			max = tmpMax
		}
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
