/*思路就是找到最高的点，然后分别从两侧遍历。
比如从左侧，如果后面节点比自己矮，那么就可以
存水。同理可得右侧。
*/
package main

import "fmt"

func trap(height []int) int {
	maxIndex := 0
	MaxValue := 0
	result := 0
	for index, value := range height{
		if value > MaxValue{
			MaxValue = value
			maxIndex = index
		}
	}
	left := 0
	lo := left + 1
	for left < maxIndex && lo < maxIndex{
		if height[left] > height[lo]{
			result = result + height[left] - height[lo]
			lo++
		}else {
			left = lo
			lo = lo + 1
		}
	}
	right := len(height) - 1
	lor := right - 1
	for right > maxIndex && lor > maxIndex{
		if height[right] > height[lor]{
			result = result + height[right] - height[lor]
			lor--
		}else {
			right = lor
			lor--
		}
	}
	return result
}

func main(){
	test := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(test))
}
