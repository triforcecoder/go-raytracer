package main

import "math"

type PointLight struct {
	position  Tuple
	intensity Color
}

type Material struct {
	color     Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

func NewMaterial() Material {
	return Material{Color{1, 1, 1}, 0.1, 0.9, 0.9, 200}
}

func Lighting(material Material, light PointLight, point Tuple, eyev Tuple, normalv Tuple) Color {
	var diffuse Color
	var specular Color

	effectiveColor := material.color.Multiply(light.intensity)
	lightv := light.position.Subtract(point).Normalize()
	ambient := effectiveColor.MultiplyScalar(material.ambient)

	lightDotNormal := lightv.Dot(normalv)
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.MultiplyScalar(material.diffuse).MultiplyScalar(lightDotNormal)

		reflectv := lightv.Multiply(-1).Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, material.shininess)
			specular = light.intensity.MultiplyScalar(material.specular).MultiplyScalar(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
