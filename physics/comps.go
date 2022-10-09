package physics

import (
	. "go-raytracer/core"
	. "go-raytracer/geometry"
	"math"
)

// stores precomputed computations
type Comps struct {
	t          float64
	object     Shape
	point      Tuple
	eyev       Tuple
	normalv    Tuple
	reflectv   Tuple
	inside     bool
	overPoint  Tuple
	underPoint Tuple
	n1         float64
	n2         float64
}

func PrepareComputations(intersection Intersection, ray Ray, xs []Intersection) Comps {
	comps := Comps{}

	comps.t = intersection.T
	comps.object = intersection.Object

	comps.point = ray.Position(comps.t)
	comps.eyev = ray.Direction.Multiply(-1)
	comps.normalv = comps.object.NormalAt(comps.point)
	comps.reflectv = ray.Direction.Reflect(comps.normalv)

	if comps.normalv.Dot(comps.eyev) < 0 {
		comps.inside = true
		comps.normalv = comps.normalv.Multiply(-1)
	}

	comps.overPoint = comps.point.Add(comps.normalv.Multiply(Epsilon))
	comps.underPoint = comps.point.Subtract(comps.normalv.Multiply(Epsilon))

	var containers []Shape

	for _, x := range xs {
		if x == intersection {
			if len(containers) == 0 {
				comps.n1 = 1
			} else {
				last := len(containers) - 1
				comps.n1 = containers[last].GetMaterial().RefractiveIndex
			}
		}

		find := -1
		for index, object := range containers {
			if object == x.Object {
				find = index
				break
			}
		}

		if find == -1 {
			containers = append(containers, x.Object)
		} else {
			// remove found element
			containers = append(containers[:find], containers[find+1:]...)
		}

		if x == intersection {
			if len(containers) == 0 {
				comps.n2 = 1
			} else {
				last := len(containers) - 1
				comps.n2 = containers[last].GetMaterial().RefractiveIndex
			}

			break
		}
	}

	return comps
}

func (comps Comps) Schlick() float64 {
	cos := comps.eyev.Dot(comps.normalv)

	if comps.n1 > comps.n2 {
		n := comps.n1 / comps.n2
		sin2T := math.Pow(n, 2) * (1.0 - math.Pow(cos, 2))

		if sin2T > 1 {
			return 1
		}

		cos = math.Sqrt(1.0 - sin2T)
	}

	r0 := math.Pow((comps.n1-comps.n2)/(comps.n1+comps.n2), 2)
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
