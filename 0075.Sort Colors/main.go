/*
题目的最简单的做法就是扫描，然后看0,1,2的个数

或者是设置两个标志位red 和blue分别指向头尾，然后遍历数组

这里要注意一个问题，找到0或者2的时候Index是不能变得，因为要
比较两次，也就是你换过来的数要和0和2都比较，只有既不等于0也不等于
2的时候才能将index加一。

blue和red只要发生互换就可以自增了。

*/


package main

import "fmt"

func sortColors(nums []int)  {
	red := 0
	blue := len(nums) - 1
	for index := 0; index <= blue;{
		if nums[index] == 0 && index != red{
			nums[index],nums[red] = nums[red],nums[index]
			red++
		}else if nums[index] == 2{
			nums[index],nums[blue] = nums[blue],nums[index]
			blue--
		}else {
			index++
		}
	}
	fmt.Println(nums)
}

func main(){
	s := []int{2,0,1}
	sortColors(s)
}