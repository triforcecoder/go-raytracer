package main

// stores precomputed computations
type Comps struct {
	t       float64
	object  Sphere
	point   Tuple
	eyev    Tuple
	normalv Tuple
	inside  bool
}

func PrepareComputations(intersection Intersection, ray Ray) Comps {
	comps := Comps{}

	comps.t = intersection.t
	comps.object = intersection.object

	comps.point = ray.Position(comps.t)
	comps.eyev = ray.direction.Multiply(-1)
	comps.normalv = comps.object.NormalAt(comps.point)

	if comps.normalv.Dot(comps.eyev) < 0 {
		comps.inside = true
		comps.normalv = comps.normalv.Multiply(-1)
	}

	return comps
}
