package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)

	assert.Equal(t, 3.5, i.T)
	assert.Equal(t, s, i.Object)
}

func TestIntersectionCollection(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)

	xs := []Intersection{i1, i2}

	assert.Equal(t, 1.0, xs[0].T)
	assert.Equal(t, 2.0, xs[1].T)
}

func TestHitAllIntersectionsPositive(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := []Intersection{i2, i1}

	i, err := Hit(xs)

	assert.Equal(t, i1, i)
	assert.Nil(t, err)
}

func TestHitSomeIntersectionsNegative(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	xs := []Intersection{i2, i1}

	i, err := Hit(xs)

	assert.Equal(t, i2, i)
	assert.Nil(t, err)
}

func TestHitAllIntersectionsNegative(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	xs := []Intersection{i2, i1}

	_, err := Hit(xs)

	assert.NotNil(t, err)
}

func TestHitLowestNonNegative(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := []Intersection{i1, i4, i3, i2}

	i, err := Hit(xs)

	assert.Equal(t, i4, i)
	assert.Nil(t, err)
}
