package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	canvas := NewCanvas(10, 20)

	assert.Equal(t, 10, canvas.width)
	assert.Equal(t, 20, canvas.height)
	assert.Equal(t, 10, len(canvas.pixel))
	assert.Equal(t, 20, len(canvas.pixel[0]))

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			assert.Equal(t, Color{}, canvas.pixel[i][j])
		}
	}
}

func TestWritePixel(t *testing.T) {
	canvas := NewCanvas(10, 20)
	red := Color{1, 0, 0}
	canvas.WritePixel(2, 3, red)

	assert.Equal(t, red, canvas.pixel[2][3])
}

func TestHeaderToPPM(t *testing.T) {
	canvas := NewCanvas(5, 3)
	ppm := canvas.ToPPM()

	result := strings.Split(ppm, "\n")

	assert.Equal(t, "P3", result[0])
	assert.Equal(t, "5 3", result[1])
	assert.Equal(t, "255", result[2])
}

func TestPixelDataToPPM(t *testing.T) {
	canvas := NewCanvas(5, 3)
	color1 := Color{1.5, 0, 0}
	color2 := Color{0, 0.5, 0}
	color3 := Color{-0.5, 0, 1}

	canvas.WritePixel(0, 0, color1)
	canvas.WritePixel(2, 1, color2)
	canvas.WritePixel(4, 2, color3)

	ppm := canvas.ToPPM()

	result := strings.Split(ppm, "\n")

	assert.Equal(t, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", result[3])
	assert.Equal(t, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0", result[4])
	assert.Equal(t, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", result[5])
}

func TestSplitLongLinesToPPM(t *testing.T) {
	canvas := NewCanvas(10, 2)

	for j := 0; j < canvas.height; j++ {
		for i := 0; i < canvas.width; i++ {
			canvas.WritePixel(i, j, Color{1, 0.8, 0.6})
		}
	}

	ppm := canvas.ToPPM()

	result := strings.Split(ppm, "\n")

	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", result[3])
	assert.Equal(t, "153 255 204 153 255 204 153 255 204 153 255 204 153", result[4])
	assert.Equal(t, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", result[5])
	assert.Equal(t, "153 255 204 153 255 204 153 255 204 153 255 204 153", result[6])
}

func TestEndNewlineToPPM(t *testing.T) {
	canvas := NewCanvas(5, 3)

	ppm := canvas.ToPPM()

	assert.True(t, strings.HasSuffix(ppm, "\n"))
}

func TestScaleFloat(t *testing.T) {
	assert.Equal(t, 0, scaleFloat(-1))
	assert.Equal(t, 0, scaleFloat(-0.5))
	assert.Equal(t, 0, scaleFloat(0))
	assert.Equal(t, 255, scaleFloat(1.6))
	assert.Equal(t, 255, scaleFloat(1))
	assert.Equal(t, 128, scaleFloat(0.5))
}
