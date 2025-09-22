package arrays1

func generate(numRows int) [][]int {
    ans := [][]int{{1}}

    for i := 1; i < numRows; i++ {
        prev := ans[i-1]
        arr := []int{1} 

        for j := 0; j < len(prev)-1; j++ {
            arr = append(arr, prev[j]+prev[j+1])
        }

        arr = append(arr, 1) 
        ans = append(ans, arr)
    }

    return ans
}
