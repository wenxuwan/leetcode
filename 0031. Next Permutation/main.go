/*
这个题目的做法就是找到下一个比较大的序列，因为是下一个比较大的
所以需要从后面找，找到第一个后面数字比前面数字大的地方，然后
互换位置，然后再把后面的数字排列一下。
eg:
43231
从后面找到2和3
然后互换2，3位置变为
34321
然后把后面的21排序变为12
所以就是34312
*/
package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int)  {
	lenth := len(nums)
	flag := 0
	j := 0
	for i := lenth - 2; i >= 0;i--{
		if nums[i] < nums[i + 1]{
			for j = i + 1;j < lenth;j++{
				if nums[j] > nums[i]{
					continue
				}else{
					break
				}
			}
			nums[j - 1],nums[i] = nums[i], nums[j - 1]
			flag = i + 1
			break
		}
	}
	sort.Ints(nums[flag:])
}

func main(){
	s := []int{1}
	nextPermutation(s)
	fmt.Println(s)
}
