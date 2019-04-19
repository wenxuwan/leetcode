package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	lenth := m + n - 1
	m = m - 1
	n = n - 1
	for m >= 0 && n >= 0{
		if nums1[m] > nums2[n]{
			nums1[lenth] = nums1[m]
			lenth--
			m--
		}else{
			nums1[lenth] = nums2[n]
			lenth--
			n--
		}
	}
	for n >= 0{
		nums1[lenth] = nums2[n]
		n--
		lenth--
	}
}

func main(){
	nums1 := []int{1,2,4,5,6,0}
	nums2 := []int{3}
	merge(nums1,5,nums2,1)

}
