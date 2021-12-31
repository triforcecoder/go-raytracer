package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointLight(t *testing.T) {
	intensity := Color{1, 1, 1}
	position := NewPoint(0, 0, 0)
	light := PointLight{position, intensity}

	assert.Equal(t, position, light.position)
	assert.Equal(t, intensity, light.intensity)
}

func TestDefaultMaterial(t *testing.T) {
	material := NewMaterial()

	assert.Equal(t, Color{1, 1, 1}, material.color)
	assert.Equal(t, 0.1, material.ambient)
	assert.Equal(t, 0.9, material.diffuse)
	assert.Equal(t, 0.9, material.specular)
	assert.Equal(t, 200.0, material.shininess)
}

func TestLightingWithEyeBetweenLightAndSurface(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), Color{1, 1, 1}}

	result := Lighting(material, light, position, eyev, normalv, false)

	assert.Equal(t, Color{1.9, 1.9, 1.9}, result)
}

func TestLightingWithEyeBetweenLightAndSurfaceAndEyeOffset45(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), Color{1, 1, 1}}

	result := Lighting(material, light, position, eyev, normalv, false)

	assert.Equal(t, Color{1, 1, 1}, result)
}

func TestLightingWithEyeOppositeSurfaceAndEyeOffset45(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), Color{1, 1, 1}}

	result := Lighting(material, light, position, eyev, normalv, false)

	EqualColor(t, Color{0.7364, 0.7364, 0.7364}, result)
}

func TestLightingWithEyeInPathOfReflectionVector(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), Color{1, 1, 1}}

	result := Lighting(material, light, position, eyev, normalv, false)

	EqualColor(t, Color{1.6364, 1.6364, 1.6364}, result)
}

func TestLightingWithLightBehindSurface(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, 10), Color{1, 1, 1}}

	result := Lighting(material, light, position, eyev, normalv, false)

	assert.Equal(t, Color{0.1, 0.1, 0.1}, result)
}

func TestLightingWithSurfaceInShadow(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), Color{1, 1, 1}}
	inShadow := true

	result := Lighting(material, light, position, eyev, normalv, inShadow)

	assert.Equal(t, Color{0.1, 0.1, 0.1}, result)
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
