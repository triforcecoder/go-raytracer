package main

type Ray struct {
	origin    Tuple
	direction Tuple
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
