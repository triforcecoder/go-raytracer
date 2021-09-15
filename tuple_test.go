package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleIsPoint(t *testing.T) {
	point := createPoint(4.3, -4.2, 3.1)
	assert.True(t, point.isPoint())
	assert.False(t, point.isVector())
}

func TestTupleIsVector(t *testing.T) {
	vector := createVector(4.3, -4.2, 3.1)
	assert.True(t, vector.isVector())
	assert.False(t, vector.isPoint())
}

func TestFloatEqual(t *testing.T) {
	assert.False(t, equal(1, 1.01))
	assert.True(t, equal(1, 1.000001))
}

func TestTupleEquals(t *testing.T) {
	point1 := createPoint(8, 8, 9)
	point2 := createPoint(50, 100, 64)
	vector1 := createVector(8, 8, 9)
	vector2 := createVector(8, 8, 9)

	assert.False(t, point1.equals(point2))
	assert.False(t, point1.equals(vector1))
	assert.True(t, vector1.equals(vector2))
}

func TestAddTuple(t *testing.T) {
	tuple1 := Tuple{3, -2, 5, 1}
	tuple2 := Tuple{-2, 3, 1, 0}
	result := Tuple{1, 1, 6, 1}

	assert.Equal(t, result, tuple1.add(tuple2))
}

func TestSubtractPoint(t *testing.T) {
	point1 := createPoint(3, 2, 1)
	point2 := createPoint(5, 6, 7)
	result := createVector(-2, -4, -6)

	assert.Equal(t, result, point1.subtract(point2))
}

func TestSubtractVector(t *testing.T) {
	point := createPoint(3, 2, 1)
	vector := createVector(5, 6, 7)
	result := createPoint(-2, -4, -6)

	assert.Equal(t, result, point.subtract(vector))
}

func TestNegateTuple(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}

	assert.Equal(t, Tuple{-1, 2, -3, 4}, tuple.negate())
}

func TestMultiplyTuple(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}

	assert.Equal(t, Tuple{3.5, -7, 10.5, -14}, tuple.multiply(3.5))
	assert.Equal(t, Tuple{0.5, -1, 1.5, -2}, tuple.multiply(0.5))
}

func TestDivideTuple(t *testing.T) {
	tuple := Tuple{1, -2, 3, -4}

	assert.Equal(t, Tuple{0.5, -1, 1.5, -2}, tuple.divide(2))
}

func TestVectorMagnitude(t *testing.T) {
	assert.Equal(t, 1.0, createVector(1, 0, 0).magnitude())
	assert.Equal(t, 1.0, createVector(0, 1, 0).magnitude())
	assert.Equal(t, 1.0, createVector(0, 0, 1).magnitude())
	assert.Equal(t, math.Sqrt(14), createVector(1, 2, 3).magnitude())
	assert.Equal(t, math.Sqrt(14), createVector(-1, -2, -3).magnitude())
}

func TestNormalizeVector(t *testing.T) {
	assert.Equal(t, createVector(1, 0, 0), createVector(4, 0, 0).normalize())
	assert.True(t, createVector(0.26726, 0.53452, 0.80178).equals(createVector(1, 2, 3).normalize()))
	assert.Equal(t, 1.0, createVector(1, 2, 3).normalize().magnitude())
}

func TestDotProductVector(t *testing.T) {
	vector1 := createVector(1, 2, 3)
	vector2 := createVector(2, 3, 4)
	assert.Equal(t, 20.0, vector1.dot(vector2))
}

func TestCrossProductVector(t *testing.T) {
	vector1 := createVector(1, 2, 3)
	vector2 := createVector(2, 3, 4)
	assert.Equal(t, createVector(-1, 2, -1), vector1.cross(vector2))
	assert.Equal(t, createVector(1, -2, 1), vector2.cross(vector1))
}
