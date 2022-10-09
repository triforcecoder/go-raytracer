package geometry

import (
	. "go-raytracer/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaneNormalIsConstantEverywhere(t *testing.T) {
	plane := NewPlane()
	n1 := plane.NormalAt(NewPoint(0, 0, 0))
	n2 := plane.NormalAt(NewPoint(10, 0, -10))
	n3 := plane.NormalAt(NewPoint(-5, 0, 150))

	assert.Equal(t, NewVector(0, 1, 0), n1)
	assert.Equal(t, NewVector(0, 1, 0), n2)
	assert.Equal(t, NewVector(0, 1, 0), n3)
}

func TestIntersectWithRayParallelToPlane(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1))

	xs := plane.Intersects(ray)

	assert.Equal(t, 0, len(xs))
}

func TestIntersectWithCoplanerRay(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))

	xs := plane.Intersects(ray)

	assert.Equal(t, 0, len(xs))
}

func TestRayIntersectsPlaneFromAbove(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, 1, 0), NewVector(0, -1, 0))

	xs := plane.Intersects(ray)

	assert.Equal(t, 1, len(xs))
	assert.Equal(t, 1.0, xs[0].T)
	assert.Equal(t, plane, xs[0].Object)
}

func TestRayIntersectsPlaneFromBelow(t *testing.T) {
	plane := NewPlane()
	ray := NewRay(NewPoint(0, -1, 0), NewVector(0, 1, 0))

	xs := plane.Intersects(ray)

	assert.Equal(t, 1, len(xs))
	assert.Equal(t, 1.0, xs[0].T)
	assert.Equal(t, plane, xs[0].Object)
}
