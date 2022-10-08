package physics

import (
	. "go-raytracer/core"
	. "go-raytracer/geometry"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	world := World{}

	assert.Nil(t, world.Light)
	assert.Nil(t, world.Objects)
}

func TestDefaultWorld(t *testing.T) {
	light := &PointLight{NewPoint(-10, 10, -10), White}
	s1 := NewSphere()
	s1.Material.Color = NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	s2 := NewSphere()
	s2.Transform = s2.Transform.Scale(0.5, 0.5, 0.5)

	world := DefaultWorld()

	assert.Equal(t, light, world.Light)
	assert.Equal(t, s1, world.Objects[0])
	assert.Equal(t, s2, world.Objects[1])
}

func TestIntersectWorldWithRay(t *testing.T) {
	world := DefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	xs := world.Intersect(ray)

	assert.Equal(t, 4, len(xs))
	assert.Equal(t, 4.0, xs[0].T)
	assert.Equal(t, 4.5, xs[1].T)
	assert.Equal(t, 5.5, xs[2].T)
	assert.Equal(t, 6.0, xs[3].T)
}

func TestColorWhenRayMisses(t *testing.T) {
	world := DefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))

	color := world.ColorAt(ray, 4)

	EqualColor(t, Black, color)
}

func TestColorWhenRayHits(t *testing.T) {
	world := DefaultWorld()
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	color := world.ColorAt(ray, 4)

	EqualColor(t, NewColor(0.38066, 0.47583, 0.2855), color)
}

func TestColorWhenIntersectionBehindRay(t *testing.T) {
	world := DefaultWorld()
	outerMaterial := world.Objects[0].GetMaterial()
	outerMaterial.Ambient = 1
	world.Objects[0].SetMaterial(outerMaterial)
	innerMaterial := world.Objects[0].GetMaterial()
	innerMaterial.Ambient = 1
	world.Objects[1].SetMaterial(innerMaterial)

	ray := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))
	color := world.ColorAt(ray, 4)

	EqualColor(t, innerMaterial.Color, color)
}

func TestShadeHitIntersectionInShadow(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	s2.Transform = s2.Transform.Translate(0, 0, 10)

	world := World{}
	world.Light = &PointLight{NewPoint(0, 0, -10), White}
	world.Objects = make([]Shape, 0)
	world.Objects = append(world.Objects, s1, s2)

	ray := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	intersection := NewIntersection(4, s2)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	EqualColor(t, NewColor(0.1, 0.1, 0.1), world.ShadeHit(comps, 4))
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
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := world.Objects[0]
	intersection := NewIntersection(4, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ShadeHit(comps, 4)

	EqualColor(t, NewColor(0.38066, 0.47583, 0.2855), color)
}

func TestShadingIntersectionFromInside(t *testing.T) {
	world := DefaultWorld()
	world.Light = &PointLight{NewPoint(0, 0.25, 0), White}
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := world.Objects[1]
	intersection := NewIntersection(0.5, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ShadeHit(comps, 4)

	EqualColor(t, NewColor(0.90498, 0.90498, 0.90498), color)
}

func TestReflectedColorOfNonReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	world.Objects[1].SetMaterial(Material{Ambient: 1})
	intersection := NewIntersection(math.Sqrt2, world.Objects[1])
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ReflectedColor(comps, 4)

	EqualColor(t, Black, color)
}

func TestReflectedColorOfReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.Transform = shape.Transform.Translate(0, -1, 0)
	world.Objects = append(world.Objects, shape)
	ray := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	intersection := NewIntersection(math.Sqrt2, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ReflectedColor(comps, 4)

	EqualColor(t, NewColor(0.19033, 0.23791, 0.14274), color)
}

func TestShadeHitOfReflectiveMaterial(t *testing.T) {
	world := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.Transform = shape.Transform.Translate(0, -1, 0)
	world.Objects = append(world.Objects, shape)
	ray := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	intersection := NewIntersection(math.Sqrt2, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ShadeHit(comps, 4)

	EqualColor(t, NewColor(0.87676, 0.92434, 0.82918), color)
}

func TestColorAtMutuallyReflectiveSurfaces(t *testing.T) {
	world := World{}
	world.Light = &PointLight{NewPoint(0, 0, 0), NewColor(1, 1, 1)}

	lower := NewPlane()
	lower.Material.Reflective = 1
	lower.Transform = lower.Transform.Translate(0, -1, 0)

	upper := NewPlane()
	upper.Material.Reflective = 1
	upper.Transform = upper.Transform.Translate(0, 1, 0)

	world.Objects = make([]Shape, 0)
	world.Objects = append(world.Objects, lower, upper)

	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 1, 0))

	// ColorAt should terminate without stack overflow
	world.ColorAt(ray, 4)
}

