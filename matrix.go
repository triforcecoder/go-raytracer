package main

type Matrix [][]float64

func (matrix Matrix) equals(other Matrix) bool {
	if len(matrix) != len(other) {
		return false
	}

	for row := range matrix {
		if len(matrix[row]) != len(other[row]) {
			return false
		}
	}

	for row := range matrix {
		for col := range matrix[row] {
			if !equal(matrix[row][col], other[row][col]) {
				return false
			}
		}
	}

	return true
}

func (matrix Matrix) multiply(other Matrix) Matrix {
	result := createMatrix(len(matrix), len(matrix[0]))

	for row := range matrix {
		for col := range matrix[row] {
			elem := 0.0

			for i := range matrix {
				elem += matrix[row][i] * other[i][col]
			}

			result[row][col] = elem
		}
	}

	return result
}

func (matrix Matrix) multiplyTuple(tuple Tuple) Tuple {
	result := [4]float64{}

	for i := 0; i < 4; i++ {
		elem := 0.0

		elem += matrix[i][0] * tuple.x
		elem += matrix[i][1] * tuple.y
		elem += matrix[i][2] * tuple.z
		elem += matrix[i][3] * tuple.w

		result[i] = elem
	}

	return Tuple{result[0], result[1], result[2], result[3]}
}

func createMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	return matrix
}
