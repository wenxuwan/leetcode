/**
该题目因为要求时间复杂度为log(n)所以，我们采用的是二分查找，但我们并不清楚
在哪里旋转了数组。所以问题关键就是怎么判断出起始坐标和终止坐标。我们看个例子
1,2,3,4,5
对于这个数组我们的旋转方式有
5,1,2,3,4
5,4,1,2,3
5,4,3,1,2
5,4,3,2,1

假设最左边为left，最右边为right

我们可以看出对于里面的任意数据,我们都可以通过与最右侧的数据比较，如果小于最右侧的
那么证明右侧是有序的！如果target在这个值和target之间那么就可以二分直接查找了。如果
target不在这个区间，那么就按照前面的步骤继续分。

*/
package main

func search(nums []int, target int) int {
	if len(nums) == 0{
		return  -1
	}
	begin := 0
	end := len(nums) - 1
	mid := 0
	for begin <= end{
		mid = (begin + end) / 2
		if nums[mid] == target{
			return mid
		}else{
			if nums[mid] < nums[end]{
				if nums[mid] < target && nums[end] >= target{
					begin = mid + 1
				}else {
					end = mid -1
				}
			}else{
				if nums[begin] <= target && nums[mid] > target{
					end = mid - 1
				}else{
					begin = mid + 1
				}
			}
		}
	}
	return  -1
}

func main(){

}
