package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func SimulateSurface(initialS *Surface, numGens int, timeStep, diffusion_cons_A, diffusion_cons_B, killRate, zerothRateConstant, bimolecularRateConstant float64) []*Surface {
	timePoints := make([]*Surface, numGens)
	// set the initial Surface object as the first time point
	timePoints[0] = initialS
	// iterate through numGens generations and update the Surface object each time.
	//dataA := make([]int, 0)
	//dataB := make([]int, 0)
	for i := 1; i < numGens; i++ {
		timePoints[i] = timePoints[i-1].Update(timeStep, bimolecularRateConstant, diffusion_cons_A, diffusion_cons_B, killRate, zerothRateConstant)
		a := len(timePoints[i].A_particles)
		//b := len(timePoints[i].B_particles)
		//fmt.Println(a, ",", b, ",")
		fmt.Println(a, ",", i, ",")
	}
	return timePoints
}

// Surface method: Update()
// Updates the Surface object given a time s
func (s *Surface) Update(timeStep, bimolecularRateConstant, diffusion_cons_A, diffusion_cons_B, killRate, zerothRateConstant float64) *Surface {
	// create a copy of the current Surface object
	newS := s.Copy()
	rand.Seed(time.Now().UnixNano())
	// iterate through the particles on the surface and diffuse them
	Astd := math.Sqrt(2 * timeStep * diffusion_cons_A)
	for _, particle := range newS.A_particles {
		// diffuse the particle
		particle.Diffuse(Astd)
	}
	Bstd := math.Sqrt(2 * timeStep * diffusion_cons_B)
	for _, particle := range newS.B_particles {
		// diffuse the particle
		particle.Diffuse(Bstd)

	}
	/**
	for _, particle := range newS.C_particles {
		// diffuse the particle
		particle.Diffuse(timeStep)
	}**/
	//newS.LoktaVolterraReaction(bimolecularRateConstant, diffusion_cons_A, diffusion_cons_B)
	//newS.AddAParticles(zerothRateConstant, timeStep)
	//newS.KillParticles(killRate, timeStep)
	newS.BimolecularReaction(bimolecularRateConstant, diffusion_cons_A, diffusion_cons_B)
	return newS
}
func (newS *Surface) LoktaVolterraReaction(rateConstant, diffusion_cons_A, diffusion_cons_B float64) {
	//kSi = 4πDσb.
	binding_radius := rateConstant / (4 * math.Pi * (diffusion_cons_A + diffusion_cons_B))
	B := &Species{
		name:          "B",
		diffusionRate: 100.0,
		radius:        3,
		red:           0,
		green:         0,
		blue:          255,
		mass:          1.0,
	}
	newB := make([]OrderedPair, 0)
	//range through a and compare it's distance with the B particles
	//if the distance between them is less than the binding radius, make C_particles
	for _, b_particle := range newS.B_particles {
		for _, a_particle := range newS.A_particles {
			particle_dist := Distance(a_particle.position, b_particle.position)
			if particle_dist < binding_radius {
				new_dist := Average_pos(a_particle.position, b_particle.position)
				newB = append(newB, new_dist)
				newS.DeleteParticleA(a_particle)
			}
		}
	}

	for _, b_position := range newB {
		B_p := Particle{
			position: b_position,
			species:  B, //pointer to a species defined in main
		}
		newS.B_particles = append(newS.B_particles, &B_p)
	}
}
func (newS *Surface) DeleteParticleA(a *Particle) {
	//range through surface to find the index of the particle
	for i, particle := range newS.A_particles {
		if particle == a {
			newS.A_particles = append(newS.A_particles[:i], newS.A_particles[i+1:]...)
		}
	}
}

// Particle method: Copy{}
// Creates a deep copy of the particle object.
func (s *Particle) Copy() *Particle {
	// create new address for newP
	var newP Particle
	// copy the position and species of the particle
	newP.position.x = s.position.x
	newP.position.y = s.position.y
	newP.species = s.species
	return &newP
}

