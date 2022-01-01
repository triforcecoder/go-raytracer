package main

type Shape interface {
	NormalAt(point Tuple) Tuple
	Intersects(ray Ray) []Intersection
	GetMaterial() Material
	SetMaterial(material Material)
}
