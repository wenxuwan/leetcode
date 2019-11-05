/*
题目的解题思路是比较明确的：
1.按照身高和位置排序
2.从头遍历，逐次插入

开始的时候通过切片来控制插入导致耗时很大，看题解用empty
来标识插入的位置，这种做法耗时很短。

主要思想就是用empty来标识下一次放置的位置。排序后people就会
按照身高升序排列，按照位置降序排列。比如main里面的data排序后
{{9,0},{7,0},{1,9},{3,0},{2,7},{5,3},{6,0},{3,4},{6,2},{5,2}}
[[1 9] [2 7] [3 4] [3 0] [5 3] [5 2] [6 2] [6 0] [7 0] [9 0]]

这样排列后相当于每次找出最矮的放到相应的位置，因为后面的都比它高，所以
它需要先放置位置。然后把该位置剔除掉，然后继续遍历，依次放入所有的人。
*/
package main

import (
	"fmt"
	"sort"
)

type twoDe [][]int

func (t twoDe) Len() int {
	return len(t)
}

func (t twoDe) Less(i, j int) bool {
	if t[i][0] > t[j][0] {
		return false
	}
	if t[i][0] == t[j][0]{
		if t[i][1] < t[j][1]{
			return false
		} else{
			return true
		}
	}
	return true
}

func (t twoDe) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func reconstructQueue(people [][]int) [][]int {
	result := make([][]int,len(people))
	var empty =  make([]int,len(people))
	var demo twoDe = people
	for i := 0; i < len(people); i++ {
		result[i] = make([]int, 2)
		empty[i]=i
	}
	/*
	for i := 0;i < len(people);i++{
		for j := 0; j < len(people) - 1;j++{
			if people[j][0] < people[j + 1][0]{
				people[j][0],people[j + 1][0] = people[j + 1][0],people[j][0]
				people[j][1],people[j + 1][1] = people[j + 1][1],people[j][1]
			}
			if people[j][0] == people[j + 1][0] && people[j][1] > people[j + 1][1]{
				people[j][1],people[j + 1][1] = people[j + 1][1],people[j][1]
			}
		}
	}*/
	sort.Sort(demo)
	fmt.Println(demo)
	for i := 0;i < len(demo);i++{
		/*切片处理
		tmp := make([][]int,len(result[demo[i][1]:]))
		copy(tmp,result[demo[i][1]:])
		result = append(result[0:demo[i][1]],demo[i])
		result = append(result,tmp...)*/
		k:=people[i][1]
		result[empty[k]] = people[i]
		empty = append(empty[:k], empty[k+1:]...)

	}
	fmt.Println(result[:len(demo)])
	return result[:len(demo)]
}

func main(){
	data := [][]int{{9,0},{7,0},{1,9},{3,0},{2,7},{5,3},{6,0},{3,4},{6,2},{5,2}}
	reconstructQueue(data)
}
