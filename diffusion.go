package main

import (
	"math/rand"
	"time"
)

/**
func (p *Particle) Diffuse(std float64) {
	//allocate a new PRNG objec for every object
	dx := rand.NormFloat64() * std
	dy := rand.NormFloat64() * std
	p.position.x += dx
	p.position.y += dy
	//runing too fast that seeds being the same?
}
**/

// calling BrownianMotion() to all particles in parallel
func (p *Particle) Diffuse(std float64) {
	//allocate a new PRNG objec for every object
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	time.Sleep(time.Nanosecond) //To generate a different PRNG

	p.BrownianMotion(generator, std)
	p.SurfaceReaction(400.0)
	//runing too fast that seeds being the same?
}

// Diffuse function update a Particle's displacement after 1 time
func (p *Particle) BrownianMotion(generator *(rand.Rand), std float64) {
	//std := math.Sqrt(2 * timeStep * p.species.diffusionRate)
	dx := generator.NormFloat64() * std
	dy := generator.NormFloat64() * std
	p.position.x += dx
	p.position.y += dy
	//probably need to handle off boundary senario
}
