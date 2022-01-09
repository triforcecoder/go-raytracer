package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrecomputingIntersection(t *testing.T) {
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	shape := NewSphere()
	i := NewIntersection(4, shape)

	comps := PrepareComputations(i, ray)

	assert.Equal(t, i.t, comps.t)
	assert.Equal(t, i.object, comps.object)
	assert.Equal(t, NewPoint(0, 0, -1), comps.point)
	assert.Equal(t, NewVector(0, 0, -1), comps.eyev)
	assert.Equal(t, NewVector(0, 0, -1), comps.normalv)
}

func TestHitWhenIntersectionOutside(t *testing.T) {
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	shape := NewSphere()
	i := NewIntersection(4, shape)

	comps := PrepareComputations(i, ray)

	assert.False(t, comps.inside)
}

func TestHitWhenIntersectionInside(t *testing.T) {
	ray := Ray{NewPoint(0, 0, 0), NewVector(0, 0, 1)}
	shape := NewSphere()
	i := NewIntersection(1, shape)

	comps := PrepareComputations(i, ray)

	assert.True(t, comps.inside)
	assert.Equal(t, NewPoint(0, 0, 1), comps.point)
	assert.Equal(t, NewVector(0, 0, -1), comps.eyev)
	assert.Equal(t, NewVector(0, 0, -1), comps.normalv)
}

func TestHitOffsetsPoint(t *testing.T) {
	const epsilon = 0.00001
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	shape := NewSphere()
	shape.transform = shape.transform.Translate(0, 0, 1)
	intersection := Intersection{5, shape}

	comps := PrepareComputations(intersection, ray)

	assert.True(t, comps.overPoint.z < -epsilon/2)
	assert.True(t, comps.point.z > comps.overPoint.z)
}

func TestShadingIntersection(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	shape := world.objects[0]
	i := NewIntersection(4, shape)

	comps := PrepareComputations(i, ray)
	color := world.ShadeHit(comps)

	EqualColor(t, Color{0.38066, 0.47583, 0.2855}, color)
}

func TestShadingIntersectionFromInside(t *testing.T) {
	world := DefaultWorld()
	world.light = &PointLight{NewPoint(0, 0.25, 0), white}
	ray := Ray{NewPoint(0, 0, 0), NewVector(0, 0, 1)}
	shape := world.objects[1]
	i := NewIntersection(0.5, shape)

	comps := PrepareComputations(i, ray)
	color := world.ShadeHit(comps)

	EqualColor(t, Color{0.90498, 0.90498, 0.90498}, color)
}
