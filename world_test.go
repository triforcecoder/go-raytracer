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
	light := &PointLight{NewPoint(-10, 10, -10), Color{1, 1, 1}}
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

	EqualColor(t, Color{0, 0, 0}, color)
}

func TestColorWhenRayHits(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}

	color := world.ColorAt(ray)

	EqualColor(t, Color{0.38066, 0.47583, 0.2855}, color)
}

func TestColorWhenIntersectionBehindRay(t *testing.T) {
	world := DefaultWorld()
	outer := &world.objects[0]
	outer.material.ambient = 1
	inner := &world.objects[1]
	inner.material.ambient = 1
	ray := Ray{NewPoint(0, 0, 0.75), NewVector(0, 0, -1)}

	color := world.ColorAt(ray)

	EqualColor(t, inner.material.color, color)
}
