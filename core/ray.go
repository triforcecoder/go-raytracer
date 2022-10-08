package core

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (ray Ray) Position(t float64) Tuple {
	return ray.Origin.Add(ray.Direction.Multiply(t))
}

func (ray Ray) Transform(m Matrix) Ray {
	result := Ray{}

	result.Origin = m.MultiplyTuple(ray.Origin)
	result.Direction = m.MultiplyTuple(ray.Direction)

	return result
}
