package physics

import (
	. "go-raytracer/core"
	. "go-raytracer/geometry"
	"math"
)

type PointLight struct {
	position  Tuple
	intensity Color
}

func NewPointLight(position Tuple, intensity Color) *PointLight {
	return &PointLight{position, intensity}
}

func Lighting(material Material, object Shape, light PointLight, point Tuple, eyev Tuple, normalv Tuple, inShadow bool) Color {
	var color Color
	var diffuse Color
	var specular Color

	if material.Pattern != nil {
		color = PatternColor(material.Pattern, object, point)
	} else {
		color = material.Color
	}

	effectiveColor := color.Multiply(light.intensity)
	lightv := light.position.Subtract(point).Normalize()
	ambient := effectiveColor.MultiplyScalar(material.Ambient)

	if !inShadow {
		lightDotNormal := lightv.Dot(normalv)
		if lightDotNormal >= 0 {
			diffuse = effectiveColor.MultiplyScalar(material.Diffuse).MultiplyScalar(lightDotNormal)

			reflectv := lightv.Multiply(-1).Reflect(normalv)
			reflectDotEye := reflectv.Dot(eyev)

			if reflectDotEye > 0 {
				factor := math.Pow(reflectDotEye, material.Shininess)
				specular = light.intensity.MultiplyScalar(material.Specular).MultiplyScalar(factor)
			}
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
