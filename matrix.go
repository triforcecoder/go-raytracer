package main

import "math"

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

func (matrix Matrix) inverse() Matrix {
	if !matrix.invertible() {
		panic("precondition - matrix is not invertible")
	}

	result := createMatrix(len(matrix), len(matrix[0]))
	determinant := matrix.determinant()

	for row := range matrix {
		for col := range matrix[0] {
			cofactor := matrix.cofactor(row, col)
			result[col][row] = cofactor / determinant
		}
	}

	return result
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

func createTranslationMatrix(x float64, y float64, z float64) Matrix {
	matrix := createIdentityMatrix()

	matrix[0][3] = x
	matrix[1][3] = y
	matrix[2][3] = z

	return matrix
}

func createScalingMatrix(x float64, y float64, z float64) Matrix {
	matrix := createMatrix(4, 4)

	matrix[0][0] = x
	matrix[1][1] = y
	matrix[2][2] = z
	matrix[3][3] = 1

	return matrix
}

func rotationX(radians float64) Matrix {
	matrix := createIdentityMatrix()

	matrix[1][1] = math.Cos(radians)
	matrix[1][2] = -math.Sin(radians)
	matrix[2][1] = math.Sin(radians)
	matrix[2][2] = math.Cos(radians)

	return matrix
}

func rotationY(radians float64) Matrix {
	matrix := createIdentityMatrix()

	matrix[0][0] = math.Cos(radians)
	matrix[0][2] = math.Sin(radians)
	matrix[2][0] = -math.Sin(radians)
	matrix[2][2] = math.Cos(radians)

	return matrix
}

func rotationZ(radians float64) Matrix {
	matrix := createIdentityMatrix()

	matrix[0][0] = math.Cos(radians)
	matrix[0][1] = -math.Sin(radians)
	matrix[1][0] = math.Sin(radians)
	matrix[1][1] = math.Cos(radians)

	return matrix
}

func shearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	matrix := createIdentityMatrix()

	matrix[0][1] = xy
	matrix[0][2] = xz
	matrix[1][0] = yx
	matrix[1][2] = yz
	matrix[2][0] = zx
	matrix[2][1] = zy

	return matrix
}
