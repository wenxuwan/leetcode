package main

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]int)
	result := 0
	for _, value := range nums {
		hashMap[value] = 1
	}
	for _, value := range nums {
		tmpResult := 1
		if tmpResult > result {
			result = tmpResult
		}
		if _, ok := hashMap[value-1]; ok {
			continue
		}
		value++
		for {
			if _, ok := hashMap[value]; ok {
				value++
				tmpResult++
			} else {
				break
			}
		}
		if tmpResult > result {
			result = tmpResult
		}
	}
	return result
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	longestConsecutive(arr)
}
