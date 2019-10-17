/*
遍历一次，将value作为下标，更改对应的value的符号。如果已经为负数，就不需要再更改了。
再遍历一次，如果value大于0证明没出现过该值。
*/
package main

func findDisappearedNumbers(nums []int) []int {
	result := make([]int,0,0)
	for _, value := range nums{
		if value < 0{
			value = -value
		}
		if nums[value - 1] > 0 {
			nums[value - 1] = -nums[value - 1]
		}
	}
	for index,value := range nums{
		if value > 0{
			result = append(result, index + 1)
		}
	}
	return result
}

func main(){

}
