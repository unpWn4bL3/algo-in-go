package main

func main() {
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for k := range matrix {
		if matrix[k][0] == 0 {
			row0 = true
		}
	}
	for k := range matrix[0] {
		if matrix[0][k] == 0 {
			col0 = true
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
			if matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for k := range matrix {
			matrix[k][0] = 0
		}
	}
	if col0 {
		for k := range matrix[0] {
			matrix[0][k] = 0
		}
	}
}

// func setZeroes(matrix [][]int) {
// 	m, n := len(matrix), len(matrix[0])
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if matrix[i][j] == 0 {
// 				// 向上
// 				for x, y := i, j; y >= 0; y-- {
// 					matrix[x][y] = 0
// 				}
// 				// 向下
// 				for x, y := i, j; y < n; y++ {
// 					matrix[x][y] = 0
// 				}
// 				// 向左
// 				for x, y := i, j; x >= 0; x-- {
// 					matrix[x][y] = 0
// 				}
// 				//向右
// 				for x, y := i, j; x < m; x++ {
// 					matrix[x][y] = 0
// 				}
// 			}
// 		}
// 	}
// }
