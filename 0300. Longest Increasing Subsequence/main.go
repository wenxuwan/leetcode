package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	bs := make([]int, len(nums), len(nums))
	max := 0
	for index, value := range nums{
		for i := 0; i < index;i++{
			if value > nums[i]{
				if bs[index] < bs[i] + 1{
					bs[index] = bs[i] + 1
				}
			}
		}
	}
	for _, value := range bs{
		if value > max{
			max = value
		}
	}
	return max + 1
}

func main(){
	data := []int{-2,-1}
	lengthOfLIS(data)
}
