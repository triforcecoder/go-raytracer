package geometry

import (
	. "go-raytracer/core"
)

type Shape interface {
	NormalAt(point Tuple) Tuple
	Intersects(ray Ray) []Intersection
	GetMaterial() Material
	SetMaterial(material Material)
	GetTransform() Matrix
	GetInverse() Matrix
}
