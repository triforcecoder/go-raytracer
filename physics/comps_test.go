package physics

import (
	. "go-raytracer/core"
	. "go-raytracer/geometry"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	index uint
	n1    float64
	n2    float64
}

func TestPrecomputingIntersection(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	intersection := NewIntersection(4, shape)

	comps := PrepareComputations(intersection, ray, []Intersection{})

	assert.Equal(t, intersection.T, comps.t)
	assert.Equal(t, intersection.Object, comps.object)
	assert.Equal(t, NewPoint(0, 0, -1), comps.point)
	assert.Equal(t, NewVector(0, 0, -1), comps.eyev)
	assert.Equal(t, NewVector(0, 0, -1), comps.normalv)
}

func TestHitWhenIntersectionOutside(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	intersection := NewIntersection(4, shape)

	comps := PrepareComputations(intersection, ray, []Intersection{})

	assert.False(t, comps.inside)
}

func TestHitWhenIntersectionInside(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := NewSphere()
	intersection := NewIntersection(1, shape)

	comps := PrepareComputations(intersection, ray, []Intersection{})

	assert.True(t, comps.inside)
	assert.Equal(t, NewPoint(0, 0, 1), comps.point)
	assert.Equal(t, NewVector(0, 0, -1), comps.eyev)
	assert.Equal(t, NewVector(0, 0, -1), comps.normalv)
}

func TestHitOffsetsPoint(t *testing.T) {
	const epsilon = 0.00001
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	shape.Transform = shape.Transform.Translate(0, 0, 1)
	intersection := NewIntersection(5, shape)

	comps := PrepareComputations(intersection, ray, []Intersection{})

	assert.True(t, comps.overPoint.Z < -epsilon/2)
	assert.True(t, comps.point.Z > comps.overPoint.Z)
}

func TestPrecomputeReflectionVector(t *testing.T) {
	shape := NewPlane()
	ray := NewRay(NewPoint(0, 1, -1), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	intersection := NewIntersection(math.Sqrt2, shape)

	comps := PrepareComputations(intersection, ray, []Intersection{})

	assert.Equal(t, NewVector(0, math.Sqrt2/2, math.Sqrt2/2), comps.reflectv)
}

func TestN1andN2VariousIntersections(t *testing.T) {
	a := NewGlassSphere()
	a.Transform.Scale(2, 2, 2)
	a.Material.RefractiveIndex = 1.5

	b := NewGlassSphere()
	b.Transform.Translate(0, 0, -0.25)
	b.Material.RefractiveIndex = 2

	c := NewGlassSphere()
	c.Transform.Translate(0, 0, 0.25)
	c.Material.RefractiveIndex = 2.5

	ray := NewRay(NewPoint(0, 0, -4), NewVector(0, 0, 1))
	xs := []Intersection{
		NewIntersection(2, a),
		NewIntersection(2.75, b),
		NewIntersection(3.25, c),
		NewIntersection(4.75, b),
		NewIntersection(5.25, c),
		NewIntersection(6, a)}

	examples := []Example{
		{0, 1.0, 1.5},
		{1, 1.5, 2.0},
		{2, 2.0, 2.5},
		{3, 2.5, 2.5},
		{4, 2.5, 1.5},
		{5, 1.5, 1.0}}

	for _, example := range examples {
		comps := PrepareComputations(xs[example.index], ray, xs)

		assert.Equal(t, example.n1, comps.n1)
		assert.Equal(t, example.n2, comps.n2)
	}
}

func TestUnderPointOffsetBelowSurface(t *testing.T) {
	const epsilon = 0.00001
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewGlassSphere()
	shape.Transform = shape.Transform.Translate(0, 0, 1)
	intersection := NewIntersection(5, shape)
	xs := []Intersection{intersection}

	comps := PrepareComputations(intersection, ray, xs)

	assert.Greater(t, comps.underPoint.Z, epsilon/2)
	assert.Less(t, comps.point.Z, comps.underPoint.Z)
}

func TestSchlickApproxUnderTotalInternalReflection(t *testing.T) {
	shape := NewGlassSphere()
	ray := NewRay(NewPoint(0, 0, math.Sqrt2/2), NewVector(0, 1, 0))
	xs := []Intersection{
		NewIntersection(-math.Sqrt2/2, shape),
		NewIntersection(math.Sqrt2/2, shape)}
	comps := PrepareComputations(xs[1], ray, xs)

	reflectance := comps.Schlick()

	assert.Equal(t, 1.0, reflectance)
}

func TestSchlickApproxWithPerpendicularViewlingAngle(t *testing.T) {
	shape := NewGlassSphere()
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 1, 0))
	xs := []Intersection{
		NewIntersection(-1, shape),
		NewIntersection(1, shape)}
	comps := PrepareComputations(xs[1], ray, xs)

	reflectance := comps.Schlick()

	assert.True(t, FloatEquals(reflectance, 0.04))
}

func TestSchlickApproxWithSmallAngle(t *testing.T) {
	shape := NewGlassSphere()
	ray := NewRay(NewPoint(0, 0.99, -2), NewVector(0, 0, 1))
	xs := []Intersection{NewIntersection(1.8589, shape)}
	comps := PrepareComputations(xs[0], ray, xs)

	reflectance := comps.Schlick()

	assert.True(t, FloatEquals(reflectance, 0.48873))
}
