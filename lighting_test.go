package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointLight(t *testing.T) {
	intensity := white
	position := NewPoint(0, 0, 0)
	light := PointLight{position, intensity}

	assert.Equal(t, position, light.position)
	assert.Equal(t, intensity, light.intensity)
}

func TestDefaultMaterial(t *testing.T) {
	material := NewMaterial()

	assert.Equal(t, white, material.color)
	assert.Equal(t, 0.1, material.ambient)
	assert.Equal(t, 0.9, material.diffuse)
	assert.Equal(t, 0.9, material.specular)
	assert.Equal(t, 200.0, material.shininess)
	assert.Equal(t, 0.0, material.reflective)
	assert.Equal(t, 0.0, material.transparency)
	assert.Equal(t, 1.0, material.refractiveIndex)
}

func TestLightingWithEyeBetweenLightAndSurface(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), white}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	assert.Equal(t, Color{1.9, 1.9, 1.9}, result)
}

func TestLightingWithEyeBetweenLightAndSurfaceAndEyeOffset45(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), white}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	assert.Equal(t, white, result)
}

func TestLightingWithEyeOppositeSurfaceAndEyeOffset45(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), white}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	EqualColor(t, Color{0.7364, 0.7364, 0.7364}, result)
}

func TestLightingWithEyeInPathOfReflectionVector(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), white}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	EqualColor(t, Color{1.6364, 1.6364, 1.6364}, result)
}

func TestLightingWithLightBehindSurface(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, 10), white}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	assert.Equal(t, Color{0.1, 0.1, 0.1}, result)
}

func TestLightingWithSurfaceInShadow(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), white}
	inShadow := true

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, inShadow)

	assert.Equal(t, Color{0.1, 0.1, 0.1}, result)
}

func TestLightingWithPatternApplied(t *testing.T) {
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))
	material := Material{}
	material.pattern = pattern
	material.ambient = 1
	material.diffuse = 0
	material.specular = 0

	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), white}
	inShadow := false

	c1 := Lighting(material, NewSphere(), light, NewPoint(0.9, 0, 0), eyev, normalv, inShadow)
	c2 := Lighting(material, NewSphere(), light, NewPoint(1.1, 0, 0), eyev, normalv, inShadow)

	assert.Equal(t, white, c1)
	assert.Equal(t, black, c2)
}
