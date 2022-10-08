package main

import (
	"flag"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	generateScene()
}

func generateScene() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	floor := NewPlane()
	floor.transform = floor.transform.Translate(1, 1, -2)
	floor.material.color = blue
	floor.material.ambient = 0.2
	floor.material.pattern = NewCheckersPattern(
		NewSolidPattern(black), NewGradientPattern(
			NewSolidPattern(green),
			NewSolidPattern(blue)))
	floor.material.reflective = 0.5
	floor.material.transparency = 0.2
	floor.material.refractiveIndex = 4.5

	middleSphere := NewSphere()
	middleSphere.transform = middleSphere.transform.
		Translate(-0.5, 1.5, 0.5).
		Scale(0.8, 0.8, 0.8)
	middleSphere.material.color = Color{0.1, 1, 0.1}
	middleSphere.material.diffuse = 0.7
	middleSphere.material.specular = 0.3
	middleSphere.material.ambient = 0.3
	middleSphere.material.transparency = 0.5
	middleSphere.material.refractiveIndex = 0
	middleSpherePattern := NewGradientPattern(
		NewSolidPattern(red), NewSolidPattern(blue))
	middleSpherePattern.transform = middleSphere.transform.
		RotateX(math.Pi / 4).
		RotateY(math.Pi / 4).
		RotateZ(math.Pi)
	middleSphere.material.pattern = middleSpherePattern

	rightSphere := NewSphere()
	rightSphere.transform = rightSphere.transform.
		Translate(1.5, 1.5, -0.5).
		Scale(0.5, 0.5, 0.5)
	rightSphere.material.color = Color{1, 0.2, 1}
	rightSphere.material.diffuse = 0.7
	rightSphere.material.specular = 0.3
	rightSphere.material.reflective = 0.5
	rightSphere.material.transparency = 0.5
	rightSphere.material.refractiveIndex = 1.5
	rightSpherePattern := NewGradientPattern(
		NewSolidPattern(red), NewSolidPattern(blue))
	rightSpherePattern.transform = rightSphere.transform.
		RotateX(math.Pi / 2).
		RotateZ(math.Pi / 4)
	rightSphere.material.pattern = rightSpherePattern

	world := World{}
	world.light = &PointLight{NewPoint(-10, 10, -10), white}
	world.objects = make([]Shape, 0)
	world.objects = append(world.objects, floor, middleSphere, rightSphere)

	camera := NewCamera(2000, 1000, math.Pi/3)
	camera.transform = ViewTransform(NewPoint(0, 1.5, -5), NewPoint(0, 1, 0), NewVector(0, 1, 0))

	canvas := camera.Render(world)

	os.WriteFile("render/scene.ppm", []byte(canvas.ToPPM()), 0666)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
