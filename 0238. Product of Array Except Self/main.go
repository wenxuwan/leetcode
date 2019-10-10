package main

func productExceptSelf(nums []int) []int {
	lenth := len(nums)
	result := make([]int, len(nums), len(nums))
	left := 1
	right := 1
	for i := 0;i < lenth;i++{
		result[i] = 1
	}
	for i := 0;i < lenth;i++{
		result[i] = result[i] * left
		left = left * nums[i]
		result[lenth - i - 1] = result[lenth - i - 1] * right
		right = right * nums[lenth - i - 1]
	}
	return result
}

func main(){

}
