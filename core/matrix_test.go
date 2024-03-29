package core

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EqualMatrix(t *testing.T, expected Matrix, actual Matrix) {
	if expected.Equals(actual) {
		assert.True(t, true)
	} else {
		assert.Equal(t, expected, actual)
	}
}

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

	assert.True(t, matrix1.Equals(matrix2))
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

	assert.False(t, matrix1.Equals(matrix2))
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

	assert.Equal(t, result, matrix1.Multiply(matrix2))
}

func TestMultiplyMatrixByTuple(t *testing.T) {
	matrix := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1}}
	tuple := Tuple{1, 2, 3, 1}
	result := Tuple{18, 24, 33, 1}

	assert.Equal(t, result, matrix.MultiplyTuple(tuple))
}

func TestIdentityMatrixWithMatrix(t *testing.T) {
	matrix := Matrix{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32}}
	identity := NewIdentityMatrix()

	assert.Equal(t, matrix, matrix.Multiply(identity))
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

	assert.Equal(t, result, matrix.Transpose())
}

func TestTransposeMatrix(t *testing.T) {
	tuple := Tuple{1, 2, 3, 4}
	identity := NewIdentityMatrix()

	assert.Equal(t, tuple, identity.MultiplyTuple(tuple))
}

func TestTransposeIdentityMatrix(t *testing.T) {
	identity := NewIdentityMatrix()

	assert.Equal(t, identity, identity.Transpose())
}

func TestDeterminant2x2Matrix(t *testing.T) {
	matrix := Matrix{
		{1, 5},
		{-3, 2}}

	assert.Equal(t, 17.0, matrix.Determinant())
}

func Test3x3Submatrix(t *testing.T) {
	matrix := Matrix{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3}}

	Submatrix := Matrix{
		{-3, 2},
		{0, 6}}

	assert.Equal(t, Submatrix, matrix.Submatrix(0, 2))
}

func Test4x4Submatrix(t *testing.T) {
	matrix := Matrix{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1}}

	Submatrix := Matrix{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1}}

	assert.Equal(t, Submatrix, matrix.Submatrix(2, 1))
}

func TestMinor3x3Matrix(t *testing.T) {
	matrix := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5}}

	assert.Equal(t, 25.0, matrix.Minor(1, 0))
}

func TestCofactor3x3Matrix(t *testing.T) {
	matrix := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5}}

	assert.Equal(t, -12.0, matrix.Cofactor(0, 0))
	assert.Equal(t, -25.0, matrix.Cofactor(1, 0))
}

func TestDeterminant3x3Matrix(t *testing.T) {
	matrix := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4}}

	assert.Equal(t, -196.0, matrix.Determinant())
}

func TestDeterminant4x4Matrix(t *testing.T) {
	matrix := Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9}}

	assert.Equal(t, -4071.0, matrix.Determinant())
}

func TestInvertibleMatrix(t *testing.T) {
	matrix := Matrix{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6}}

	assert.True(t, matrix.Invertible())
}

func TestNonInvertibleMatrix(t *testing.T) {
	matrix := Matrix{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0}}

	assert.False(t, matrix.Invertible())
}

func TestInverseMatrix(t *testing.T) {
	matrix := Matrix{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4}}

	Inverse := matrix.Inverse()

	expectedInverse := Matrix{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639}}

	assert.Equal(t, 532.0, matrix.Determinant())
	assert.Equal(t, -160.0, matrix.Cofactor(2, 3))
	assert.Equal(t, -160.0/532, Inverse[3][2])
	assert.Equal(t, 105.0/532, Inverse[2][3])
	assert.True(t, expectedInverse.Equals(Inverse))
}

func TestInverseSecondMatrix(t *testing.T) {
	matrix := Matrix{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4}}

	expectedInverse := Matrix{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308}}

	assert.True(t, expectedInverse.Equals(matrix.Inverse()))
}

func TestInverseThirdMatrix(t *testing.T) {
	matrix := Matrix{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2}}

	expectedInverse := Matrix{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333}}

	assert.True(t, expectedInverse.Equals(matrix.Inverse()))
}

