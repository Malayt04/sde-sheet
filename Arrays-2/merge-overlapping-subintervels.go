package arrays2

func mergeOverlappingSubintervels(intervals [][]int) [][]int{
	ans:=make([][] int, 0, len(intervals))
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		if current[1] > next[0]{
			current[1] = max(current[1], next[1])
		}else{
			ans = append(ans, current)
			current = next
		}
	}

	return ans
}