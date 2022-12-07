package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func SimulateSurfaceCollision(initialS *Surface, numGens int, timeStep, rateConstant0, rateConstant2, diffusion_cons_A, diffusion_cons_B float64) []*Surface {
	timePoints := make([]*Surface, numGens)
	// set the initial Surface object as the first time point
	timePoints[0] = initialS
	initialS.SetInitialVelocity(timeStep)
	fmt.Print(initialS.A_particles[50].velocity)
	// iterate through numGens generations and update the Surface object each time.
	for i := 1; i < numGens; i++ {
		timePoints[i] = timePoints[i-1].UpdateCollision(timeStep, rateConstant0, rateConstant2, diffusion_cons_A, diffusion_cons_B)
		//fmt.Println("Generation: ", i)
	}
	return timePoints
}

func (s *Surface) SetInitialVelocity(timeStep float64) {
	rand.Seed(time.Now().UnixNano())
	liveList := append(s.A_particles, s.B_particles...)
	liveList = append(liveList, s.C_particles...)
	for _, p := range liveList {
		p.SetInitialVelocity(timeStep)
	}

}

// Surface method: Update()
// Updates the Surface object given a time s
func (s *Surface) UpdateCollision(timeStep, rateConstant0, rateConstant2, diffusion_cons_A, diffusion_cons_B float64) *Surface {
	// create a copy of the current Surface object
	newS := s.CopyCollision()

	// iterate through the particles on the surface
	newS.DiffuseCollision(timeStep)
	fmt.Println(newS.A_particles[3].velocity)
	//newS.BimolecularReaction(rateConstant2, diffusion_cons_A, diffusion_cons_B)

	//fmt.Println("survival of A: ", len(newS.A_particles))
	return newS
}

// Particle method: Copy{}
// Creates a deep copy of the particle object.
func (s *Particle) CopyCollision() *Particle {
	// create new address for newP
	var newP Particle
	// copy the position and species of the particle
	newP.position.x = s.position.x
	newP.position.y = s.position.y
	newP.velocity.x = s.velocity.x
	newP.velocity.y = s.velocity.y
	newP.species = s.species
	return &newP
}

// Surface method: Copy()
// Creates a deep copy of the Surface object.
func (s *Surface) CopyCollision() *Surface {
	// create new address for newS
	var newS Surface
	newS.A_particles = make([]*Particle, 0)
	newS.B_particles = make([]*Particle, 0)
	newS.C_particles = make([]*Particle, 0)
	newS.width = s.width
	// iterate through the particles on the surface
	for _, particle := range s.A_particles {
		newParticle := &Particle{
			position: particle.position,
			velocity: particle.velocity,
			species:  particle.species,
		}
		newS.A_particles = append(newS.A_particles, newParticle)
	}
	for _, particle := range s.B_particles {
		newParticle := &Particle{
			position: particle.position,
			velocity: particle.velocity,
			species:  particle.species,
		}
		newS.B_particles = append(newS.B_particles, newParticle)
	}
	for _, particle := range s.C_particles {
		newParticle := &Particle{
			position: particle.position,
			velocity: particle.velocity,
			species:  particle.species,
		}
		newS.C_particles = append(newS.C_particles, newParticle)
	}
	// append the new particle to the new surface

	return &newS
}

// Particle method: SurfaceReaction(), this method takes into account the interaction of the particles with the surface
// the simulation is kept simple by defining boundaries, we simulate an inert permeable boundary
// this function takes the witdth of the surface and reflects particles back into the medium when they hit the surface
func (p *Particle) SurfaceReactionCollision(width float64) {
	if p.position.x > width {
		p.position.x = p.position.x - (p.position.x - width)
	} else {
		if p.position.x < 0 {
			p.position.x = width - (p.position.x * (width / p.position.x))
		}
	}

	if p.position.y > width {
		p.position.y = p.position.y - (p.position.y - width)
	} else {
		if p.position.y < 0 {
			p.position.y = width - (p.position.y * (width / p.position.y))
		}
	}
}

// this function simulates the bimolecular reaction
// input: takes the rate constant of the reaction, calculates a binding radius from it which determines how far
// two species need to be from each other to initiate collision and consequently a chemical reaction
func (newS *Surface) BimolecularReactionCollision(rateConstant, diffusion_cons_A, diffusion_cons_B float64) {
	//kSi = 4πDσb.
	binding_radius := rateConstant / (4 * math.Pi * (diffusion_cons_A + diffusion_cons_B))
	//fmt.Println(binding_radius)
	C := &Species{
		name:          "C",
		diffusionRate: 1.0,
		radius:        1,
		red:           0,
		green:         255,
		blue:          0,
		mass:          1.0,
	}
	//range through a and compare it's distance with the B particles
	//if the distance between them is less than the binding radius, make C_particles
	for _, a_particle := range newS.A_particles {
		for _, b_particle := range newS.B_particles {
			particle_dist := Distance(a_particle.position, b_particle.position)

			if particle_dist < binding_radius {
				new_dist := Average_pos(a_particle.position, b_particle.position)
				C_p := Particle{
					position: new_dist,
					velocity: b_particle.velocity,
					species:  C, //pointer to a species defined in main
				}
				newS.DeleteParticles(a_particle, b_particle)
				newS.C_particles = append(newS.C_particles, &C_p)

			}
		}
	}

}