func TestMultiplyProductByInverse(t *testing.T) {
	matrix1 := Matrix{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1}}

	matrix2 := Matrix{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5}}

	result := matrix1.Multiply(matrix2)

	assert.True(t, matrix1.Equals(result.Multiply(matrix2.Inverse())))
}

func TestMultiplyByTranslationMatrix(t *testing.T) {
	matrix := NewIdentityMatrix().Translate(5, -3, 2)
	point := NewPoint(-3, 4, 5)

	result := NewPoint(2, 1, 7)

	assert.Equal(t, result, matrix.MultiplyTuple(point))
}

func TestMultiplyInverseByTranslationMatrix(t *testing.T) {
	matrix := NewIdentityMatrix().Translate(5, -3, 2).Inverse()
	point := NewPoint(-3, 4, 5)

	result := NewPoint(-8, 7, 3)

	assert.Equal(t, result, matrix.MultiplyTuple(point))
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	matrix := NewIdentityMatrix().Translate(5, -3, 2)
	vector := NewVector(-3, 4, 5)

	assert.Equal(t, vector, matrix.MultiplyTuple(vector))
}

func TestMultiplyByScalingMatrix(t *testing.T) {
	matrix := NewIdentityMatrix().Scale(2, 3, 4)
	point := NewPoint(-4, 6, 8)

	result := NewPoint(-8, 18, 32)

	assert.Equal(t, result, matrix.MultiplyTuple(point))
}

func TestMultiplyInverseByScalingMatrix(t *testing.T) {
	matrix := NewIdentityMatrix().Scale(2, 3, 4).Inverse()
	point := NewPoint(-4, 6, 8)

	result := NewPoint(-2, 2, 2)

	assert.Equal(t, result, matrix.MultiplyTuple(point))
}

func TestScalingDoesNotAffectVectors(t *testing.T) {
	matrix := NewIdentityMatrix().Scale(2, 3, 4)
	vector := NewVector(-4, 6, 8)

	result := NewVector(-8, 18, 32)

	assert.Equal(t, result, matrix.MultiplyTuple(vector))
}

func TestRotatePointXAxis(t *testing.T) {
	point := NewPoint(0, 1, 0)
	halfQuarter := NewIdentityMatrix().RotateX(math.Pi / 4)
	fullQuarter := NewIdentityMatrix().RotateX(math.Pi / 2)

	result1 := NewPoint(0, math.Sqrt2/2, math.Sqrt2/2)
	result2 := NewPoint(0, 0, 1)

	EqualTuple(t, result1, halfQuarter.MultiplyTuple(point))
	EqualTuple(t, result2, fullQuarter.MultiplyTuple(point))
}

func TestInverseRotatePointXAxis(t *testing.T) {
	point := NewPoint(0, 1, 0)
	halfQuarter := NewIdentityMatrix().RotateX(math.Pi / 4)
	inv := halfQuarter.Inverse()

	result := NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2)

	EqualTuple(t, result, inv.MultiplyTuple(point))
}

func TestRotatePointYAxis(t *testing.T) {
	point := NewPoint(0, 0, 1)
	halfQuarter := NewIdentityMatrix().RotateY(math.Pi / 4)
	fullQuarter := NewIdentityMatrix().RotateY(math.Pi / 2)

	result1 := NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2)
	result2 := NewPoint(1, 0, 0)

	EqualTuple(t, result1, halfQuarter.MultiplyTuple(point))
	EqualTuple(t, result2, fullQuarter.MultiplyTuple(point))
}

func TestRotatePointZAxis(t *testing.T) {
	point := NewPoint(0, 1, 0)
	halfQuarter := NewIdentityMatrix().RotateZ(math.Pi / 4)
	fullQuarter := NewIdentityMatrix().RotateZ(math.Pi / 2)

	result1 := NewPoint(-math.Sqrt2/2, math.Sqrt2/2, 0)
	result2 := NewPoint(-1, 0, 0)

	EqualTuple(t, result1, halfQuarter.MultiplyTuple(point))
	EqualTuple(t, result2, fullQuarter.MultiplyTuple(point))
}

