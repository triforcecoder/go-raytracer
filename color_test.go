package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateColor(t *testing.T) {
	color := Color{-0.5, 0.4, 1.7}
	assert.Equal(t, -0.5, color.red)
	assert.Equal(t, 0.4, color.green)
	assert.Equal(t, 1.7, color.blue)
}

func TestAddColor(t *testing.T) {
	color1 := Color{0.9, 0.6, 0.75}
	color2 := Color{0.7, 0.1, 0.25}
	result := Color{1.6, 0.7, 1.0}

	assert.Equal(t, result, color1.add(color2))
}

func TestSubtractColor(t *testing.T) {
	color1 := Color{0.9, 0.6, 0.75}
	color2 := Color{0.7, 0.1, 0.25}
	result := Color{0.2, 0.5, 0.5}

	assert.True(t, color1.subtract(color2).equals(result))
}

func TestMultiplyColorByScalar(t *testing.T) {
	color := Color{0.2, 0.3, 0.4}

	assert.Equal(t, Color{0.4, 0.6, 0.8}, color.multiplyScalar(2))
}

func TestMultiplyColor(t *testing.T) {
	color1 := Color{1, 0.2, 0.4}
	color2 := Color{0.9, 1, 0.1}
	result := Color{0.9, 0.2, 0.04}

	assert.True(t, color1.multiply(color2).equals(result))
}

func TestScaleFloat(t *testing.T) {
	assert.Equal(t, 0, scaleFloat(-1))
	assert.Equal(t, 0, scaleFloat(-0.5))
	assert.Equal(t, 0, scaleFloat(0))
	assert.Equal(t, 255, scaleFloat(1.6))
	assert.Equal(t, 255, scaleFloat(1))
	assert.Equal(t, 128, scaleFloat(0.5))
}
