package arrays2

import "sort"

func mergeSortedArrays(nums1 []int, nums2 []int) []int{

	for i := 0; i < len(nums2); i++{
		nums1 = append(nums1, nums2[i])
	} 

	sort.Ints(nums1)

	return nums1
}
