package main

import (
	"fmt"
	"math"
)
/*暴力解法，属于O(n^2)的时间复杂度
func minSubArrayLen(s int, nums []int) int {
	sum := 0
	begin := 0
	end := 0
	maxLenth := 0
	index :=0
	for index = 0; index < len(nums);index++{
		sum += nums[index]
		if sum >= s{
			end = index
			break
		}
	}
	if index == len(nums){
		return 0
	}
	maxLenth = index + 1

	for end < len(nums){
		//fmt.Println("sum:", sum, begin, end)
		if sum > s{
			for begin < end{
				sum -= nums[begin]
				begin++
				//fmt.Println(begin, end, sum)
				if sum >= s{
					//fmt.Println(begin, end, sum)
					if end - begin + 1< maxLenth{
						maxLenth = end - begin + 1
					}
				}else{
					break
				}
			}
		}
		end++
		if end < len(nums){
			sum += nums[end]
		}
	}
	return maxLenth
}
*/

func minSubArrayLen(s int, nums []int) int {
	begin := 0
	end := 0
	sum := 0
	minLen := math.MaxInt32
	for end < len(nums){
		if(sum + nums[end] < s){
			sum += nums[end]
			end++
		}else{
			if end - begin < minLen{
				minLen = end - begin + 1
			}
			sum = sum - nums[begin]
			begin++
		}
	}
	if minLen == math.MaxInt32{
		return 0
	}
	return minLen
}
func main(){
	//nums := []int{5,1,3,5,10,7,4,9,2,8}
	//nums := []int{2, 3, 1, 2, 4, 3}
	nums := []int{1,2,3,4}
	fmt.Println(minSubArrayLen(10,nums))
}
