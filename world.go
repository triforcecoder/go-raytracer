package main

import "sort"

type World struct {
	light   *PointLight
	objects []Sphere
}

func DefaultWorld() World {
	s1 := NewSphere()
	s1.material.color = Color{0.8, 1.0, 0.6}
	s1.material.diffuse = 0.7
	s1.material.specular = 0.2
	s2 := NewSphere()
	s2.transform = s2.transform.Scale(0.5, 0.5, 0.5)

	world := World{}
	world.light = &PointLight{NewPoint(-10, 10, -10), Color{1, 1, 1}}
	world.objects = make([]Sphere, 0)
	world.objects = append(world.objects, s1)
	world.objects = append(world.objects, s2)

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

func (world World) ShadeHit(comps Comps) Color {
	return Lighting(
		comps.object.material,
		*world.light,
		comps.point, comps.eyev, comps.normalv)
}

func (world World) ColorAt(ray Ray) Color {
	intersections := world.Intersect(ray)
	hit, err := Hit(intersections)

	if err != nil {
		return Color{0, 0, 0}
	}

	comps := PrepareComputations(hit, ray)
	return world.ShadeHit(comps)
}
