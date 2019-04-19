/*****
和动态递归有点像，主要就是保存两个值，通过更新两个值来得到最后结果

当a[i] =< first的时候更新first，当second >= a[i]的时候更新second，当
两种都不满足也就是second < a[i]的时候就找到了第三个。

 */

package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func increasingTriplet(nums []int) bool {
	if len(nums) < 3{
		return  false
	}
	first, second := INT_MAX, INT_MAX
	for i := 0;i < len(nums);i++{
		if first >= nums[i]{
			first = nums[i]
		}else if second >= nums[i]{
			second = nums[i]
		}else if nums[i] > second{
			return true
		}
	}
	return false
}

func main(){
	pre := []int{1,1,-2,6}
	fmt.Println(increasingTriplet(pre))
}
