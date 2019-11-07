package main

func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	result := make([]int,0,len(nums1))
	for i := 0; i < len(nums1);i++{
		if _,ok := hashMap[nums1[i]];!ok{
			hashMap[nums1[i]]++
		}
	}
	for j := 0;j < len(nums2);j++{
		if _,ok := hashMap[nums2[j]];ok{
			if hashMap[nums2[j]] == 1{
				result = append(result,nums2[j])
				hashMap[nums2[j]]++
			}
		}
	}
	return result
}

func main(){

}
