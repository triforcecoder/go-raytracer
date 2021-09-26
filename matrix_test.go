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

	assert.True(t, matrix1.multiply(matrix2).equals(result))
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
