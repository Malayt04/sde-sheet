package arrays2

import "sort"

func repeatNumbers(nums []int) int {
	sort.Ints(nums)
	var missingNumber int

	for i := 0; i < len(nums) - 1; i++ {
		if nums[i + 1] != i+1 {
			missingNumber = i + 1
		}
	}
	return missingNumber
}