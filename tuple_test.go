package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleIsPoint(t *testing.T) {
	point := NewPoint(4.3, -4.2, 3.1)
	assert.True(t, point.IsPoint())
	assert.False(t, point.IsVector())
}

func TestTupleIsVector(t *testing.T) {
	vector := NewVector(4.3, -4.2, 3.1)
	assert.True(t, vector.IsVector())
	assert.False(t, vector.IsPoint())
}

func TestFloatEqual(t *testing.T) {
	assert.False(t, floatEquals(1, 1.01))
	assert.True(t, floatEquals(1, 1.000001))
}

func TestTupleEquals(t *testing.T) {
	point1 := NewPoint(8, 8, 9)
	point2 := NewPoint(50, 100, 64)
	vector1 := NewVector(8, 8, 9)
	vector2 := NewVector(8, 8, 9)

	assert.False(t, point1.Equals(point2))
	assert.False(t, point1.Equals(vector1))
	assert.True(t, vector1.Equals(vector2))
}

func TestAddTuple(t *testing.T) {
	tuple1 := Tuple{3, -2, 5, 1}
	tuple2 := Tuple{-2, 3, 1, 0}
	result := Tuple{1, 1, 6, 1}

	assert.Equal(t, result, tuple1.Add(tuple2))
}

func TestSubtractPoint(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)
	result := NewVector(-2, -4, -6)

	assert.Equal(t, result, point1.Subtract(point2))
}

func TestSubtractVector(t *testing.T) {
	point := NewPoint(3, 2, 1)
	vector := NewVector(5, 6, 7)
	result := NewPoint(-2, -4, -6)

	assert.Equal(t, result, point.Subtract(vector))
}

func TestNegateTuple(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}

	assert.Equal(t, Tuple{-1, 2, -3, 4}, tuple.Negate())
}

func TestMultiplyTuple(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}

	assert.Equal(t, Tuple{3.5, -7, 10.5, -14}, tuple.Multiply(3.5))
	assert.Equal(t, Tuple{0.5, -1, 1.5, -2}, tuple.Multiply(0.5))
}

func TestDivideTuple(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}

	assert.Equal(t, Tuple{0.5, -1, 1.5, -2}, tuple.Divide(2))
}

func TestVectorMagnitude(t *testing.T) {
	assert.Equal(t, 1.0, NewVector(1, 0, 0).Magnitude())
	assert.Equal(t, 1.0, NewVector(0, 1, 0).Magnitude())
	assert.Equal(t, 1.0, NewVector(0, 0, 1).Magnitude())
	assert.Equal(t, math.Sqrt(14), NewVector(1, 2, 3).Magnitude())
	assert.Equal(t, math.Sqrt(14), NewVector(-1, -2, -3).Magnitude())
}

func TestNormalizeVector(t *testing.T) {
	assert.Equal(t, NewVector(1, 0, 0), NewVector(4, 0, 0).Normalize())
	assert.True(t, NewVector(0.26726, 0.53452, 0.80178).Equals(NewVector(1, 2, 3).Normalize()))
	assert.Equal(t, 1.0, NewVector(1, 2, 3).Normalize().Magnitude())
}

func TestDotProductVector(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)
	assert.Equal(t, 20.0, vector1.Dot(vector2))
}

func TestCrossProductVector(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)
	assert.Equal(t, NewVector(-1, 2, -1), vector1.Cross(vector2))
	assert.Equal(t, NewVector(1, -2, 1), vector2.Cross(vector1))
}
