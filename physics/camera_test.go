package physics

import (
	. "go-raytracer/core"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EqualTuple(t *testing.T, expected Tuple, actual Tuple) {
	if expected.Equals(actual) {
		assert.True(t, true)
	} else {
		assert.Equal(t, expected, actual)
	}
}

func EqualColor(t *testing.T, expected Color, actual Color) {
	if expected.Equals(actual) {
		assert.True(t, true)
	} else {
		assert.Equal(t, expected, actual)
	}
}

func TestNewCamera(t *testing.T) {
	var hsize uint = 160
	var vsize uint = 120
	fieldOfView := math.Pi / 2

	camera := NewCamera(hsize, vsize, fieldOfView)

	assert.Equal(t, uint(160), camera.hsize)
	assert.Equal(t, uint(120), camera.vsize)
	assert.Equal(t, math.Pi/2, camera.fieldOfView)
	assert.Equal(t, NewIdentityMatrix(), camera.Transform)
}

func TestPixelSizeHorizontalCanvas(t *testing.T) {
	var hsize uint = 200
	var vsize uint = 125
	fieldOfView := math.Pi / 2

	camera := NewCamera(hsize, vsize, fieldOfView)

	assert.Equal(t, 0.01, camera.pixelSize)
}

func TestPixelSizeVerticalCanvas(t *testing.T) {
	var hsize uint = 125
	var vsize uint = 200
	fieldOfView := math.Pi / 2

	camera := NewCamera(hsize, vsize, fieldOfView)

	assert.Equal(t, 0.01, camera.pixelSize)
}

func TestConstructRayCenterOfCanvas(t *testing.T) {
	var hsize uint = 201
	var vsize uint = 101
	fieldOfView := math.Pi / 2
	camera := NewCamera(hsize, vsize, fieldOfView)

	ray := camera.RayForPixel(100, 50)

	assert.Equal(t, NewPoint(0, 0, 0), ray.Origin)
	assert.Equal(t, NewVector(0, 0, -1), ray.Direction)
}

func TestConstructRayCornerOfCanvas(t *testing.T) {
	var hsize uint = 201
	var vsize uint = 101
	fieldOfView := math.Pi / 2
	camera := NewCamera(hsize, vsize, fieldOfView)

	ray := camera.RayForPixel(0, 0)

	assert.Equal(t, NewPoint(0, 0, 0), ray.Origin)
	EqualTuple(t, NewVector(0.66519, 0.33259, -0.66851), ray.Direction)
}

func TestConstructRayCameraTransformed(t *testing.T) {
	var hsize uint = 201
	var vsize uint = 101
	fieldOfView := math.Pi / 2
	camera := NewCamera(hsize, vsize, fieldOfView)
	camera.Transform = camera.Transform.RotateY(math.Pi/4).Translate(0, -2, 5)

	ray := camera.RayForPixel(100, 50)

	assert.Equal(t, NewPoint(0, 2, -5), ray.Origin)
	EqualTuple(t, NewVector(math.Sqrt2/2, 0, -math.Sqrt2/2), ray.Direction)
}

func TestRenderWorldWithCamera(t *testing.T) {
	world := DefaultWorld()
	camera := NewCamera(11, 11, math.Pi/2)
	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)
	camera.Transform = ViewTransform(from, to, up)

	image := camera.Render(world)
	result := image.Pixel[5][5]

	EqualColor(t, NewColor(0.38066, 0.47583, 0.2855), result)
}
