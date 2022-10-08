package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)
	r := Ray{origin, direction}

	assert.Equal(t, origin, r.Origin)
	assert.Equal(t, direction, r.Direction)
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

func TestTranslateRay(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	m := NewIdentityMatrix().Translate(3, 4, 5)

	r2 := r.Transform(m)

	assert.Equal(t, NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, NewVector(0, 1, 0), r2.Direction)
}

func TestScaleRay(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	m := NewIdentityMatrix().Scale(2, 3, 4)

	r2 := r.Transform(m)

	assert.Equal(t, NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, NewVector(0, 3, 0), r2.Direction)
}
