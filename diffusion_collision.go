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
					Collision(p1, p2)
				}
			}
		}
	}

}

func Collision(p1, p2 *Particle) {

	r := Subtract(p1.position, p2.position)
	u1Parallel := Projection(p1.velocity, r)
	u1Perpendicular := Subtract(p1.velocity, u1Parallel)

	u2Parallel := Projection(p2.velocity, r)
	u2Perpendicular := Subtract(p2.velocity, u2Parallel)
	u1ParallelTranslated := Subtract(u1Parallel, u2Parallel)

	v1ParallelTranslated := Scale((p1.species.mass-p2.species.mass)/(p1.species.mass+p2.species.mass), u1ParallelTranslated)
	v1Parallel := Add(v1ParallelTranslated, u2Parallel)
	v2Parallel := Add(Subtract(u1Parallel, u2Parallel), v1Parallel)
	v1Perpendicular := u1Perpendicular
	v2Perpendicular := u2Perpendicular
	new_v1 := Add(v1Parallel, v1Perpendicular)
	new_v2 := Add(v2Parallel, v2Perpendicular)
	p1.velocity = new_v1
	p2.velocity = new_v2
}

// vector scaling
func Scale(n float64, v OrderedPair) OrderedPair {
	v.x *= n
	v.y *= n
	return v
}

// vector subtraction
func Subtract(p1, p2 OrderedPair) OrderedPair {
	p1.x -= p2.x
	p1.y -= p2.y
	return p1
}

// vector subtraction. shouldnt do it to ptr
func Add(p1, p2 OrderedPair) OrderedPair {
	p1.x += p2.x
	p1.y += p2.y
	return p1
}

// dot product of two vectors
func Dot(v1, v2 OrderedPair) float64 {
	return v1.x*v2.x + v1.y*v2.y
}

// return v1 projected on v2
func Projection(v1, v2 OrderedPair) OrderedPair {
	if Dot(v2, v2) != 0 {
		v1Magnitude := Dot(v1, v2) / Dot(v2, v2)

		v2.x *= v1Magnitude
		v2.y *= v1Magnitude
		return v2
	} else {
		v := OrderedPair{
			x: 0,
			y: 0,
		}
		return v
	}
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
