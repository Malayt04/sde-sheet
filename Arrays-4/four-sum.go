package arrays4

import "sort"

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    ans := make([][]int, 0)

    for i := 0; i < n-3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        a := nums[i]

        for j := i + 1; j < n-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            b := nums[j]

            l := j + 1
            r := n - 1

            for l < r {
                sum := nums[l] + nums[r]
                if sum < target-(a+b) {
                    l++
                } else if sum > target-(a+b) {
                    r--
                } else {
                    ans = append(ans, []int{a, b, nums[l], nums[r]})
                    for l < r && nums[l] == nums[l+1] {
                        l++
                    }
                    for l < r && nums[r] == nums[r-1] {
                        r--
                    }
                    l++
                    r--
                }
            }
        }
    }

    return ans
}
