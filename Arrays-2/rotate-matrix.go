package arrays2

func rotateMatrix(nums[][] int){

	top := 0
	bottom := len(nums) - 1

	for top < bottom{
		for i := 0; i < len(nums[top]); i++{
			nums[top][i], nums[bottom][i] = nums[bottom][i], nums[top][i]
		}

		top++
		bottom--
	}

	for i := 0; i < len(nums); i++{
		for j := 0; j < i; j++{
			nums[i][j], nums[j][i] = nums[j][i], nums[i][j]
		}
	}

}