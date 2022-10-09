package geometry

import (
	. "go-raytracer/core"
	"math"
)

type Plane struct {
	origin        Tuple
	Transform     Matrix
	cachedInverse Matrix
	Material      Material
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

	if math.Abs(ray.Direction.Y) < Epsilon {
		return xs
	}

	t := -ray.Origin.Y / ray.Direction.Y
	xs = append(xs, NewIntersection(t, plane))
	return xs
}

func (plane *Plane) GetMaterial() Material {
	return plane.Material
}

func (plane *Plane) SetMaterial(material Material) {
	plane.Material = material
}

func (plane *Plane) GetTransform() Matrix {
	return plane.Transform
}

func (plane *Plane) GetInverse() Matrix {
	if plane.cachedInverse == nil {
		plane.cachedInverse = plane.Transform.Inverse()
	}

	return plane.cachedInverse
}
