package main

import "math"

type Pattern interface {
	ColorAt(point Tuple) Color
	GetTransform() Matrix
}

type PatternImpl struct {
	a         Color
	b         Color
	transform Matrix
}

type StripePattern struct {
	PatternImpl
}

type GradientPattern struct {
	PatternImpl
}

type RingPattern struct {
	PatternImpl
}

type CheckersPattern struct {
	PatternImpl
}

func NewStripePattern(a Color, b Color) StripePattern {
	return StripePattern{PatternImpl{a, b, NewIdentityMatrix()}}
}

func NewGradientPattern(a Color, b Color) GradientPattern {
	return GradientPattern{PatternImpl{a, b, NewIdentityMatrix()}}
}

func NewRingPattern(a Color, b Color) RingPattern {
	return RingPattern{PatternImpl{a, b, NewIdentityMatrix()}}
}

func NewCheckersPattern(a Color, b Color) CheckersPattern {
	return CheckersPattern{PatternImpl{a, b, NewIdentityMatrix()}}
}

func PatternColor(pattern Pattern, object Shape, worldPoint Tuple) Color {
	objectPoint := object.GetTransform().Inverse().MultiplyTuple(worldPoint)
	patternPoint := pattern.GetTransform().Inverse().MultiplyTuple(objectPoint)

	return pattern.ColorAt(patternPoint)
}

func (pattern PatternImpl) GetTransform() Matrix {
	return pattern.transform
}

func (pattern StripePattern) ColorAt(point Tuple) Color {
	if int(math.Floor(point.x))%2 == 0 {
		return pattern.a
	} else {
		return pattern.b
	}
}

func (pattern GradientPattern) ColorAt(point Tuple) Color {
	distance := pattern.b.Subtract(pattern.a)
	fraction := point.x - math.Floor(point.x)

	return pattern.a.Add(distance.MultiplyScalar(fraction))
}

func (pattern RingPattern) ColorAt(point Tuple) Color {
	if int(math.Floor(math.Sqrt(math.Pow(point.x, 2)+math.Pow(point.z, 2))))%2 == 0 {
		return pattern.a
	} else {
		return pattern.b
	}
}

func (pattern CheckersPattern) ColorAt(point Tuple) Color {
	if int(math.Abs(point.x)+math.Abs(point.y)+math.Abs(point.z))%2 == 0 {
		return pattern.a
	} else {
		return pattern.b
	}
}