// Surface method: Copy()
// Creates a deep copy of the Surface object.
func (s *Surface) Copy() *Surface {
	// create new address for newS
	var newS Surface
	newS.A_particles = make([]*Particle, 0)
	newS.B_particles = make([]*Particle, 0)
	newS.C_particles = make([]*Particle, 0)
	newS.width = s.width
	// iterate through the particles on the surface
	for _, particle := range s.A_particles {
		newParticle := particle.Copy()
		newS.A_particles = append(newS.A_particles, newParticle)
	}
	for _, particle := range s.B_particles {
		newParticle := particle.Copy()
		newS.B_particles = append(newS.B_particles, newParticle)
	}
	for _, particle := range s.C_particles {
		newParticle := particle.Copy()
		newS.C_particles = append(newS.C_particles, newParticle)
	}
	return &newS
}

// Particle method: SurfaceReaction(), this method takes into account the interaction of the particles with the surface
// the simulation is kept simple by defining boundaries, we simulate an inert permeable boundary
// this function takes the witdth of the surface and reflects particles back into the medium when they hit the surface
func (p *Particle) SurfaceReaction(width float64) {
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
func (newS *Surface) BimolecularReaction(rateConstant, diffusion_cons_A, diffusion_cons_B float64) {
	//kSi = 4πDσb.
	binding_radius := rateConstant / (4 * math.Pi * (diffusion_cons_A + diffusion_cons_B))

	C := &Species{
		name:          "C",
		diffusionRate: 1.0,
		radius:        1,
		red:           255,
		green:         255,
		blue:          0,
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
					species:  C, //pointer to a species defined in main
				}
				newS.DeleteParticles(a_particle, b_particle)
				newS.C_particles = append(newS.C_particles, &C_p)
			}
		}

	}
}

func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return (math.Sqrt(deltaX*deltaX + deltaY*deltaY))
}

func Average_pos(p1, p2 OrderedPair) OrderedPair {
	//calculates avergae of the two positions
	var dist OrderedPair
	deltaX := (p1.x + p2.x) / 2
	deltaY := (p1.y + p2.y) / 2

	dist.x = deltaX
	dist.y = deltaY

	return dist
}

func (newS *Surface) DeleteParticles(a, b *Particle) {
	//range through surface to find the index of the particle
	for i, particle := range newS.A_particles {
		if particle == a {
			newS.A_particles = append(newS.A_particles[:i], newS.A_particles[i+1:]...)
		}
	}
	for i, particle := range newS.B_particles {
		if particle == b {
			newS.B_particles = append(newS.B_particles[:i], newS.B_particles[i+1:]...)
		}
	}
}

// Surface method: Delete random B particles
func (s *Surface) DeleteRandomBParticle(i int) {
	//range through surface to find the index of the particle
	s.B_particles = append(s.B_particles[:i], s.B_particles[i+1:]...)

}

func (newS *Surface) KillParticles(killRate, timeStep float64) {
	// initialize global pseudo random generator
	rand.Seed(time.Now().UnixNano())
	prob := 1.0 - math.Exp(-killRate*timeStep)
	deathList := make([]int, 0)
	for i := range newS.B_particles {
		if rand.Float64() < prob {
			deathList = append(deathList, i)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(deathList)))
	for i := range deathList {
		newS.DeleteRandomBParticle(deathList[i])
	}
}

func (newS *Surface) AddAParticles(zerothRateConstant, timeStep float64) {
	// initialize global pseudo random generator
	rand.Seed(time.Now().UnixNano())
	prob := 1.0 - math.Exp(-zerothRateConstant*timeStep)
	//numParticles := int(timeStep * zerothRateConstant)
	A := &Species{
		name:          "A",
		diffusionRate: 100.0,
		radius:        3,
		red:           255,
		green:         0,
		blue:          0,
		mass:          1.0,
	}

	for i := 0; i < len(newS.A_particles); i++ {
		if rand.Float64() < prob {

			newParticle := Particle{
				position: OrderedPair{rand.Float64() * newS.width, rand.Float64() * newS.width},
				species:  A,
			}
			newS.A_particles = append(newS.A_particles, &newParticle)
		}
	}
}
