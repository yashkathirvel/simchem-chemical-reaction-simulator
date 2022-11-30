package main

import (
	"math"
	"math/rand"
	"time"
)

// calling BrownianMotion() to all particles in parallel
func (p *Particle) Diffuse(timeStep float64) {
	//allocate a new PRNG objec for every object
	sourceX := rand.NewSource(time.Now().UnixNano())
	generatorX := rand.New(sourceX)
	time.Sleep(time.Nanosecond) //To generate a different PRNG
	sourceY := rand.NewSource(time.Now().UnixNano())
	generatorY := rand.New(sourceY)
	p.BrownianMotion(generatorX, generatorY, timeStep)
	//runing too fast that seeds being the same?
}

// Diffuse function update a Particle's displacement after 1 time
func (p *Particle) BrownianMotion(generatorX, generatorY *(rand.Rand), timeStep float64) {
	std := math.Sqrt(2 * timeStep * p.species.diffusionRate)
	dx := generatorX.NormFloat64() * std
	dy := generatorY.NormFloat64() * std
	p.position.x += dx
	p.position.y += dy
	//probably need to handle off boundary senario
}
