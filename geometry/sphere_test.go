package geometry

import (
	. "go-raytracer/core"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EqualTuple(t *testing.T, expected Tuple, actual Tuple) {
	if expected.Equals(actual) {
		assert.True(t, true)
	} else {
		assert.Equal(t, expected, actual)
	}
}

func TestIntersectsSphereAtTwoPoints(t *testing.T) {
	origin := NewPoint(0, 0, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := NewSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 4.0, xs[0].T)
	assert.Equal(t, 6.0, xs[1].T)
}

func TestIntersectsSphereAtTangent(t *testing.T) {
	origin := NewPoint(0, -1, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := NewSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 5.0, xs[0].T)
	assert.Equal(t, 5.0, xs[1].T)
}

func TestMissesSphere(t *testing.T) {
	origin := NewPoint(0, 2, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := NewSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayOriginInsideSphere(t *testing.T) {
	origin := NewPoint(0, 0, 0)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := NewSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -1.0, xs[0].T)
	assert.Equal(t, 1.0, xs[1].T)
}

func TestSphereBehindRay(t *testing.T) {
	origin := NewPoint(0, 0, 5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := NewSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -6.0, xs[0].T)
	assert.Equal(t, -4.0, xs[1].T)
}

func TestIntersectSetsObject(t *testing.T) {
	origin := NewPoint(0, 0, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := NewSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, s, xs[0].Object)
	assert.Equal(t, s, xs[1].Object)
}

func TestSphereDefaultTransformation(t *testing.T) {
	s := NewSphere()

	assert.Equal(t, NewIdentityMatrix(), s.Transform)
}

func TestChangeSphereTransformation(t *testing.T) {
	s := NewSphere()
	transform := NewIdentityMatrix().Translate(2, 3, 4)

	s.Transform = transform

	assert.Equal(t, transform, s.Transform)
}

func TestIntersectingScaledSphereWithRay(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := NewSphere()

	s.Transform = NewIdentityMatrix().Scale(2, 2, 2)
	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 3.0, xs[0].T)
	assert.Equal(t, 7.0, xs[1].T)
}

func TestIntersectingTranslatedSphereWithRay(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := NewSphere()

	s.Transform = NewIdentityMatrix().Translate(5, 0, 0)
	xs := s.Intersects(r)

	assert.Equal(t, 0, len(xs))
}

func TestSphereNormalXAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(1, 0, 0))

	assert.Equal(t, NewVector(1, 0, 0), n)
}

func TestSphereNormalYAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(0, 1, 0))

	assert.Equal(t, NewVector(0, 1, 0), n)
}

func TestSphereNormalZAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(0, 0, 1))

	assert.Equal(t, NewVector(0, 0, 1), n)
}

func TestSphereNormalNonAxialPoint(t *testing.T) {
	s := NewSphere()
	x := math.Sqrt(3) / 3
	n := s.NormalAt(NewPoint(x, x, x))

	assert.Equal(t, NewVector(x, x, x), n)
}

func TestNormalIsNormalizedVector(t *testing.T) {
	s := NewSphere()
	x := math.Sqrt(3) / 3
	n := s.NormalAt(NewPoint(x, x, x))

	assert.Equal(t, n.Normalize(), n)
}

func TestNormalTranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.Transform = s.Transform.Translate(0, 1, 0)
	n := s.NormalAt(NewPoint(0, 1.70711, -0.70711))

	EqualTuple(t, NewVector(0, 0.70711, -0.70711), n)
}

func TestNormalTransformedSphere(t *testing.T) {
	s := NewSphere()
	s.Transform = NewIdentityMatrix()
	s.Transform = s.Transform.Scale(1, 0.5, 1).RotateZ(math.Pi / 5)
	n := s.NormalAt(NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2))

	EqualTuple(t, NewVector(0, 0.97014, -0.24254), n)
}
