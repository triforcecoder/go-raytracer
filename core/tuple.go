package core

import "math"

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func NewPoint(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func NewVector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func (tuple Tuple) IsPoint() bool {
	return tuple.W == 1
}

func (tuple Tuple) IsVector() bool {
	return tuple.W == 0
}

func (tuple Tuple) Equals(other Tuple) bool {
	return FloatEquals(tuple.X, other.X) &&
		FloatEquals(tuple.Y, other.Y) &&
		FloatEquals(tuple.Z, other.Z) &&
		FloatEquals(tuple.W, other.W)
}

func (tuple Tuple) Add(other Tuple) Tuple {
	return Tuple{tuple.X + other.X,
		tuple.Y + other.Y,
		tuple.Z + other.Z,
		tuple.W + other.W}
}

func (tuple Tuple) Subtract(other Tuple) Tuple {
	return Tuple{tuple.X - other.X,
		tuple.Y - other.Y,
		tuple.Z - other.Z,
		tuple.W - other.W}
}

func (tuple Tuple) Negate() Tuple {
	return Tuple{tuple.X * -1,
		tuple.Y * -1,
		tuple.Z * -1,
		tuple.W * -1}
}

func (tuple Tuple) Multiply(n float64) Tuple {
	return Tuple{tuple.X * n,
		tuple.Y * n,
		tuple.Z * n,
		tuple.W * n}
}

func (tuple Tuple) Divide(n float64) Tuple {
	return Tuple{tuple.X / n,
		tuple.Y / n,
		tuple.Z / n,
		tuple.W / n}
}

func (tuple Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(tuple.X, 2) +
		math.Pow(tuple.Y, 2) +
		math.Pow(tuple.Z, 2) +
		math.Pow(tuple.W, 2))
}

func (tuple Tuple) Normalize() Tuple {
	return tuple.Divide(tuple.Magnitude())
}

func (tuple Tuple) Dot(other Tuple) float64 {
	return tuple.X*other.X +
		tuple.Y*other.Y +
		tuple.Z*other.Z +
		tuple.W*other.W
}

func (tuple Tuple) Cross(other Tuple) Tuple {
	if !tuple.IsVector() || !other.IsVector() {
		panic("precondition - Cross can only be used with vectors")
	}

	return NewVector(tuple.Y*other.Z-tuple.Z*other.Y,
		tuple.Z*other.X-tuple.X*other.Z,
		tuple.X*other.Y-tuple.Y*other.X)
}

func (tuple Tuple) Reflect(normal Tuple) Tuple {
	// Wish I had operator overloading 😭
	return tuple.Subtract(normal.Multiply(2).Multiply(tuple.Dot(normal)))
}

func FloatEquals(x, y float64) bool {
	const epsilon = 0.00001

	if math.Abs(x-y) < epsilon {
		return true
	} else {
		return false
	}
}
