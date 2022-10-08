package main

import "math"

type Camera struct {
	hsize         uint
	vsize         uint
	fieldOfView   float64
	transform     Matrix
	cachedInverse Matrix
	pixelSize     float64
	halfWidth     float64
	halfHeight    float64
}

func NewCamera(hsize uint, vsize uint, fieldOfView float64) Camera {
	var halfWidth float64
	var halfHeight float64

	halfView := math.Tan(fieldOfView / 2)
	aspect := float64(hsize) / float64(vsize)

	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}

	pixelSize := halfWidth * 2 / float64(hsize)

	return Camera{hsize, vsize, fieldOfView, NewIdentityMatrix(),
		nil, pixelSize, halfWidth, halfHeight}
}

func (camera *Camera) RayForPixel(px uint, py uint) Ray {
	if camera.cachedInverse == nil {
		camera.cachedInverse = camera.transform.Inverse()
	}

	xoffset := (float64(px) + 0.5) * camera.pixelSize
	yoffset := (float64(py) + 0.5) * camera.pixelSize

	worldX := camera.halfWidth - xoffset
	worldY := camera.halfHeight - yoffset

	pixel := camera.cachedInverse.MultiplyTuple(NewPoint(worldX, worldY, -1))
	origin := camera.cachedInverse.MultiplyTuple(NewPoint(0, 0, 0))
	direction := pixel.Subtract(origin).Normalize()

	return Ray{origin, direction}
}

func (camera *Camera) Render(world World) Canvas {
	image := NewCanvas(camera.hsize, camera.vsize)

	for y := uint(0); y < camera.vsize; y++ {
		for x := uint(0); x < camera.hsize; x++ {
			ray := camera.RayForPixel(x, y)
			color := world.ColorAt(ray, 4)
			image.WritePixel(x, y, color)
		}
	}

	return image
}
