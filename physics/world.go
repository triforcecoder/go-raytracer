package physics

import (
	. "go-raytracer/core"
	. "go-raytracer/geometry"
	"math"
	"sort"
)

type World struct {
	Light   *PointLight
	Objects []Shape
}

func DefaultWorld() World {
	s1 := NewSphere()
	material := s1.GetMaterial()
	material.Color = NewColor(0.8, 1.0, 0.6)
	material.Diffuse = 0.7
	material.Specular = 0.2
	s1.SetMaterial(material)
	s2 := NewSphere()
	s2.Transform = s2.Transform.Scale(0.5, 0.5, 0.5)

	world := World{}
	world.Light = &PointLight{NewPoint(-10, 10, -10), White}
	world.Objects = make([]Shape, 0)
	world.Objects = append(world.Objects, s1, s2)

	return world
}

func (world World) Intersect(ray Ray) []Intersection {
	intersections := make([]Intersection, 0)

	for _, object := range world.Objects {
		temp := object.Intersects(ray)

		for _, t := range temp {
			intersections = append(intersections, t)
		}
	}

	sort.Slice(intersections, func(i, j int) bool {
		if intersections[i].T < intersections[j].T {
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
		*world.Light,
		comps.point, comps.eyev, comps.normalv, shadowed)

	reflected := world.ReflectedColor(comps, remaining)
	refracted := world.RefractedColor(comps, remaining)

	material := comps.object.GetMaterial()
	if material.Reflective > 0 && material.Transparency > 0 {
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
		return Black
	}

	comps := PrepareComputations(hit, ray, intersections)
	return world.ShadeHit(comps, remaining)
}

func (world World) IsShadowed(point Tuple) bool {
	vector := world.Light.position.Subtract(point)
	distance := vector.Magnitude()
	direction := vector.Normalize()
	ray := NewRay(point, direction)

	intersection := world.Intersect(ray)
	hit, err := Hit(intersection)

	if err == nil && hit.T < distance {
		return true
	}

	return false
}

func (world World) ReflectedColor(comps Comps, remaining uint) Color {
	if remaining == 0 || comps.object.GetMaterial().Reflective == 0 {
		return Black
	}

	reflectRay := NewRay(comps.overPoint, comps.reflectv)
	color := world.ColorAt(reflectRay, remaining-1)

	return color.MultiplyScalar(comps.object.GetMaterial().Reflective)
}

func (world World) RefractedColor(comps Comps, remaining uint) Color {
	if remaining == 0 || comps.object.GetMaterial().Transparency == 0 {
		return Black
	}

	nRatio := comps.n1 / comps.n2
	cosI := comps.eyev.Dot(comps.normalv)
	sin2T := math.Pow(nRatio, 2) * (1 - math.Pow(cosI, 2))

	if sin2T > 1 {
		// total internal reflection
		return Black
	}

	cosT := math.Sqrt(1.0 - sin2T)
	direction := comps.normalv.Multiply(nRatio*cosI - cosT).Subtract(comps.eyev.Multiply(nRatio))
	refractRay := NewRay(comps.underPoint, direction)
	color := world.ColorAt(refractRay, remaining-1).MultiplyScalar(comps.object.GetMaterial().Transparency)

	return color
}
