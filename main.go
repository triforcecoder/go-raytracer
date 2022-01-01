package main

import (
	"math"
	"os"
)

func main() {
	generateScene()
}

func generateScene() {
	floor := NewPlane()
	floor.material.color = Color{0, 0, 1}
	floor.material.ambient = 0.2

	leftSphere := NewSphere()
	leftSphere.transform = leftSphere.transform.
		Translate(-1.5, 0.33, -0.75).
		Scale(0.33, 0.33, 0.33)
	leftSphere.material.color = Color{1, 0, 0.1}
	leftSphere.material.diffuse = 0.7
	leftSphere.material.specular = 0.3

	middleSphere := NewSphere()
	middleSphere.transform = middleSphere.transform.Translate(-0.5, 1, 0.5)
	middleSphere.material.color = Color{0.1, 1, 0.1}
	middleSphere.material.diffuse = 0.7
	middleSphere.material.specular = 0.3

	rightSphere := NewSphere()
	rightSphere.transform = rightSphere.transform.
		Translate(1.5, 0.5, -0.5).
		Scale(0.5, 0.5, 0.5)
	rightSphere.material.color = Color{1, 0.2, 1}
	rightSphere.material.diffuse = 0.7
	rightSphere.material.specular = 0.3

	world := World{}
	world.light = &PointLight{NewPoint(-10, 10, -10), Color{1, 1, 1}}
	world.objects = make([]Shape, 0)
	world.objects = append(world.objects, floor, leftSphere, middleSphere, rightSphere)

	camera := NewCamera(2000, 1000, math.Pi/3)
	camera.transform = ViewTransform(NewPoint(0, 1.5, -5), NewPoint(0, 1, 0), NewVector(0, 1, 0))

	canvas := camera.Render(world)

	os.WriteFile("scene.ppm", []byte(canvas.ToPPM()), 0666)
}
