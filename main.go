package main

import (
	"flag"
	. "go-raytracer/core"
	. "go-raytracer/geometry"
	. "go-raytracer/physics"
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
	floor.Transform = floor.Transform.Translate(1, 1, -2)
	floor.Material.Color = Blue
	floor.Material.Ambient = 0.2
	floor.Material.Pattern = NewCheckersPattern(
		NewSolidPattern(Black), NewGradientPattern(
			NewSolidPattern(Green),
			NewSolidPattern(Blue)))
	floor.Material.Reflective = 0.5
	floor.Material.Transparency = 0.2
	floor.Material.RefractiveIndex = 4.5

	middleSphere := NewSphere()
	middleSphere.Transform = middleSphere.Transform.
		Translate(-0.5, 1.5, 0.5).
		Scale(0.8, 0.8, 0.8)
	middleSphere.Material.Color = NewColor(0.1, 1, 0.1)
	middleSphere.Material.Diffuse = 0.7
	middleSphere.Material.Specular = 0.3
	middleSphere.Material.Ambient = 0.3
	middleSphere.Material.Transparency = 0.5
	middleSphere.Material.RefractiveIndex = 0
	middleSpherePattern := NewGradientPattern(
		NewSolidPattern(Red), NewSolidPattern(Blue))
	middleSpherePattern.Transform = middleSpherePattern.Transform.
		RotateX(math.Pi / 4).
		RotateY(math.Pi / 4).
		RotateZ(math.Pi)
	middleSphere.Material.Pattern = middleSpherePattern

	rightSphere := NewSphere()
	rightSphere.Transform = rightSphere.Transform.
		Translate(1.5, 1.5, -0.5).
		Scale(0.5, 0.5, 0.5)
	rightSphere.Material.Color = NewColor(1, 0.2, 1)
	rightSphere.Material.Diffuse = 0.7
	rightSphere.Material.Specular = 0.3
	rightSphere.Material.Reflective = 0.5
	rightSphere.Material.Transparency = 0.5
	rightSphere.Material.RefractiveIndex = 1.5
	rightSpherePattern := NewGradientPattern(
		NewSolidPattern(Red), NewSolidPattern(Blue))
	rightSpherePattern.Transform = rightSphere.Transform.
		RotateX(math.Pi / 2).
		RotateZ(math.Pi / 4)
	rightSphere.Material.Pattern = rightSpherePattern

	world := World{}
	world.Light = NewPointLight(NewPoint(-10, 10, -10), White)
	world.Objects = make([]Shape, 0)
	world.Objects = append(world.Objects, floor, middleSphere, rightSphere)

	camera := NewCamera(2000, 1000, math.Pi/3)
	camera.Transform = ViewTransform(NewPoint(0, 1.5, -5), NewPoint(0, 1, 0), NewVector(0, 1, 0))

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