func TestShearingXtoZ(t *testing.T) {
	transform := NewIdentityMatrix().Shear(0, 1, 0, 0, 0, 0)
	point := NewPoint(2, 3, 4)

	result := NewPoint(6, 3, 4)

	EqualTuple(t, result, transform.MultiplyTuple(point))
}

func TestShearingYtoX(t *testing.T) {
	transform := NewIdentityMatrix().Shear(0, 0, 1, 0, 0, 0)
	point := NewPoint(2, 3, 4)

	result := NewPoint(2, 5, 4)

	EqualTuple(t, result, transform.MultiplyTuple(point))
}

func TestShearingYtoZ(t *testing.T) {
	transform := NewIdentityMatrix().Shear(0, 0, 0, 1, 0, 0)
	point := NewPoint(2, 3, 4)

	result := NewPoint(2, 7, 4)

	EqualTuple(t, result, transform.MultiplyTuple(point))
}

func TestShearingZtoX(t *testing.T) {
	transform := NewIdentityMatrix().Shear(0, 0, 0, 0, 1, 0)
	point := NewPoint(2, 3, 4)

	result := NewPoint(2, 3, 6)

	EqualTuple(t, result, transform.MultiplyTuple(point))
}

func TestShearingZtoY(t *testing.T) {
	transform := NewIdentityMatrix().Shear(0, 0, 0, 0, 0, 1)
	point := NewPoint(2, 3, 4)

	result := NewPoint(2, 3, 7)

	EqualTuple(t, result, transform.MultiplyTuple(point))
}

func TestIndividualTransformationsInSequence(t *testing.T) {
	point := NewPoint(1, 0, 1)
	a := NewIdentityMatrix().RotateX(math.Pi / 2)
	b := NewIdentityMatrix().Scale(5, 5, 5)
	c := NewIdentityMatrix().Translate(10, 5, 7)

	// rotation
	point2 := a.MultiplyTuple(point)
	EqualTuple(t, NewPoint(1, -1, 0), point2)

	// scaling
	point3 := b.MultiplyTuple(point2)
	EqualTuple(t, NewPoint(5, -5, 0), point3)

	// translation
	point4 := c.MultiplyTuple(point3)
	EqualTuple(t, NewPoint(15, 0, 7), point4)
}

func TestChainedTransformationsInReverseOrder(t *testing.T) {
	point := NewPoint(1, 0, 1)

	result := NewIdentityMatrix().
		Translate(10, 5, 7).
		Scale(5, 5, 5).
		RotateX(math.Pi / 2)

	EqualTuple(t, NewPoint(15, 0, 7), result.MultiplyTuple(point))
}

func TestViewTransformDefaultOrientation(t *testing.T) {
	from := NewPoint(0, 0, 0)
	to := NewPoint(0, 0, -1)
	up := NewVector(0, 1, 0)

	transform := ViewTransform(from, to, up)

	assert.Equal(t, NewIdentityMatrix(), transform)
}

func TestViewTransformLookingPositiveZDirection(t *testing.T) {
	from := NewPoint(0, 0, 0)
	to := NewPoint(0, 0, 1)
	up := NewVector(0, 1, 0)

	transform := ViewTransform(from, to, up)

	assert.Equal(t, NewIdentityMatrix().Scale(-1, 1, -1), transform)
}

func TestViewTransformMovesWorld(t *testing.T) {
	from := NewPoint(0, 0, 8)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)

	transform := ViewTransform(from, to, up)

	assert.Equal(t, NewIdentityMatrix().Translate(0, 0, -8), transform)
}

func TestArbitraryViewTransform(t *testing.T) {
	from := NewPoint(1, 3, 2)
	to := NewPoint(4, -2, 8)
	up := NewVector(1, 1, 0)

	transform := ViewTransform(from, to, up)

	result := NewMatrix(4, 4)
	result[0] = []float64{-0.50709, 0.50709, 0.67612, -2.36643}
	result[1] = []float64{0.76772, 0.60609, 0.12122, -2.82843}
	result[2] = []float64{-0.35857, 0.59761, -0.71714, 0.00000}
	result[3] = []float64{0.00000, 0.00000, 0.00000, 1.00000}

	EqualMatrix(t, result, transform)
}
