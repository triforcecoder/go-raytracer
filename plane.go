package main

import "math"

type Plane struct {
	origin        Tuple
	transform     Matrix
	cachedInverse Matrix
	material      Material
}

func NewPlane() *Plane {
	return &Plane{NewPoint(0, 0, 0), NewIdentityMatrix(), nil, NewMaterial()}
}

func (plane *Plane) NormalAt(point Tuple) Tuple {
	return NewVector(0, 1, 0)
}

func (plane *Plane) Intersects(ray Ray) []Intersection {
	xs := []Intersection{}

	ray = ray.Transform(plane.GetInverse())

	const epsilon = 0.00001
	if math.Abs(ray.direction.y) < epsilon {
		return xs
	}

	t := -ray.origin.y / ray.direction.y
	xs = append(xs, Intersection{t, plane})
	return xs
}

func (plane *Plane) GetMaterial() Material {
	return plane.material
}

func (plane *Plane) SetMaterial(material Material) {
	plane.material = material
}

func (plane *Plane) GetTransform() Matrix {
	return plane.transform
}

func (plane *Plane) GetInverse() Matrix {
	if plane.cachedInverse == nil {
		plane.cachedInverse = plane.transform.Inverse()
	}

	return plane.cachedInverse
}
