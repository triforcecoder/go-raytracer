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

func (matrix Matrix) transpose() Matrix {
	result := createMatrix(len(matrix), len(matrix[0]))

	for row := range result {
		for col := range result[0] {
			result[row][col] = matrix[col][row]
		}
	}

	return result
}

func (matrix Matrix) determinant() float64 {
	if len(matrix) == 2 && len(matrix[0]) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	determinant := 0.0

	for col := range matrix[0] {
		determinant += matrix[0][col] * matrix.cofactor(0, col)
	}

	return determinant
}

func (matrix Matrix) submatrix(skipRow, skipCol int) Matrix {
	result := createMatrix(len(matrix)-1, len(matrix[0])-1)

	var row, col int
	for i := range result {
		if row == skipRow {
			row++
		}

		col = 0
		for j := range result[0] {
			if col == skipCol {
				col++
			}

			result[i][j] = matrix[row][col]
			col++
		}

		row++
	}

	return result
}

func (matrix Matrix) minor(skipRow, skipCol int) float64 {
	return matrix.submatrix(skipRow, skipCol).determinant()
}

func (matrix Matrix) cofactor(skipRow, skipCol int) float64 {
	minor := matrix.submatrix(skipRow, skipCol).determinant()

	if (skipRow+skipCol)%2 == 0 {
		return minor
	} else {
		return minor * -1
	}
}

func (matrix Matrix) invertible() bool {
	return matrix.determinant() != 0
}

func createMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	return matrix
}

func createIdentityMatrix() Matrix {
	matrix := createMatrix(4, 4)

	matrix[0][0] = 1
	matrix[1][1] = 1
	matrix[2][2] = 1
	matrix[3][3] = 1

	return matrix
}
