package arrays3

func majorityElementTwo(nums []int) []int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
    }

    ans := make([]int, 0)
    threshold := len(nums) / 3

    for num, count := range m {
        if count > threshold {
            ans = append(ans, num)
        }
    }

    return ans
}
