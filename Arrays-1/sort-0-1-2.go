package arrays1

import "sort"

func sortColors(nums []int) []int{
	sort.Ints(nums)
	return nums
}