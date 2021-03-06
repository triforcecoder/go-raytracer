package main

import "math"

type Sphere struct {
	origin        Tuple
	transform     Matrix
	cachedInverse Matrix
	material      Material
}

func NewSphere() *Sphere {
	return &Sphere{NewPoint(0, 0, 0), NewIdentityMatrix(), nil, NewMaterial()}
}

func (sphere *Sphere) Intersects(ray Ray) []Intersection {
	xs := []Intersection{}

	ray = ray.Transform(sphere.GetInverse())
	sphereToRay := ray.origin.Subtract(sphere.origin)

	a := ray.direction.Dot(ray.direction)
	b := 2 * ray.direction.Dot(sphereToRay)
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
	worldNormal.w = 0

	return worldNormal.Normalize()
}

func (sphere *Sphere) GetMaterial() Material {
	return sphere.material
}

func (sphere *Sphere) SetMaterial(material Material) {
	sphere.material = material
}

func (sphere *Sphere) GetTransform() Matrix {
	return sphere.transform
}

func (sphere *Sphere) GetInverse() Matrix {
	if sphere.cachedInverse == nil {
		sphere.cachedInverse = sphere.transform.Inverse()
	}

	return sphere.cachedInverse
}
