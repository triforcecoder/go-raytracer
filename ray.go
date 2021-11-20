package main

import "math"

type Ray struct {
	origin    Tuple
	direction Tuple
}

type Sphere struct {
	origin    Tuple
	transform Matrix
}

type Intersection struct {
	t      float64
	object Sphere
}

func NewSphere() Sphere {
	return Sphere{NewPoint(0, 0, 0), NewIdentityMatrix()}
}

func NewIntersection(t float64, s Sphere) Intersection {
	return Intersection{t, s}
}

func (ray Ray) Position(t float64) Tuple {
	return ray.origin.Add(ray.direction.Multiply(t))
}

func (ray Ray) Transform(m Matrix) Ray {
	result := Ray{}

	result.origin = m.MultiplyTuple(ray.origin)
	result.direction = m.MultiplyTuple(ray.direction)

	return result
}

func (sphere Sphere) Intersects(ray Ray) []Intersection {
	xs := []Intersection{}

	ray = ray.Transform(sphere.transform.Inverse())
	sphereToRay := ray.origin.Subtract(sphere.origin)

	a := ray.direction.Dot(ray.direction)
	b := 2 * ray.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	// no intersections
	if discriminant < 0 {
		return xs
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	xs = append(xs, Intersection{t1, sphere})
	xs = append(xs, Intersection{t2, sphere})

	return xs
}

func Hit(intersections []Intersection) *Intersection {
	var x *Intersection

	for _, intersection := range intersections {
		if intersection.t > 0 {
			if x == nil {
				x = new(Intersection)
				*x = intersection
			} else if intersection.t < x.t {
				*x = intersection
			}
		}
	}

	return x
}
