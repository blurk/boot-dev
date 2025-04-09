package main

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)

	for row := range rows {
		matrix[row] = make([]int, cols)

		for col := range cols {
			matrix[row][col] = row * col
		}
	}

	return matrix
}
