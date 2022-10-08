package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	world := World{}

	assert.Nil(t, world.light)
	assert.Nil(t, world.objects)
}

func TestDefaultWorld(t *testing.T) {
	light := &PointLight{NewPoint(-10, 10, -10), white}
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

	color := world.ColorAt(ray, 4)

	EqualColor(t, black, color)
}

func TestColorWhenRayHits(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}

	color := world.ColorAt(ray, 4)

	EqualColor(t, Color{0.38066, 0.47583, 0.2855}, color)
}

func TestColorWhenIntersectionBehindRay(t *testing.T) {
	world := DefaultWorld()
	outerMaterial := world.objects[0].GetMaterial()
	outerMaterial.ambient = 1
	world.objects[0].SetMaterial(outerMaterial)
	innerMaterial := world.objects[0].GetMaterial()
	innerMaterial.ambient = 1
	world.objects[1].SetMaterial(innerMaterial)

	ray := Ray{NewPoint(0, 0, 0.75), NewVector(0, 0, -1)}
	color := world.ColorAt(ray, 4)

	EqualColor(t, innerMaterial.color, color)
}

func TestShadeHitIntersectionInShadow(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	s2.transform = s2.transform.Translate(0, 0, 10)

	world := World{}
	world.light = &PointLight{NewPoint(0, 0, -10), white}
	world.objects = make([]Shape, 0)
	world.objects = append(world.objects, s1, s2)

	ray := Ray{NewPoint(0, 0, 5), NewVector(0, 0, 1)}
	intersection := Intersection{4, s2}
	comps := PrepareComputations(intersection, ray, []Intersection{})

	EqualColor(t, Color{0.1, 0.1, 0.1}, world.ShadeHit(comps, 4))
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

func TestShadingIntersection(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	shape := world.objects[0]
	intersection := NewIntersection(4, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ShadeHit(comps, 4)

	EqualColor(t, Color{0.38066, 0.47583, 0.2855}, color)
}

func TestShadingIntersectionFromInside(t *testing.T) {
	world := DefaultWorld()
	world.light = &PointLight{NewPoint(0, 0.25, 0), white}
	ray := Ray{NewPoint(0, 0, 0), NewVector(0, 0, 1)}
	shape := world.objects[1]
	intersection := NewIntersection(0.5, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ShadeHit(comps, 4)

	EqualColor(t, Color{0.90498, 0.90498, 0.90498}, color)
}

func TestReflectedColorOfNonReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{NewPoint(0, 0, 0), NewVector(0, 0, 1)}
	world.objects[1].SetMaterial(Material{ambient: 1})
	intersection := NewIntersection(math.Sqrt2, world.objects[1])
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ReflectedColor(comps, 4)

	EqualColor(t, black, color)
}

func TestReflectedColorOfReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	shape := NewPlane()
	shape.material.reflective = 0.5
	shape.transform = shape.transform.Translate(0, -1, 0)
	world.objects = append(world.objects, shape)
	ray := Ray{NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2)}
	intersection := NewIntersection(math.Sqrt2, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ReflectedColor(comps, 4)

	EqualColor(t, Color{0.19033, 0.23791, 0.14274}, color)
}

func TestShadeHitOfReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	shape := NewPlane()
	shape.material.reflective = 0.5
	shape.transform = shape.transform.Translate(0, -1, 0)
	world.objects = append(world.objects, shape)
	ray := Ray{NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2)}
	intersection := NewIntersection(math.Sqrt2, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ShadeHit(comps, 4)

	EqualColor(t, Color{0.87676, 0.92434, 0.82918}, color)
}

func TestColorAtMutuallyReflectiveSurfaces(t *testing.T) {
	world := World{}
	world.light = &PointLight{NewPoint(0, 0, 0), Color{1, 1, 1}}

	lower := NewPlane()
	lower.material.reflective = 1
	lower.transform = lower.transform.Translate(0, -1, 0)

	upper := NewPlane()
	upper.material.reflective = 1
	upper.transform = upper.transform.Translate(0, 1, 0)

	world.objects = make([]Shape, 0)
	world.objects = append(world.objects, lower, upper)

	ray := Ray{NewPoint(0, 0, 0), NewVector(0, 1, 0)}

	// ColorAt should terminate without stack overflow
	world.ColorAt(ray, 4)
}

