package main

import "math"

type Pattern interface {
	ColorAt(point Tuple) Color
	GetTransform() Matrix
	GetInverse() Matrix
}

type PatternImpl struct {
	transform     Matrix
	cachedInverse Matrix
}

type SolidPattern struct {
	c Color
	PatternImpl
}

type StripePattern struct {
	a Pattern
	b Pattern
	PatternImpl
}

type GradientPattern struct {
	a Pattern
	b Pattern
	PatternImpl
}

type RingPattern struct {
	a Pattern
	b Pattern
	PatternImpl
}

type CheckersPattern struct {
	a Pattern
	b Pattern
	PatternImpl
}

func NewSolidPattern(c Color) *SolidPattern {
	return &SolidPattern{c, PatternImpl{NewIdentityMatrix(), nil}}
}

func NewStripePattern(a Pattern, b Pattern) *StripePattern {
	return &StripePattern{a, b, PatternImpl{NewIdentityMatrix(), nil}}
}

func NewGradientPattern(a Pattern, b Pattern) *GradientPattern {
	return &GradientPattern{a, b, PatternImpl{NewIdentityMatrix(), nil}}
}

func NewRingPattern(a Pattern, b Pattern) *RingPattern {
	return &RingPattern{a, b, PatternImpl{NewIdentityMatrix(), nil}}
}

func NewCheckersPattern(a Pattern, b Pattern) *CheckersPattern {
	return &CheckersPattern{a, b, PatternImpl{NewIdentityMatrix(), nil}}
}

func PatternColor(pattern Pattern, object Shape, worldPoint Tuple) Color {
	objectPoint := object.GetInverse().MultiplyTuple(worldPoint)
	patternPoint := pattern.GetInverse().MultiplyTuple(objectPoint)

	return pattern.ColorAt(patternPoint)
}

func (pattern *PatternImpl) GetTransform() Matrix {
	return pattern.transform
}

func (pattern *PatternImpl) GetInverse() Matrix {
	if pattern.cachedInverse == nil {
		pattern.cachedInverse = pattern.transform.Inverse()
	}

	return pattern.cachedInverse
}

func (pattern *SolidPattern) ColorAt(point Tuple) Color {
	return pattern.c
}

func (pattern *StripePattern) ColorAt(point Tuple) Color {
	if int(math.Floor(point.x))%2 == 0 {
		return pattern.a.ColorAt(point)
	} else {
		return pattern.b.ColorAt(point)
	}
}

func (pattern *GradientPattern) ColorAt(point Tuple) Color {
	distance := pattern.b.ColorAt(point).Subtract(pattern.a.ColorAt(point))
	fraction := point.x - math.Floor(point.x)

	return pattern.a.ColorAt(point).Add(distance.MultiplyScalar(fraction))
}

func (pattern *RingPattern) ColorAt(point Tuple) Color {
	if int(math.Floor(math.Sqrt(math.Pow(point.x, 2)+math.Pow(point.z, 2))))%2 == 0 {
		return pattern.a.ColorAt(point)
	} else {
		return pattern.b.ColorAt(point)
	}
}

func (pattern *CheckersPattern) ColorAt(point Tuple) Color {
	if int(math.Abs(point.x)+math.Abs(point.y)+math.Abs(point.z))%2 == 0 {
		return pattern.a.ColorAt(point)
	} else {
		return pattern.b.ColorAt(point)
	}
}
