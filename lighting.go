package main

import "math"

type PointLight struct {
	position  Tuple
	intensity Color
}

type Material struct {
	color           Color
	pattern         Pattern
	ambient         float64
	diffuse         float64
	specular        float64
	shininess       float64
	reflective      float64
	transparency    float64
	refractiveIndex float64
}

func NewMaterial() Material {
	return Material{white, nil, 0.1, 0.9, 0.9, 200, 0, 0, 1}
}

func Lighting(material Material, object Shape, light PointLight, point Tuple, eyev Tuple, normalv Tuple, inShadow bool) Color {
	var color Color
	var diffuse Color
	var specular Color

	if material.pattern != nil {
		color = PatternColor(material.pattern, object, point)
	} else {
		color = material.color
	}

	effectiveColor := color.Multiply(light.intensity)
	lightv := light.position.Subtract(point).Normalize()
	ambient := effectiveColor.MultiplyScalar(material.ambient)

	if !inShadow {
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
	}

	return ambient.Add(diffuse).Add(specular)
}
