package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)
	r := Ray{origin, direction}

	assert.Equal(t, origin, r.origin)
	assert.Equal(t, direction, r.direction)
}

func TestPosition(t *testing.T) {
	origin := NewPoint(2, 3, 4)
	direction := NewVector(1, 0, 0)
	r := Ray{origin, direction}

	assert.Equal(t, NewPoint(2, 3, 4), r.Position(0))
	assert.Equal(t, NewPoint(3, 3, 4), r.Position(1))
	assert.Equal(t, NewPoint(1, 3, 4), r.Position(-1))
	assert.Equal(t, NewPoint(4.5, 3, 4), r.Position(2.5))
}

func TestIntersectsSphereAtTwoPoints(t *testing.T) {
	origin := NewPoint(0, 0, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := CreateSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 4.0, xs[0].t)
	assert.Equal(t, 6.0, xs[1].t)
}

func TestIntersectsSphereAtTangent(t *testing.T) {
	origin := NewPoint(0, -1, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := CreateSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 5.0, xs[0].t)
	assert.Equal(t, 5.0, xs[1].t)
}

func TestMissesSphere(t *testing.T) {
	origin := NewPoint(0, 2, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := CreateSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayOriginInsideSphere(t *testing.T) {
	origin := NewPoint(0, 0, 0)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := CreateSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -1.0, xs[0].t)
	assert.Equal(t, 1.0, xs[1].t)
}

func TestSphereBehindRay(t *testing.T) {
	origin := NewPoint(0, 0, 5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := CreateSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -6.0, xs[0].t)
	assert.Equal(t, -4.0, xs[1].t)
}

func TestIntersection(t *testing.T) {
	s := CreateSphere()
	i := CreateIntersection(3.5, s)

	assert.Equal(t, 3.5, i.t)
	assert.Equal(t, s, i.object)
}

func TestIntersectionCollection(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(1, s)
	i2 := CreateIntersection(2, s)

	xs := []Intersection{i1, i2}

	assert.Equal(t, 1.0, xs[0].t)
	assert.Equal(t, 2.0, xs[1].t)
}

func TestIntersectSetsObject(t *testing.T) {
	origin := NewPoint(0, 0, -5)
	direction := NewVector(0, 0, 1)
	r := Ray{origin, direction}
	s := CreateSphere()

	xs := s.Intersects(r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, s, xs[0].object)
	assert.Equal(t, s, xs[1].object)

}

func TestHitAllIntersectionsPositive(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(1, s)
	i2 := CreateIntersection(2, s)
	xs := []Intersection{i2, i1}

	i := Hit(xs)

	assert.Equal(t, i1, *i)
}

func TestHitSomeIntersectionsNegative(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(-1, s)
	i2 := CreateIntersection(1, s)
	xs := []Intersection{i2, i1}

	i := Hit(xs)

	assert.Equal(t, i2, *i)
}

func TestHitAllIntersectionsNegative(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(-2, s)
	i2 := CreateIntersection(-1, s)
	xs := []Intersection{i2, i1}

	i := Hit(xs)

	assert.Nil(t, i)
}

func TestHitLowestNonNegative(t *testing.T) {
	s := CreateSphere()
	i1 := CreateIntersection(5, s)
	i2 := CreateIntersection(7, s)
	i3 := CreateIntersection(-3, s)
	i4 := CreateIntersection(2, s)
	xs := []Intersection{i1, i4, i3, i2}

	i := Hit(xs)

	assert.Equal(t, i4, *i)
}
