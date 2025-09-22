package arrays3

func binarySearch(nums[] int, target int) int{

	l, r := 0, len(nums) - 1

	for l <= r{
		mid := l + (r - l) / 2

		if nums[mid] == target {
			return mid
		}else if nums[mid] < target {
			l = mid + 1
		}else{
			r = mid - 1
		}
	}

	return l

}

func searchMatrix(matrix[][] int, target int) int{

	ans := 0

	for i := 0; i < len(matrix); i++{
		ans = binarySearch(matrix[i], target)
	}
	return ans
}