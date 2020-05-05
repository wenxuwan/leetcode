package main

import (
	"fmt"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	flag1 := 0
	len1 := len(a)
	flag2 := 0
	len2 := len(b)
	result := 2147483649
	for flag1 != len1 && flag2 != len2{
		tmp := a[flag1] - b[flag2]
		if tmp < 0{
			tmp = -tmp
		}
		if tmp < result{
			result = tmp
		}
		if tmp == 0{
			flag1++
			flag2++
		}
		if a[flag1] > b[flag2]{
			flag2++
		}else{
			flag1++
		}
	}
	return result
}

func main(){
	a := []int{0}
	b := []int{2147483647}
	fmt.Print(smallestDifference(a, b))
}