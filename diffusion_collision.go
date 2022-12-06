package main

import (
	"math"
	"math/rand"
)

// Top level function of collision-driven diffusion.
func (s *Surface) DiffuseCollision(timeStep float64) {
	liveList := append(s.A_particles, s.B_particles...)
	liveList = append(liveList, s.C_particles...)
	for _, p := range liveList {
		p.UpdatePosition(timeStep)
		p.SurfaceReaction(s.width)
	}

	for _, p1 := range liveList {
		for _, p2 := range liveList {
			if p2 != p1 {
				d := Distance(p1.position, p2.position)
				if d <= (p1.species.radius + p2.species.radius) {
					//Collision(p1, p2)
				}
			}
		}
	}

}

func Collision(p1, p2 *Particle) {

}

// seting inital velocity for all particles
func (p *Particle) SetInitialVelocity(timeStep float64) {
	//allocate a new PRNG objec for every object
	//source := rand.NewSource(time.Now().UnixNano())
	//generator := rand.New(source)
	//time.Sleep(time.Nanosecond) //To generate a different PRNG
	std := math.Sqrt(2 * timeStep * p.species.diffusionRate)
	dx := rand.NormFloat64() * std
	dy := rand.NormFloat64() * std
	p.velocity.x += dx
	p.velocity.y += dy
	//runing too fast that seeds being the same?
}

func (p *Particle) UpdatePosition(timeStep float64) {
	p.position.x += p.velocity.x * timeStep
	p.position.y += p.velocity.y * timeStep
}
