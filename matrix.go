package main

import "math"

type Matrix [][]float64

func NewMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	return matrix
}

func NewIdentityMatrix() Matrix {
	matrix := NewMatrix(4, 4)

	matrix[0][0] = 1
	matrix[1][1] = 1
	matrix[2][2] = 1
	matrix[3][3] = 1

	return matrix
}

func ViewTransform(from Tuple, to Tuple, up Tuple) Matrix {
	forward := to.Subtract(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)
	orientation := NewMatrix(4, 4)
	orientation[0] = []float64{left.x, left.y, left.z, 0}
	orientation[1] = []float64{trueUp.x, trueUp.y, trueUp.z, 0}
	orientation[2] = []float64{-forward.x, -forward.y, -forward.z, 0}
	orientation[3] = []float64{0, 0, 0, 1}

	return orientation.Translate(-from.x, -from.y, -from.z)
}

func (matrix Matrix) Equals(other Matrix) bool {
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
			if !floatEquals(matrix[row][col], other[row][col]) {
				return false
			}
		}
	}

	return true
}

func (matrix Matrix) Multiply(other Matrix) Matrix {
	result := NewMatrix(len(matrix), len(matrix[0]))

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

func (matrix Matrix) MultiplyTuple(tuple Tuple) Tuple {
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

func (matrix Matrix) Transpose() Matrix {
	result := NewMatrix(len(matrix), len(matrix[0]))

	for row := range result {
		for col := range result[0] {
			result[row][col] = matrix[col][row]
		}
	}

	return result
}

func (matrix Matrix) Determinant() float64 {
	if len(matrix) == 2 && len(matrix[0]) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	Determinant := 0.0

	for col := range matrix[0] {
		Determinant += matrix[0][col] * matrix.Cofactor(0, col)
	}

	return Determinant
}

func (matrix Matrix) Submatrix(skipRow, skipCol int) Matrix {
	result := NewMatrix(len(matrix)-1, len(matrix[0])-1)

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

func (matrix Matrix) Minor(skipRow, skipCol int) float64 {
	return matrix.Submatrix(skipRow, skipCol).Determinant()
}

func (matrix Matrix) Cofactor(skipRow, skipCol int) float64 {
	Minor := matrix.Submatrix(skipRow, skipCol).Determinant()

	if (skipRow+skipCol)%2 == 0 {
		return Minor
	} else {
		return Minor * -1
	}
}

func (matrix Matrix) Invertible() bool {
	return matrix.Determinant() != 0
}

func (matrix Matrix) Inverse() Matrix {
	if !matrix.Invertible() {
		panic("precondition - matrix is not Invertible")
	}

	result := NewMatrix(len(matrix), len(matrix[0]))
	Determinant := matrix.Determinant()

	for row := range matrix {
		for col := range matrix[0] {
			Cofactor := matrix.Cofactor(row, col)
			result[col][row] = Cofactor / Determinant
		}
	}

	return result
}

func (matrix Matrix) Translate(x float64, y float64, z float64) Matrix {
	translation := NewIdentityMatrix()

	translation[0][3] = x
	translation[1][3] = y
	translation[2][3] = z

	return matrix.Multiply(translation)
}

func (matrix Matrix) Scale(x float64, y float64, z float64) Matrix {
	scaling := NewMatrix(4, 4)

	scaling[0][0] = x
	scaling[1][1] = y
	scaling[2][2] = z
	scaling[3][3] = 1

	return matrix.Multiply(scaling)
}

func (matrix Matrix) RotateX(radians float64) Matrix {
	rotation := NewIdentityMatrix()

	rotation[1][1] = math.Cos(radians)
	rotation[1][2] = -math.Sin(radians)
	rotation[2][1] = math.Sin(radians)
	rotation[2][2] = math.Cos(radians)

	return matrix.Multiply(rotation)
}

func (matrix Matrix) RotateY(radians float64) Matrix {
	rotation := NewIdentityMatrix()

	rotation[0][0] = math.Cos(radians)
	rotation[0][2] = math.Sin(radians)
	rotation[2][0] = -math.Sin(radians)
	rotation[2][2] = math.Cos(radians)

	return matrix.Multiply(rotation)
}

func (matrix Matrix) RotateZ(radians float64) Matrix {
	rotation := NewIdentityMatrix()

	rotation[0][0] = math.Cos(radians)
	rotation[0][1] = -math.Sin(radians)
	rotation[1][0] = math.Sin(radians)
	rotation[1][1] = math.Cos(radians)

	return matrix.Multiply(rotation)
}

func (matrix Matrix) Shear(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) Matrix {
	shearing := NewIdentityMatrix()

	shearing[0][1] = xy
	shearing[0][2] = xz
	shearing[1][0] = yx
	shearing[1][2] = yz
	shearing[2][0] = zx
	shearing[2][1] = zy

	return matrix.Multiply(shearing)
}
