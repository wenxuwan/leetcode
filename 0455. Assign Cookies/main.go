package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	result := 0
	for _, value := range g{
		for j, cake := range s{
			if cake >= value{
				result++
				s = s[j+1:len(s)]
				break
			}
		}
	}
	return result
}

func main(){
	g := []int{10,9,8,7}
	s := []int{5,6,7,8}
	fmt.Println(findContentChildren(g,s))
}
