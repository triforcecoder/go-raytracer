package main

import "math"

type Plane struct {
	origin    Tuple
	transform Matrix
	material  Material
}

func NewPlane() *Plane {
	return &Plane{NewPoint(0, 0, 0), NewIdentityMatrix(), NewMaterial()}
}

func (plane *Plane) NormalAt(point Tuple) Tuple {
	return NewVector(0, 1, 0)
}

func (plane *Plane) Intersects(ray Ray) []Intersection {
	xs := []Intersection{}

	ray = ray.Transform(plane.transform.Inverse())

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
