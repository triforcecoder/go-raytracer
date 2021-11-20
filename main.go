package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	redSphere()
}

func redSphere() {
	rayOrigin := NewPoint(0, 0, -5)
	wallZ := 10
	wallSize := 7.0
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := NewCanvas(canvasPixels, canvasPixels)
	red := Color{1, 0, 0}
	shape := NewSphere()

	for y := 0; y < canvas.height; y++ {
		// compute the world y coordinate (top = +half, bottom = -half)
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvas.width; x++ {
			// compute the world x coordinate (left = -half, right = half)
			worldX := -half + pixelSize*float64(x)

			position := NewPoint(worldX, worldY, float64(wallZ))

			r := Ray{rayOrigin, position.Subtract(rayOrigin).Normalize()}
			xs := shape.Intersects(r)

			if Hit(xs) != nil {
				canvas.WritePixel(x, y, red)
			}
		}
	}

	os.WriteFile("red-sphere.ppm", []byte(canvas.ToPPM()), 0666)
}

func clock() {
	length := 20
	centerPos := float64(length / 2)
	radius := 3.0 / 8 * float64(length)
	canvas := NewCanvas(length, length)
	hour := NewPoint(0, 0, 1)
	rotation := NewIdentityMatrix().RotateY(2 * math.Pi / 12)

	for i := 0; i < 12; i++ {
		canvas.WritePixel(int((hour.x*radius)+centerPos), int((hour.z*radius)+centerPos), Color{1, 1, 1})
		hour = rotation.MultiplyTuple(hour)
	}

	os.WriteFile("clock.ppm", []byte(canvas.ToPPM()), 0666)
}

type Projectile struct {
	position Tuple
	velocity Tuple
}

type Environment struct {
	gravity Tuple
	wind    Tuple
}

func simulatedProjectile() {
	start := NewPoint(0, 1, 0)
	velocity := NewVector(1, 1.8, 0).Normalize().Multiply(11.25)
	proj := Projectile{start, velocity}

	gravity := NewVector(0, -0.1, 0)
	wind := NewVector(-0.01, 0, 0)
	env := Environment{gravity, wind}

	width := 900
	height := 550
	canvas := NewCanvas(width, height)

	ticks := 0

	for proj.position.y > 0 {
		fmt.Println("projectile position = ", proj.position)
		canvas.WritePixel(int(proj.position.x), height-int(proj.position.y), Color{1, 0, 0})

		proj = tick(env, proj)
		ticks++
	}

	fmt.Println(ticks, " ticks to hit the ground")
	os.WriteFile("projectile.ppm", []byte(canvas.ToPPM()), 0666)
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add((env.wind))
	return Projectile{position, velocity}
}
