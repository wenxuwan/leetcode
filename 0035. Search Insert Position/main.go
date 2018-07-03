package  main

import "fmt"

func searchInsert(nums []int, target int) int {
	medium := 0
	begin := 0
	end := len(nums) - 1
	for begin <= end{
		medium = (end + begin)/2
		if nums[medium] < target{
			begin = medium + 1
		}else{
			if nums[medium] == target{
				return medium
			}else{
				end  = medium - 1
			}
		}
	}

	return begin
}

func main(){
	array := []int{1,3,5,6}
	fmt.Print(searchInsert(array,0))
}