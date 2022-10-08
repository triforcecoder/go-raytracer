package physics

import (
	. "go-raytracer/core"
	. "go-raytracer/geometry"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointLight(t *testing.T) {
	intensity := White
	position := NewPoint(0, 0, 0)
	light := PointLight{position, intensity}

	assert.Equal(t, position, light.position)
	assert.Equal(t, intensity, light.intensity)
}

func TestDefaultMaterial(t *testing.T) {
	material := NewMaterial()

	assert.Equal(t, White, material.Color)
	assert.Equal(t, 0.1, material.Ambient)
	assert.Equal(t, 0.9, material.Diffuse)
	assert.Equal(t, 0.9, material.Specular)
	assert.Equal(t, 200.0, material.Shininess)
	assert.Equal(t, 0.0, material.Reflective)
	assert.Equal(t, 0.0, material.Transparency)
	assert.Equal(t, 1.0, material.RefractiveIndex)
}

func TestLightingWithEyeBetweenLightAndSurface(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), White}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	assert.Equal(t, NewColor(1.9, 1.9, 1.9), result)
}

func TestLightingWithEyeBetweenLightAndSurfaceAndEyeOffset45(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), White}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	assert.Equal(t, White, result)
}

func TestLightingWithEyeOppositeSurfaceAndEyeOffset45(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), White}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	EqualColor(t, NewColor(0.7364, 0.7364, 0.7364), result)
}

func TestLightingWithEyeInPathOfReflectionVector(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), White}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	EqualColor(t, NewColor(1.6364, 1.6364, 1.6364), result)
}

func TestLightingWithLightBehindSurface(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, 10), White}

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, false)

	assert.Equal(t, NewColor(0.1, 0.1, 0.1), result)
}

func TestLightingWithSurfaceInShadow(t *testing.T) {
	material := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), White}
	inShadow := true

	result := Lighting(material, NewSphere(), light, position, eyev, normalv, inShadow)

	assert.Equal(t, NewColor(0.1, 0.1, 0.1), result)
}

func TestLightingWithPatternApplied(t *testing.T) {
	pattern := NewStripePattern(NewSolidPattern(White), NewSolidPattern(Black))
	material := Material{}
	material.Pattern = pattern
	material.Ambient = 1
	material.Diffuse = 0
	material.Specular = 0

	eyev := NewVector(0, 0, -1)
	normalv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), White}
	inShadow := false

	c1 := Lighting(material, NewSphere(), light, NewPoint(0.9, 0, 0), eyev, normalv, inShadow)
	c2 := Lighting(material, NewSphere(), light, NewPoint(1.1, 0, 0), eyev, normalv, inShadow)

	assert.Equal(t, White, c1)
	assert.Equal(t, Black, c2)
}