func TestReflectedColorMaxRecursiveDepth(t *testing.T) {
	world := DefaultWorld()
	shape := NewPlane()
	shape.material.reflective = 0.5
	shape.transform = shape.transform.Translate(0, -1, 0)
	world.objects = append(world.objects, shape)
	ray := Ray{NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2)}
	intersection := NewIntersection(math.Sqrt2, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ReflectedColor(comps, 0)

	EqualColor(t, black, color)
}

func TestRefractedColorOpaqueObject(t *testing.T) {
	world := DefaultWorld()
	shape := world.objects[0]
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	xs := []Intersection{NewIntersection(4, shape), NewIntersection(6, shape)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.RefractedColor(comps, 5)

	assert.Equal(t, black, color)
}

func TestRefractedColorMaxRecursiveDepth(t *testing.T) {
	world := DefaultWorld()
	shape := world.objects[0]
	material := shape.GetMaterial()
	material.transparency = 1
	material.refractiveIndex = 1.5
	shape.SetMaterial(material)
	ray := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	xs := []Intersection{NewIntersection(4, shape), NewIntersection(6, shape)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.RefractedColor(comps, 0)

	assert.Equal(t, black, color)
}

func TestRefractedColorUnderTotalInternalReflection(t *testing.T) {
	world := DefaultWorld()
	shape := world.objects[0]
	material := shape.GetMaterial()
	material.transparency = 1
	material.refractiveIndex = 1.5
	shape.SetMaterial(material)
	ray := Ray{NewPoint(0, 0, math.Sqrt2/2), NewVector(0, 1, 0)}
	xs := []Intersection{NewIntersection(-math.Sqrt2/2, shape), NewIntersection(math.Sqrt2/2, shape)}
	comps := PrepareComputations(xs[1], ray, xs) // xs[1] is inside the sphere

	color := world.RefractedColor(comps, 5)

	assert.Equal(t, black, color)
}

func TestRefractedColorWithRefractedRay(t *testing.T) {
	world := DefaultWorld()

	a := world.objects[0]
	materialA := a.GetMaterial()
	materialA.ambient = 1
	materialA.pattern = NewTestPattern()
	a.SetMaterial(materialA)

	b := world.objects[1]
	materialB := b.GetMaterial()
	materialB.transparency = 1
	materialB.refractiveIndex = 1.5
	b.SetMaterial(materialB)

	ray := Ray{NewPoint(0, 0, 0.1), NewVector(0, 1, 0)}
	xs := []Intersection{
		NewIntersection(-0.9899, a),
		NewIntersection(-0.4899, b),
		NewIntersection(0.4899, b),
		NewIntersection(0.9899, a)}
	comps := PrepareComputations(xs[2], ray, xs)

	color := world.RefractedColor(comps, 5)

	EqualColor(t, Color{0, 0.99888, 0.04722}, color)
}

func TestShadeHitWithTransparentMaterial(t *testing.T) {
	world := DefaultWorld()

	floor := NewPlane()
	floor.transform = floor.transform.Translate(0, -1, 0)
	floor.material.transparency = 0.5
	floor.material.refractiveIndex = 1.5

	ball := NewSphere()
	ball.material.color = Color{1, 0, 0}
	ball.material.ambient = 0.5
	ball.transform = ball.transform.Translate(0, -3.5, -0.5)

	world.objects = append(world.objects, floor, ball)

	ray := Ray{NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2)}
	xs := []Intersection{NewIntersection(math.Sqrt2, floor)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.ShadeHit(comps, 5)

	EqualColor(t, Color{0.93642, 0.68642, 0.68642}, color)
}

func TestShadeHitWithReflectiveTransparentMaterial(t *testing.T) {
	world := DefaultWorld()

	floor := NewPlane()
	floor.transform = floor.transform.Translate(0, -1, 0)
	floor.material.reflective = 0.5
	floor.material.transparency = 0.5
	floor.material.refractiveIndex = 1.5

	ball := NewSphere()
	ball.material.color = Color{1, 0, 0}
	ball.material.ambient = 0.5
	ball.transform = ball.transform.Translate(0, -3.5, -0.5)

	world.objects = append(world.objects, floor, ball)

	ray := Ray{NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2)}
	xs := []Intersection{NewIntersection(math.Sqrt2, floor)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.ShadeHit(comps, 5)

	EqualColor(t, Color{0.93391, 0.69643, 0.69243}, color)
}
