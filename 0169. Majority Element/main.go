package main


func majorityElement(nums []int) int {
	storeMap := make(map[int]int,len(nums))
	result := 0
	for i := 0 ; i < len(nums);i++{
		storeMap[nums[i]]++
		if storeMap[nums[i]] > len(nums) / 2{
			result = nums[i]
			break
		}
	}
	return result
}

func main(){

}