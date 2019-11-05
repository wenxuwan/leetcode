package main

func twoSum(numbers []int, target int) []int {
	begin := 0
	end := len(numbers) - 1
	result := make([]int,2)
	for begin <= end{
		if numbers[begin] + numbers[end] > target{
			end--
		}
		if numbers[begin] + numbers[end] < target{
			begin++
		}
		if numbers[begin] + numbers[end] == target{
			result[0] = begin + 1
			result[1] = end + 1
			break
		}
	}
	return result
}

func main(){

}
