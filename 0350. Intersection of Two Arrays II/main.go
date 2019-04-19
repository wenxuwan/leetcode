package main

func intersect(nums1 []int, nums2 []int) []int {
	var mapValue map[int]int
	mapValue = make(map[int]int)
	result := make([]int,0,len(nums1))
	for i := 0; i < len(nums1);i++{
		mapValue[nums1[i]] = mapValue[nums1[i]] + 1
	}
	for j := 0;j < len(nums2);j++{
		if mapValue[nums2[j]] != 0{
			mapValue[nums2[j]]--
			result = append(result, nums2[j])
		}
	}
	return result
}

func main(){

}