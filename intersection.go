package main

import "errors"

type Intersection struct {
	t      float64
	object Sphere
}

func NewIntersection(t float64, s Sphere) Intersection {
	return Intersection{t, s}
}

func Hit(intersections []Intersection) (Intersection, error) {
	var x Intersection
	var found bool

	for _, intersection := range intersections {
		if intersection.t > 0 {
			if !found {
				x = intersection
				found = true
			} else if intersection.t < x.t {
				x = intersection
			}
		}
	}

	if !found {
		return x, errors.New("hit not found")
	}

	return x, nil
}
