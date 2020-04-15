/*
利用栈

新元素和栈顶元素比较，如果新元素小，那么就pop出来栈顶元素，这个地方要注意的就是k不要超了。

*/
package main

import (
	"fmt"
	"sync"
)

type (
	Stack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	flag := 0
	for i := 0; i < length; i++ {
		if tmp[length-i-1] == 48 && flag == 0{
			continue
		}
		flag = 1
		result = append(result, tmp[length-i-1])
	}
	if len(result) == 0{
		return "0"
	}
	return string(result)
}

func removeKdigits(num string, k int) string {
	st := NewStack()
	l := len(num)
	result := ""
	removeNum := 0
	if k >= l{
		return "0"
	}
	for _, value := range num{
		for  removeNum < k  && st.length != 0 && value < st.Peek().(int32){
			st.Pop()
			removeNum++
		}
		st.Push(value)
	}

	if st.length == 0{
		return "0"
	}
	for st.length > l - k{
		st.Pop()
	}
	for st.length > 0{
		result += string(st.Pop().(int32))
	}
	return reverse1(result)
}

func main(){
	num := "112"
	k := 1
	fmt.Print(removeKdigits(num, k))
}
