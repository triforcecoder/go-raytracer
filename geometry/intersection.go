package geometry

import (
	"errors"
)

type Intersection struct {
	T      float64
	Object Shape
}

func NewIntersection(t float64, s Shape) Intersection {
	return Intersection{t, s}
}

func Hit(intersections []Intersection) (Intersection, error) {
	var x Intersection
	var found bool

	for _, intersection := range intersections {
		if intersection.T > 0 {
			if !found {
				x = intersection
				found = true
			} else if intersection.T < x.T {
				x = intersection
			}
		}
	}

	if !found {
		return x, errors.New("hit not found")
	}

	return x, nil
}
