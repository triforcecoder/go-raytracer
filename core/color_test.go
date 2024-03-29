package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func EqualColor(t *testing.T, expected Color, actual Color) {
	if expected.Equals(actual) {
		assert.True(t, true)
	} else {
		assert.Equal(t, expected, actual)
	}
}

func TestCreateColor(t *testing.T) {
	color := Color{-0.5, 0.4, 1.7}
	assert.Equal(t, -0.5, color.Red)
	assert.Equal(t, 0.4, color.Green)
	assert.Equal(t, 1.7, color.Blue)
}

func TestAddColor(t *testing.T) {
	color1 := Color{0.9, 0.6, 0.75}
	color2 := Color{0.7, 0.1, 0.25}
	result := Color{1.6, 0.7, 1.0}

	assert.Equal(t, result, color1.Add(color2))
}

func TestSubtractColor(t *testing.T) {
	color1 := Color{0.9, 0.6, 0.75}
	color2 := Color{0.7, 0.1, 0.25}
	result := Color{0.2, 0.5, 0.5}

	assert.True(t, color1.Subtract(color2).Equals(result))
}

func TestMultiplyColorByScalar(t *testing.T) {
	color := Color{0.2, 0.3, 0.4}

	assert.Equal(t, Color{0.4, 0.6, 0.8}, color.MultiplyScalar(2))
}

func TestMultiplyColor(t *testing.T) {
	color1 := Color{1, 0.2, 0.4}
	color2 := Color{0.9, 1, 0.1}
	result := Color{0.9, 0.2, 0.04}

	EqualColor(t, result, color1.Multiply(color2))
}

func TestConstantColors(t *testing.T) {
	assert.Equal(t, Color{0, 0, 0}, Black)
	assert.Equal(t, Color{1, 1, 1}, White)
}
