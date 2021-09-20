package main

import (
	"fmt"
	"os"
)

func main() {
	start := createPoint(0, 1, 0)
	velocity := createVector(1, 1.8, 0).normalize().multiply(11.25)
	proj := Projectile{start, velocity}

	gravity := createVector(0, -0.1, 0)
	wind := createVector(-0.01, 0, 0)
	env := Environment{gravity, wind}

	width := 900
	height := 550
	canvas := createCanvas(width, height)

	ticks := 0

	for proj.position.y > 0 {
		fmt.Println("projectile position = ", proj.position)
		canvas.writePixel(int(proj.position.x), height-int(proj.position.y), Color{1, 0, 0})

		proj = tick(env, proj)
		ticks++
	}

	fmt.Println(ticks, " ticks to hit the ground")
	os.WriteFile("canvas.ppm", []byte(canvas.toPPM()), 0666)
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
	position := proj.position.add(proj.velocity)
	velocity := proj.velocity.add(env.gravity).add((env.wind))
	return Projectile{position, velocity}
}
