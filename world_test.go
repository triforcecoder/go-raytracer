package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	world := World{}

	assert.Nil(t, world.light)
	assert.Nil(t, world.objects)
}

func TestDefaultWorld(t *testing.T) {
	light := &PointLight{NewPoint(-10, 10, -10), white}
	s1 := NewSphere()
	s1.material.color = Color{0.8, 1.0, 0.6}
	s1.material.diffuse = 0.7
	s1.material.specular = 0.2
	s2 := NewSphere()
	s2.transform = s2.transform.Scale(0.5, 0.5, 0.5)

	world := DefaultWorld()

	assert.Equal(t, light, world.light)
	assert.Equal(t, s1, world.objects[0])
	assert.Equal(t, s2, world.objects[1])
}

func TestIntersectWorldWithRay(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}

	xs := world.Intersect(ray)

	assert.Equal(t, 4, len(xs))
	assert.Equal(t, 4.0, xs[0].t)
	assert.Equal(t, 4.5, xs[1].t)
	assert.Equal(t, 5.5, xs[2].t)
	assert.Equal(t, 6.0, xs[3].t)
}

func TestColorWhenRayMisses(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 1, 0)}

	color := world.ColorAt(ray)

	EqualColor(t, black, color)
}

func TestColorWhenRayHits(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}

	color := world.ColorAt(ray)

	EqualColor(t, Color{0.38066, 0.47583, 0.2855}, color)
}

func TestColorWhenIntersectionBehindRay(t *testing.T) {
	world := DefaultWorld()
	outerMaterial := world.objects[0].GetMaterial()
	outerMaterial.ambient = 1
	world.objects[0].SetMaterial(outerMaterial)
	innerMaterial := world.objects[0].GetMaterial()
	innerMaterial.ambient = 1
	world.objects[1].SetMaterial(innerMaterial)

	ray := Ray{NewPoint(0, 0, 0.75), NewVector(0, 0, -1)}
	color := world.ColorAt(ray)

	EqualColor(t, innerMaterial.color, color)
}

func TestShadeHitIntersectionInShadow(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	s2.transform = s2.transform.Translate(0, 0, 10)

	world := World{}
	world.light = &PointLight{NewPoint(0, 0, -10), white}
	world.objects = make([]Shape, 0)
	world.objects = append(world.objects, s1, s2)

	ray := Ray{NewPoint(0, 0, 5), NewVector(0, 0, 1)}
	intersection := Intersection{4, s2}
	comps := PrepareComputations(intersection, ray)

	EqualColor(t, Color{0.1, 0.1, 0.1}, world.ShadeHit(comps))
}

func TestNoShadowWhenNothingCollinearWithPointAndLight(t *testing.T) {
	world := DefaultWorld()
	point := NewPoint(0, 10, 0)

	assert.Equal(t, false, world.IsShadowed(point))
}

func TestShadowWhenObjectBetweenPointAndLight(t *testing.T) {
	world := DefaultWorld()
	point := NewPoint(10, -10, 10)

	assert.Equal(t, true, world.IsShadowed(point))
}

func TestNoShadowWhenObjectBehindLight(t *testing.T) {
	world := DefaultWorld()
	point := NewPoint(-20, 20, -20)

	assert.Equal(t, false, world.IsShadowed(point))
}

func TestNoShadowWhenObjectBehindPoint(t *testing.T) {
	world := DefaultWorld()
	point := NewPoint(-2, 2, -2)

	assert.Equal(t, false, world.IsShadowed(point))
}
