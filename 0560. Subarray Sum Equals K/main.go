package main

/*
暴力解，用两个for遍历所有的可能性。从i开始一直加到最后，看
有多少个解
*/
func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums);i++{
		sum := 0
		for j := i; j < len(nums);j++{
			sum += nums[j]
			if sum == k{
				count++
			}
		}
	}
	return count
}

/*
这种做法是通过哈希表来做
定义一个sum字典，该字典以和为key，以出现的次数为value。

该题的重点在于连续：

1.假设 sum[i]代表的是0-i的所有元素的和，那么sum[j]代表的
就是0-j所有元素的和。
2.如果sum[j] - sum[i] == k，也就说i - j之间的和为k
3.可能出现多个序列的值相同，所以用value来保存出现的子序列的个数。

*/
func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	sumMap := make(map[int]int)
	sumMap[0] = 1
	for i := 0; i < len(nums);i++{
		sum += nums[i]
		if _, ok := sumMap[sum - k];ok{
			count = count + sumMap[sum - k]
		}
		sumMap[sum]++
	}
	return count
}

func main(){

}
