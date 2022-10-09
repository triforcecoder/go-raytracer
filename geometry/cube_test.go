package geometry

import (
	. "go-raytracer/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayIntersectsCube(t *testing.T) {
	examples := [][]Tuple{
		{NewPoint(5, 0.5, 0), NewVector(-1, 0, 0)},
		{NewPoint(-5, 0.5, 0), NewVector(1, 0, 0)},
		{NewPoint(0.5, 5, 0), NewVector(0, -1, 0)},
		{NewPoint(0.5, -5, 0), NewVector(0, 1, 0)},
		{NewPoint(0.5, 0, 5), NewVector(0, 0, -1)},
		{NewPoint(0.5, 0, -5), NewVector(0, 0, 1)},
		{NewPoint(0, 0.5, 0), NewVector(0, 0, 1)},
	}

	expected := [][]float64{
		{4, 6},
		{4, 6},
		{4, 6},
		{4, 6},
		{4, 6},
		{4, 6},
		{-1, 1},
	}

	c := NewCube()

	for i := range examples {
		r := NewRay(examples[i][0], examples[i][1])
		xs := c.Intersects(r)

		assert.Equal(t, 2, len(xs))
		assert.Equal(t, expected[i][0], xs[0].T)
		assert.Equal(t, expected[i][1], xs[1].T)
	}
}

func TestRayMissesCube(t *testing.T) {
	examples := [][]Tuple{
		{NewPoint(-2, 0, 0), NewVector(0.2673, 0.5345, 0.8018)},
		{NewPoint(0, -2, 0), NewVector(0.8018, 0.2673, 0.5345)},
		{NewPoint(0, 0, -2), NewVector(0.5345, 0.8018, 0.2673)},
		{NewPoint(2, 0, 2), NewVector(0, 0, -1)},
		{NewPoint(0, 2, 2), NewVector(0, -1, 0)},
		{NewPoint(2, 2, 0), NewVector(-1, 0, 0)},
	}

	c := NewCube()

	for i := range examples {
		r := NewRay(examples[i][0], examples[i][1])
		xs := c.Intersects(r)

		assert.Equal(t, 0, len(xs))
	}
}

func TestNormalOnSurfaceOfCube(t *testing.T) {
	examples := [][]Tuple{
		{NewPoint(1, 0.5, -0.8), NewVector(1, 0, 0)},
		{NewPoint(-1, -0.2, 0.9), NewVector(-1, 0, 0)},
		{NewPoint(-0.4, 1, -0.1), NewVector(0, 1, 0)},
		{NewPoint(0.3, -1, -0.7), NewVector(0, -1, 0)},
		{NewPoint(-0.6, 0.3, 1), NewVector(0, 0, 1)},
		{NewPoint(0.4, 0.4, -1), NewVector(0, 0, -1)},
		{NewPoint(1, 1, 1), NewVector(1, 0, 0)},
		{NewPoint(-1, -1, -1), NewVector(-1, 0, 0)},
	}

	c := NewCube()

	for i := range examples {
		assert.Equal(t, examples[i][1], c.NormalAt(examples[i][0]))
	}
}
