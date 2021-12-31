package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	generateScene()
}

func generateScene() {
	floor := NewSphere()
	floor.transform = floor.transform.Scale(10, 0.01, 10)
	floor.material.color = Color{1, 0.9, 0.9}
	floor.material.specular = 0

	leftWall := NewSphere()
	leftWall.transform = leftWall.transform.
		Translate(0, 0, 5).
		RotateY(-math.Pi/4).
		RotateX(math.Pi/2).
		Scale(10, 0.01, 10)
	leftWall.material = floor.material

	rightWall := NewSphere()
	rightWall.transform = rightWall.transform.
		Translate(0, 0, 5).
		RotateY(math.Pi/4).
		RotateX(math.Pi/2).
		Scale(10, 0.01, 10)
	rightWall.material = floor.material

	middleSphere := NewSphere()
	middleSphere.transform = middleSphere.transform.Translate(-0.5, 1, 0.5)
	middleSphere.material.color = Color{0.1, 1, 0.5}
	middleSphere.material.diffuse = 0.7
	middleSphere.material.specular = 0.3

	rightSphere := NewSphere()
	rightSphere.transform = rightSphere.transform.
		Translate(1.5, 0.5, -0.5).
		Scale(0.5, 0.5, 0.5)
	rightSphere.material.color = Color{0.5, 1, 0.1}
	rightSphere.material.diffuse = 0.7
	rightSphere.material.specular = 0.3

	leftSphere := NewSphere()
	leftSphere.transform = leftSphere.transform.
		Translate(-1.5, 0.33, -0.75).
		Scale(0.33, 0.33, 0.33)
	leftSphere.material.color = Color{1, 0.8, 0.1}
	leftSphere.material.diffuse = 0.7
	leftSphere.material.specular = 0.3

	world := World{}
	world.light = &PointLight{NewPoint(-10, 10, -10), Color{1, 1, 1}}
	world.objects = make([]Sphere, 0)
	world.objects = append(world.objects, floor, leftWall, rightWall, middleSphere, rightSphere, leftSphere)

	camera := NewCamera(1000, 500, math.Pi/3)
	camera.transform = ViewTransform(NewPoint(0, 1.5, -5), NewPoint(0, 1, 0), NewVector(0, 1, 0))

	canvas := camera.Render(world)

	os.WriteFile("scene.ppm", []byte(canvas.ToPPM()), 0666)
}

func generateSphere() {
	rayOrigin := NewPoint(0, 0, -5)
	wallZ := 10
	wallSize := 7.0
	var canvasPixels uint = 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := NewCanvas(canvasPixels, canvasPixels)
	shape := NewSphere()
	shape.material.color = Color{1, 0.2, 1}

	lightPosition := NewPoint(-10, 10, -10)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

	for y := uint(0); y < canvas.height; y++ {
		// compute the world y coordinate (top = +half, bottom = -half)
		worldY := half - pixelSize*float64(y)

		for x := uint(0); x < canvas.width; x++ {
			// compute the world x coordinate (left = -half, right = half)
			worldX := -half + pixelSize*float64(x)

			position := NewPoint(worldX, worldY, float64(wallZ))

			r := Ray{rayOrigin, position.Subtract(rayOrigin).Normalize()}
			xs := shape.Intersects(r)

			hit, err := Hit(xs)
			if err == nil {
				point := r.Position(hit.t)
				normal := hit.object.NormalAt(point)
				eye := r.direction.Multiply(-1)
				color := Lighting(hit.object.material, light, point, eye, normal, false)
				canvas.WritePixel(x, y, color)
			}
		}
	}

	os.WriteFile("sphere.ppm", []byte(canvas.ToPPM()), 0666)
}

func clock() {
	var length uint = 20
	centerPos := float64(length / 2)
	radius := 3.0 / 8 * float64(length)
	canvas := NewCanvas(length, length)
	hour := NewPoint(0, 0, 1)
	rotation := NewIdentityMatrix().RotateY(2 * math.Pi / 12)

	for i := 0; i < 12; i++ {
		canvas.WritePixel(uint((hour.x*radius)+centerPos), uint((hour.z*radius)+centerPos), Color{1, 1, 1})
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

	var width uint = 900
	var height uint = 550
	canvas := NewCanvas(width, height)

	ticks := 0

	for proj.position.y > 0 {
		fmt.Println("projectile position = ", proj.position)
		canvas.WritePixel(uint(proj.position.x), height-uint(proj.position.y), Color{1, 0, 0})

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
