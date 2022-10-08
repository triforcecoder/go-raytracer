package main

import (
	"math"
	"sort"
)

type World struct {
	light   *PointLight
	objects []Shape
}

func DefaultWorld() World {
	s1 := NewSphere()
	s1.material.color = Color{0.8, 1.0, 0.6}
	s1.material.diffuse = 0.7
	s1.material.specular = 0.2
	s2 := NewSphere()
	s2.transform = s2.transform.Scale(0.5, 0.5, 0.5)

	world := World{}
	world.light = &PointLight{NewPoint(-10, 10, -10), white}
	world.objects = make([]Shape, 0)
	world.objects = append(world.objects, s1, s2)

	return world
}

func (world World) Intersect(ray Ray) []Intersection {
	intersections := make([]Intersection, 0)

	for _, object := range world.objects {
		temp := object.Intersects(ray)

		for _, t := range temp {
			intersections = append(intersections, t)
		}
	}

	sort.Slice(intersections, func(i, j int) bool {
		if intersections[i].t < intersections[j].t {
			return true
		}

		return false
	})

	return intersections
}

func (world World) ShadeHit(comps Comps, remaining uint) Color {
	shadowed := world.IsShadowed(comps.overPoint)

	surface := Lighting(
		comps.object.GetMaterial(),
		comps.object,
		*world.light,
		comps.point, comps.eyev, comps.normalv, shadowed)

	reflected := world.ReflectedColor(comps, remaining)
	refracted := world.RefractedColor(comps, remaining)

	material := comps.object.GetMaterial()
	if material.reflective > 0 && material.transparency > 0 {
		reflectance := comps.Schlick()
		return surface.Add(
			reflected.MultiplyScalar(reflectance)).Add(
			refracted.MultiplyScalar(1 - reflectance))
	}

	return surface.Add(reflected).Add(refracted)
}

func (world World) ColorAt(ray Ray, remaining uint) Color {
	intersections := world.Intersect(ray)
	hit, err := Hit(intersections)

	if err != nil {
		return black
	}

	comps := PrepareComputations(hit, ray, intersections)
	return world.ShadeHit(comps, remaining)
}

func (world World) IsShadowed(point Tuple) bool {
	vector := world.light.position.Subtract(point)
	distance := vector.Magnitude()
	direction := vector.Normalize()
	ray := Ray{point, direction}

	intersection := world.Intersect(ray)
	hit, err := Hit(intersection)

	if err == nil && hit.t < distance {
		return true
	}

	return false
}

func (world World) ReflectedColor(comps Comps, remaining uint) Color {
	if remaining == 0 || comps.object.GetMaterial().reflective == 0 {
		return black
	}

	reflectRay := Ray{comps.overPoint, comps.reflectv}
	color := world.ColorAt(reflectRay, remaining-1)

	return color.MultiplyScalar(comps.object.GetMaterial().reflective)
}

func (world World) RefractedColor(comps Comps, remaining uint) Color {
	if remaining == 0 || comps.object.GetMaterial().transparency == 0 {
		return black
	}

	nRatio := comps.n1 / comps.n2
	cosI := comps.eyev.Dot(comps.normalv)
	sin2T := math.Pow(nRatio, 2) * (1 - math.Pow(cosI, 2))

	if sin2T > 1 {
		// total internal reflection
		return black
	}

	cosT := math.Sqrt(1.0 - sin2T)
	direction := comps.normalv.Multiply(nRatio*cosI - cosT).Subtract(comps.eyev.Multiply(nRatio))
	refractRay := Ray{comps.underPoint, direction}
	color := world.ColorAt(refractRay, remaining-1).MultiplyScalar(comps.object.GetMaterial().transparency)

	return color
}