func TestReflectedColorMaxRecursiveDepth(t *testing.T) {
	world := DefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.Transform = shape.Transform.Translate(0, -1, 0)
	world.Objects = append(world.Objects, shape)
	ray := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	intersection := NewIntersection(math.Sqrt2, shape)
	comps := PrepareComputations(intersection, ray, []Intersection{})

	color := world.ReflectedColor(comps, 0)

	EqualColor(t, Black, color)
}

func TestRefractedColorOpaqueObject(t *testing.T) {
	world := DefaultWorld()
	shape := world.Objects[0]
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := []Intersection{NewIntersection(4, shape), NewIntersection(6, shape)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.RefractedColor(comps, 5)

	assert.Equal(t, Black, color)
}

func TestRefractedColorMaxRecursiveDepth(t *testing.T) {
	world := DefaultWorld()
	shape := world.Objects[0]
	material := shape.GetMaterial()
	material.Transparency = 1
	material.RefractiveIndex = 1.5
	shape.SetMaterial(material)
	ray := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := []Intersection{NewIntersection(4, shape), NewIntersection(6, shape)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.RefractedColor(comps, 0)

	assert.Equal(t, Black, color)
}

func TestRefractedColorUnderTotalInternalReflection(t *testing.T) {
	world := DefaultWorld()
	shape := world.Objects[0]
	material := shape.GetMaterial()
	material.Transparency = 1
	material.RefractiveIndex = 1.5
	shape.SetMaterial(material)
	ray := NewRay(NewPoint(0, 0, math.Sqrt2/2), NewVector(0, 1, 0))
	xs := []Intersection{NewIntersection(-math.Sqrt2/2, shape), NewIntersection(math.Sqrt2/2, shape)}
	comps := PrepareComputations(xs[1], ray, xs) // xs[1] is inside the sphere

	color := world.RefractedColor(comps, 5)

	assert.Equal(t, Black, color)
}

func TestRefractedColorWithRefractedRay(t *testing.T) {
	world := DefaultWorld()

	a := world.Objects[0]
	materialA := a.GetMaterial()
	materialA.Ambient = 1
	materialA.Pattern = NewTestPattern()
	a.SetMaterial(materialA)

	b := world.Objects[1]
	materialB := b.GetMaterial()
	materialB.Transparency = 1
	materialB.RefractiveIndex = 1.5
	b.SetMaterial(materialB)

	ray := NewRay(NewPoint(0, 0, 0.1), NewVector(0, 1, 0))
	xs := []Intersection{
		NewIntersection(-0.9899, a),
		NewIntersection(-0.4899, b),
		NewIntersection(0.4899, b),
		NewIntersection(0.9899, a)}
	comps := PrepareComputations(xs[2], ray, xs)

	color := world.RefractedColor(comps, 5)

	EqualColor(t, NewColor(0, 0.99888, 0.04722), color)
}

func TestShadeHitWithTransparentMaterial(t *testing.T) {
	world := DefaultWorld()

	floor := NewPlane()
	floor.Transform = floor.Transform.Translate(0, -1, 0)
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := NewSphere()
	ball.Material.Color = NewColor(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.Transform = ball.Transform.Translate(0, -3.5, -0.5)

	world.Objects = append(world.Objects, floor, ball)

	ray := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	xs := []Intersection{NewIntersection(math.Sqrt2, floor)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.ShadeHit(comps, 5)

	EqualColor(t, NewColor(0.93642, 0.68642, 0.68642), color)
}

func TestShadeHitWithReflectiveTransparentMaterial(t *testing.T) {
	world := DefaultWorld()

	floor := NewPlane()
	floor.Transform = floor.Transform.Translate(0, -1, 0)
	floor.Material.Reflective = 0.5
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := NewSphere()
	ball.Material.Color = NewColor(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.Transform = ball.Transform.Translate(0, -3.5, -0.5)

	world.Objects = append(world.Objects, floor, ball)

	ray := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	xs := []Intersection{NewIntersection(math.Sqrt2, floor)}
	comps := PrepareComputations(xs[0], ray, xs)

	color := world.ShadeHit(comps, 5)

	EqualColor(t, NewColor(0.93391, 0.69643, 0.69243), color)
}
