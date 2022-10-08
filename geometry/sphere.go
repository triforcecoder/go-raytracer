package geometry

import (
	. "go-raytracer/core"
	"math"
)

type Sphere struct {
	origin        Tuple
	Transform     Matrix
	cachedInverse Matrix
	Material      Material
}

func NewSphere() *Sphere {
	return &Sphere{NewPoint(0, 0, 0), NewIdentityMatrix(), nil, NewMaterial()}
}

func NewGlassSphere() *Sphere {
	sphere := Sphere{NewPoint(0, 0, 0), NewIdentityMatrix(), nil, NewMaterial()}
	sphere.Material.Transparency = 1
	sphere.Material.RefractiveIndex = 1.5

	return &sphere
}

func (sphere *Sphere) Intersects(ray Ray) []Intersection {
	xs := []Intersection{}

	ray = ray.Transform(sphere.GetInverse())
	sphereToRay := ray.Origin.Subtract(sphere.origin)

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	// no intersections
	if discriminant < 0 {
		return xs
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	xs = append(xs, Intersection{t1, sphere})
	xs = append(xs, Intersection{t2, sphere})

	return xs
}

func (sphere *Sphere) NormalAt(point Tuple) Tuple {
	objectPoint := sphere.GetInverse().MultiplyTuple(point)
	objectNormal := objectPoint.Subtract(sphere.origin)
	worldNormal := sphere.GetInverse().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}

func (sphere *Sphere) GetMaterial() Material {
	return sphere.Material
}

func (sphere *Sphere) SetMaterial(material Material) {
	sphere.Material = material
}

func (sphere *Sphere) GetTransform() Matrix {
	return sphere.Transform
}

func (sphere *Sphere) GetInverse() Matrix {
	if sphere.cachedInverse == nil {
		sphere.cachedInverse = sphere.Transform.Inverse()
	}

	return sphere.cachedInverse
}
