package main

import (
	"container/list"
	"fmt"
	"math"
)

func decodeString(s string) string {
	var result []int32
	l := list.New()
	for _, value := range s {
		if value != ']' {
			l.PushBack(value)
		} else {
			var tmp []int32
			var number int32
			t := (l.Back().Value).(int32)
			l.Remove(l.Back())
			for t != '[' {
				tmp = append(tmp, t)
				t = (l.Back().Value).(int32)
				l.Remove(l.Back())
			}
			t = (l.Back().Value).(int32)
			l.Remove(l.Back())
			llenth := l.Len()
			number = t - '0'
			for m := 0; m < llenth; m++ {
				t = (l.Back().Value).(int32)
				if '0' <= t && t <= '9' {
					n := math.Pow(10, float64(m+1))
					number = (t-'0')*int32(n) + number
					l.Remove(l.Back())
					continue
				}
				break
			}
			for m := 0; m < len(tmp)/2; m++ {
				tmp[m], tmp[len(tmp)-1-m] = tmp[len(tmp)-1-m], tmp[m]
			}
			for j := 0; j < int(number); j++ {
				for _, v := range tmp {
					l.PushBack(v)
				}
			}
		}
	}
	lenth := l.Len()
	for j := 0; j < lenth; j++ {
		element := l.Front()
		l.Remove(element)
		result = append(result, element.Value.(int32))
	}
	return string(result)
}

func main() {
	fmt.Println(decodeString("100[leetcode]"))
}
