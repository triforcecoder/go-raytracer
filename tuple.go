package main

import "math"

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func (tuple Tuple) IsPoint() bool {
	return tuple.w == 1
}

func (tuple Tuple) IsVector() bool {
	return tuple.w == 0
}

func (tuple Tuple) Equals(other Tuple) bool {
	return floatEquals(tuple.x, other.x) &&
		floatEquals(tuple.y, other.y) &&
		floatEquals(tuple.z, other.z) &&
		floatEquals(tuple.w, other.w)
}

func (tuple Tuple) Add(other Tuple) Tuple {
	return Tuple{tuple.x + other.x,
		tuple.y + other.y,
		tuple.z + other.z,
		tuple.w + other.w}
}

func (tuple Tuple) Subtract(other Tuple) Tuple {
	return Tuple{tuple.x - other.x,
		tuple.y - other.y,
		tuple.z - other.z,
		tuple.w - other.w}
}

func (tuple Tuple) Negate() Tuple {
	return Tuple{tuple.x * -1,
		tuple.y * -1,
		tuple.z * -1,
		tuple.w * -1}
}

func (tuple Tuple) Multiply(n float64) Tuple {
	return Tuple{tuple.x * n,
		tuple.y * n,
		tuple.z * n,
		tuple.w * n}
}

func (tuple Tuple) Divide(n float64) Tuple {
	return Tuple{tuple.x / n,
		tuple.y / n,
		tuple.z / n,
		tuple.w / n}
}

func (tuple Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(tuple.x, 2) +
		math.Pow(tuple.y, 2) +
		math.Pow(tuple.z, 2) +
		math.Pow(tuple.w, 2))
}

func (tuple Tuple) Normalize() Tuple {
	return tuple.Divide(tuple.Magnitude())
}

func (tuple Tuple) Dot(other Tuple) float64 {
	return tuple.x*other.x +
		tuple.y*other.y +
		tuple.z*other.z +
		tuple.w*other.w
}

func (tuple Tuple) Cross(other Tuple) Tuple {
	if !tuple.IsVector() || !other.IsVector() {
		panic("precondition - Cross can only be used with vectors")
	}

	return NewVector(tuple.y*other.z-tuple.z*other.y,
		tuple.z*other.x-tuple.x*other.z,
		tuple.x*other.y-tuple.y*other.x)
}

func floatEquals(x, y float64) bool {
	const epsilon = 0.00001

	if math.Abs(x-y) < epsilon {
		return true
	} else {
		return false
	}
}
