package arrays1

import "slices"

func nextPermutation(nums []int) {
    n := len(nums)
    if n < 2 {
        return
    }

    i := n - 1
    for i > 0 && nums[i-1] >= nums[i] {
        i--
    }

    if i == 0 {
        slices.Reverse(nums)
        return
    }

    j := n - 1
    for nums[j] <= nums[i-1] {
        j--
    }

    nums[i-1], nums[j] = nums[j], nums[i-1]

    slices.Reverse(nums[i:])
}
