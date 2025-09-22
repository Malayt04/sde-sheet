package arrays1

func stockBuyAndSell(nums[] int) int{
	profit := 0
	buyPrice := nums[0]

	for i := 1; i < len(nums); i++{
		if(buyPrice > nums[i]){
			buyPrice = nums[i]
		}

		profit = max(profit, nums[i] - buyPrice) 
	}

	return profit
}