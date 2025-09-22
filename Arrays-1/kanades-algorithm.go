package arrays1

func maxSubarray(nums[]int) int {
	local := 0
	global := 0

	for i := 0; i < len(nums); i++{
		local += nums[i]

		global =max(global, local)

		if local < 0{
			local = 0
		}
	}

	return global
}