package geometry

import (
	. "go-raytracer/core"
	"math"
)

type Cube struct {
	origin        Tuple
	Transform     Matrix
	cachedInverse Matrix
	Material      Material
}

func NewCube() *Cube {
	return &Cube{NewPoint(0, 0, 0), NewIdentityMatrix(), nil, NewMaterial()}
}

func (cube *Cube) NormalAt(point Tuple) Tuple {
	objectPoint := cube.GetInverse().MultiplyTuple(point)
	var objectNormal Tuple

	maxC := math.Max(math.Max(math.Abs(objectPoint.X), math.Abs(objectPoint.Y)), math.Abs(objectPoint.Z))

	if maxC == math.Abs(objectPoint.X) {
		objectNormal = NewVector(objectPoint.X, 0, 0)
	} else if maxC == math.Abs(objectPoint.Y) {
		objectNormal = NewVector(0, objectPoint.Y, 0)
	} else {
		objectNormal = NewVector(0, 0, objectPoint.Z)
	}

	worldNormal := cube.GetInverse().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}

func (cube *Cube) Intersects(ray Ray) []Intersection {
	ray = ray.Transform(cube.GetInverse())

	xtMin, xtMax := checkAxis(ray.Origin.X, ray.Direction.X)
	ytMin, ytMax := checkAxis(ray.Origin.Y, ray.Direction.Y)
	ztMin, ztMax := checkAxis(ray.Origin.Z, ray.Direction.Z)
	tMin := math.Max(math.Max(xtMin, ytMin), ztMin)
	tMax := math.Min(math.Min(xtMax, ytMax), ztMax)

	if tMin > tMax {
		return []Intersection{}
	}

	return []Intersection{{tMin, cube}, {tMax, cube}}
}

func (cube *Cube) GetMaterial() Material {
	return cube.Material
}

func (cube *Cube) SetMaterial(material Material) {
	cube.Material = material
}

func (cube *Cube) GetTransform() Matrix {
	return cube.Transform
}

func (cube *Cube) GetInverse() Matrix {
	if cube.cachedInverse == nil {
		cube.cachedInverse = cube.Transform.Inverse()
	}

	return cube.cachedInverse
}

func checkAxis(origin, direction float64) (float64, float64) {
	var tmin, tmax float64

	tMinNumerator := -1 - origin
	tMaxNumerator := 1 - origin

	if math.Abs(direction) >= Epsilon {
		tmin = tMinNumerator / direction
		tmax = tMaxNumerator / direction
	} else {
		tmin = math.Inf(int(tMinNumerator))
		tmax = math.Inf(int(tMaxNumerator))
	}

	if tmin > tmax {
		return tmax, tmin
	} else {
		return tmin, tmax
	}
}
