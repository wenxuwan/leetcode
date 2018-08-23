/*
三次交换
第一次交换 (0 ,len - k -1)
第二次交换 (len - k ,len(nums) - 1)
第三次交换 （0，len(nums)）
1,2,3,4,5  k = 2
3,2,1,4,5
3,2,1,5,4
4,5,1,2,3

其实原理就是（0，len-k-1）区间的数值顺序不会乱
只是往后移动，但(len - k ,len(nums) - 1)的数据
会和前面的互换，所以反转三次
*/
package main

import "fmt"

func revert(nums []int, begin int, end int){
	i:= begin
	j :=end
	for i < j{
		nums[i],nums[j] = nums[j],nums[i]
		i++
		j--
	}
}
func rotate(nums []int, k int)  {
	k = k % len(nums)
	revert(nums,0,len(nums) - k - 1)
	revert(nums,len(nums) - k,len(nums) - 1)
	revert(nums,0,len(nums) - 1)
}

func main(){
	testDate := []int{1,2,3,4,5,6,7}
	rotate(testDate,2)
	fmt.Println(testDate)
}
