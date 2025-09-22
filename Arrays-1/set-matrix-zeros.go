package arrays1

type point struct{
    x, y int
}

func makeMatrixZero(matrix [][]int, x, y int) {

    for i := 0; i < len(matrix[0]); i++ {
        matrix[x][i] = 0
    }
    for i := 0; i < len(matrix); i++ {
        matrix[i][y] = 0
    }
}

func setZeroes(matrix [][]int) {
    var zeros []point

    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                zeros = append(zeros, point{i, j})
            }
        }
    }

    for _, p := range zeros {
        makeMatrixZero(matrix, p.x, p.y)
    }
}