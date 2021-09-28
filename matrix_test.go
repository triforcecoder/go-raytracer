package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate2x2Matrix(t *testing.T) {
	matrix := Matrix{
		{-3, 5},
		{1, -2}}

	assert.Equal(t, -3.0, matrix[0][0])
	assert.Equal(t, 5.0, matrix[0][1])
	assert.Equal(t, 1.0, matrix[1][0])
	assert.Equal(t, -2.0, matrix[1][1])
}

func TestCreate3x3Matrix(t *testing.T) {
	matrix := Matrix{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1}}

	assert.Equal(t, -3.0, matrix[0][0])
	assert.Equal(t, -2.0, matrix[1][1])
	assert.Equal(t, 1.0, matrix[2][2])
}
func TestCreate4x4Matrix(t *testing.T) {
	matrix := Matrix{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5}}

	assert.Equal(t, 1.0, matrix[0][0])
	assert.Equal(t, 4.0, matrix[0][3])
	assert.Equal(t, 5.5, matrix[1][0])
	assert.Equal(t, 7.5, matrix[1][2])
	assert.Equal(t, 11.0, matrix[2][2])
	assert.Equal(t, 13.5, matrix[3][0])
	assert.Equal(t, 15.5, matrix[3][2])
}

func TestMatrixEqual(t *testing.T) {
	matrix1 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2}}

	matrix2 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2}}

	assert.True(t, matrix1.equals(matrix2))
}

func TestMatrixNotEqual(t *testing.T) {
	matrix1 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2}}

	matrix2 := Matrix{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1}}

	assert.False(t, matrix1.equals(matrix2))
}

func TestMultiplyMatrix(t *testing.T) {
	matrix1 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2}}

	matrix2 := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8}}

	result := Matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42}}

	assert.Equal(t, result, matrix1.multiply(matrix2))
}

func TestMultiplyMatrixByTuple(t *testing.T) {
	matrix := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1}}
	tuple := Tuple{1, 2, 3, 1}
	result := Tuple{18, 24, 33, 1}

	assert.Equal(t, result, matrix.multiplyTuple(tuple))
}

func TestIdentityMatrixWithMatrix(t *testing.T) {
	matrix := Matrix{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32}}
	identity := createIdentityMatrix()

	assert.Equal(t, matrix, matrix.multiply(identity))
}

func TestIdentityMatrixWithTuple(t *testing.T) {
	matrix := Matrix{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8}}
	result := Matrix{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8}}

	assert.Equal(t, result, matrix.transpose())
}

func TestTransposeMatrix(t *testing.T) {
	tuple := Tuple{1, 2, 3, 4}
	identity := createIdentityMatrix()

	assert.Equal(t, tuple, identity.multiplyTuple(tuple))
}

func TestTransposeIdentityMatrix(t *testing.T) {
	identity := createIdentityMatrix()

	assert.Equal(t, identity, identity.transpose())
}

func TestDeterminant2x2Matrix(t *testing.T) {
	matrix := Matrix{
		{1, 5},
		{-3, 2}}

	assert.Equal(t, 17.0, matrix.determinant())
}

func Test3x3Submatrix(t *testing.T) {
	matrix := Matrix{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3}}

	submatrix := Matrix{
		{-3, 2},
		{0, 6}}

	assert.Equal(t, submatrix, matrix.submatrix(0, 2))
}

func Test4x4Submatrix(t *testing.T) {
	matrix := Matrix{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1}}

	submatrix := Matrix{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1}}

	assert.Equal(t, submatrix, matrix.submatrix(2, 1))
}

func TestMinor3x3Matrix(t *testing.T) {
	matrix := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5}}

	assert.Equal(t, 25.0, matrix.minor(1, 0))
}

func TestCofactor3x3Matrix(t *testing.T) {
	matrix := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5}}

	assert.Equal(t, -12.0, matrix.cofactor(0, 0))
	assert.Equal(t, -25.0, matrix.cofactor(1, 0))
}

func TestDeterminant3x3Matrix(t *testing.T) {
	matrix := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4}}

	assert.Equal(t, -196.0, matrix.determinant())
}

func TestDeterminant4x4Matrix(t *testing.T) {
	matrix := Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9}}

	assert.Equal(t, -4071.0, matrix.determinant())
}

func TestInvertibleMatrix(t *testing.T) {
	matrix := Matrix{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6}}

	assert.True(t, matrix.invertible())
}

func TestNonInvertibleMatrix(t *testing.T) {
	matrix := Matrix{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0}}

	assert.False(t, matrix.invertible())
}
