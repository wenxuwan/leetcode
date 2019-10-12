/*
时间复杂度 O（n）
空间复杂度 O（1）

遍历数组，用三个变量分别记录该变量出现的次数，往前移动的下表，以及当前正在统计的value值

*/

package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	currentIndex := 0
	currentValue := nums[0]
	currentCount := 0
	for _, value := range nums{
		if currentValue == value{
			currentCount++
			if currentCount > 2 {
				//DO nothing
			}else{
				nums[currentIndex] = value
				currentIndex++
			}
		}else{
			nums[currentIndex] = value
			currentIndex++
			currentValue = value
			currentCount = 1
		}
	}
	if currentIndex == 0{
		currentIndex = len(nums)
	}
	return currentIndex
}

func main(){

}
