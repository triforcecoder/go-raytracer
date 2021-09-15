package main

import "math"

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (tuple Tuple) isPoint() bool {
	return tuple.w == 1
}

func (tuple Tuple) isVector() bool {
	return tuple.w == 0
}

func (tuple Tuple) equals(other Tuple) bool {
	return equal(tuple.x, other.x) &&
		equal(tuple.y, other.y) &&
		equal(tuple.z, other.z) &&
		equal(tuple.w, other.w)
}

func (tuple Tuple) add(other Tuple) Tuple {
	return Tuple{tuple.x + other.x,
		tuple.y + other.y,
		tuple.z + other.z,
		tuple.w + other.w}
}

func (tuple Tuple) subtract(other Tuple) Tuple {
	return Tuple{tuple.x - other.x,
		tuple.y - other.y,
		tuple.z - other.z,
		tuple.w - other.w}
}

func (tuple Tuple) negate() Tuple {
	return Tuple{tuple.x * -1,
		tuple.y * -1,
		tuple.z * -1,
		tuple.w * -1}
}

func (tuple Tuple) multiply(n float64) Tuple {
	return Tuple{tuple.x * n,
		tuple.y * n,
		tuple.z * n,
		tuple.w * n}
}

func (tuple Tuple) divide(n float64) Tuple {
	return Tuple{tuple.x / n,
		tuple.y / n,
		tuple.z / n,
		tuple.w / n}
}

func (tuple Tuple) magnitude() float64 {
	return math.Sqrt(math.Pow(tuple.x, 2) +
		math.Pow(tuple.y, 2) +
		math.Pow(tuple.z, 2) +
		math.Pow(tuple.w, 2))
}

func (tuple Tuple) normalize() Tuple {
	return tuple.divide(tuple.magnitude())
}

func (tuple Tuple) dot(other Tuple) float64 {
	return tuple.x*other.x +
		tuple.y*other.y +
		tuple.z*other.z +
		tuple.w*other.w
}

func (tuple Tuple) cross(other Tuple) Tuple {
	if !tuple.isVector() || !other.isVector() {
		panic("precondition - cross can only be used with vectors")
	}

	return createVector(tuple.y*other.z-tuple.z*other.y,
		tuple.z*other.x-tuple.x*other.z,
		tuple.x*other.y-tuple.y*other.x)
}

func createPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func createVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func equal(x, y float64) bool {
	const epsilon = 0.00001

	if math.Abs(x-y) < epsilon {
		return true
	} else {
		return false
	}
}
