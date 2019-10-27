/*
拿到题目首先分析，如果所有的数据的和为奇数，那么肯定没法划分成两份。
如果是偶数，就有两种可能，一种就是可以分成两半，一种是不可以分为两半。偶数的一半就是half = sum / 2

然后继续分解问题，变为在数组中查找组合的和等于half的组合！！这个就回到了我们熟悉的递归了。

首先排序一下，为了后面的数据好判断。然后判断最大的数据如果大于half那直接return false，然后递归的时候
从后面往前走。先从大得开始，有助于寻找。

然后就是递归求值了。

作者：fish_of_cat
链接：https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/goyu-yan-zhan-sheng-bai-fen-bai-by-fish_of_cat/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"sort"
)

func helper(nums []int, index int, sum int)bool{
	if index < 0 || sum < 0{
		return false
	}
	if sum == 0{
		return true
	}
	for i := index; i > 0;i--{
		result := helper(nums, i - 1,sum - nums[i])
		if result == true{
			return true
		}
	}
	return false
}

func canPartition(nums []int) bool {
	sum := 0
	for _, value := range(nums){
		sum += value
	}
	if sum % 2 != 0{
		return false
	}
	half := sum / 2
	//fmt.Println(half)
	sort.Ints(nums)
	if nums[len(nums) - 1] > half {
		return false
	}
	return helper(nums, len(nums) - 1, half)
}

func main(){
	//data := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,100}
	data := []int{1,1,2,5,5,5,5}
	fmt.Println(canPartition(data))
}
