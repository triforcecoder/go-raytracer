package geometry

import . "go-raytracer/core"

type Material struct {
	Color           Color
	Pattern         Pattern
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64
}

func NewMaterial() Material {
	return Material{White, nil, 0.1, 0.9, 0.9, 200, 0, 0, 1}
}
