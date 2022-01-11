package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStripePattern(t *testing.T) {
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, NewSolidPattern(white), pattern.a)
	assert.Equal(t, NewSolidPattern(black), pattern.b)
}

func TestStripePatternConstantY(t *testing.T) {
	object := NewSphere()
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 1, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 2, 0)))
}

func TestStripePatternConstantZ(t *testing.T) {
	object := NewSphere()
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 1)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 2)))
}

func TestStripePatternAlternatesX(t *testing.T) {
	object := NewSphere()
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0.9, 0, 0)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(1, 0, 0)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(-0.1, 0, 0)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(-1, 0, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(-1.1, 0, 0)))
}

func TestStripeWithObjectTransformation(t *testing.T) {
	object := NewSphere()
	object.transform = object.transform.Scale(2, 2, 2)
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(1.5, 0, 0)))
}

func TestStripeWithPatternTransformation(t *testing.T) {
	object := NewSphere()
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))
	pattern.transform = pattern.transform.Scale(2, 2, 2)

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(1.5, 0, 0)))
}

func TestStripeWithObjectAndPatternTransformation(t *testing.T) {
	object := NewSphere()
	object.transform = object.transform.Scale(2, 2, 2)
	pattern := NewStripePattern(NewSolidPattern(white), NewSolidPattern(black))
	pattern.transform = pattern.transform.Translate(0.5, 0, 0)

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(2.5, 0, 0)))
}

func TestGradientLinearlyInterpolatesColors(t *testing.T) {
	object := NewSphere()
	pattern := NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, Color{0.75, 0.75, 0.75}, PatternColor(pattern, object, NewPoint(0.25, 0, 0)))
	assert.Equal(t, Color{0.5, 0.5, 0.5}, PatternColor(pattern, object, NewPoint(0.5, 0, 0)))
	assert.Equal(t, Color{0.25, 0.25, 0.25}, PatternColor(pattern, object, NewPoint(0.75, 0, 0)))
}

func TestRingExtendsBothXandZ(t *testing.T) {
	object := NewSphere()
	pattern := NewRingPattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(1, 0, 0)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(0, 0, 1)))
	// 0.708 = just slightly more than √2/2
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(0.708, 0, 0.708)))
}

func TestCheckersShouldRepeatInX(t *testing.T) {
	object := NewSphere()
	pattern := NewCheckersPattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0.99, 0, 0)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(1.01, 0, 0)))
}

func TestCheckersShouldRepeatInY(t *testing.T) {
	object := NewSphere()
	pattern := NewCheckersPattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0.99, 0)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(0, 1.01, 0)))
}

func TestCheckersShouldRepeatInZ(t *testing.T) {
	object := NewSphere()
	pattern := NewCheckersPattern(NewSolidPattern(white), NewSolidPattern(black))

	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0)))
	assert.Equal(t, white, PatternColor(pattern, object, NewPoint(0, 0, 0.99)))
	assert.Equal(t, black, PatternColor(pattern, object, NewPoint(0, 0, 1.01)))
}
