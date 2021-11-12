package main

import (
	"fmt"
	"os"
)

func main() {
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
	os.WriteFile("canvas.ppm", []byte(canvas.ToPPM()), 0666)
}

type Projectile struct {
	position Tuple
	velocity Tuple
}

type Environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add((env.wind))
	return Projectile{position, velocity}
}
