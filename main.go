package main

import "fmt"

func main() {
	proj := Projectile{createPoint(0, 1, 0), createVector(1, 1, 0).normalize()}
	env := Environment{createVector(0, -0.1, 0), createVector(-0.01, 0, 0)}
	ticks := 0

	fmt.Println("projectile position = ", proj.position)

	for proj.position.y > 0 {
		proj = tick(env, proj)
		ticks++
		fmt.Println("projectile position = ", proj.position)
	}

	fmt.Println(ticks, " ticks to hit the ground")
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
