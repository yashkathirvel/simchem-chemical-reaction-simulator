package main

import (
	"fmt"
)

func SimulateSurface(initialS *Surface, numGens int, timeStep float64) []*Surface {
	timePoints := make([]*Surface, numGens)
	// set the initial Surface object as the first time point
	timePoints[0] = initialS
	// iterate through numGens generations and update the Surface object each time.
	for i := 1; i < numGens; i++ {
		timePoints[i] = timePoints[i-1].Update(timeStep, 0)
		fmt.Println("Generation: ", i)
	}
	return timePoints
}

// Surface method: Update()
// Updates the Surface object given a time s
func (s *Surface) Update(timeStep float64, rate float64) *Surface {
	// create a copy of the current Surface object
	newS := s.Copy()
	newS.Diffuse(timeStep)
	// update the position of each particle
	// for _, p := range newS.particles {
	// 	p = ZerothUpdatePosition(p,rate)
	// }
	return newS
}

// in zeroth order reactions, the reaction progresses at a rate that is
// independent of all chemical concentrations. this means products
// are formed spontaneously.

// zeroth update position takes a particle and the underlying rate constant
// updates position based simply on rate constant, with no relation to other particles
// in the system
// func ZerothUpdatePosition(p Particle, rateConstant float64) OrderedPair {
// 	var pos OrderedPair // initializes new position
// 	std := math.Sqrt(2 * time * rateConstant)

// 	if rateConstant > 1 {
// 		// updates position based on rate constant
// 		newParticle := p.CopyParticle
// 		dx := generatorX.NormFloat64() * std
// 		dy := generatorY.NormFloat64() * std
// 		newParticle.x += dx
// 		newParticle.y += dy
// 		return newParticle
// 	}

// 	dx := generatorX.NormFloat64() * std
// 	dy := generatorY.NormFloat64() * std
// 	pos.x += dx
// 	pos.y += dy

// 	return pos
// }

// // particle method: Copy
// // creates a deep copy of the particle object at hand
// func (s *Particle) Copy() *Particle {
// 	// create new address for newP
// 	var newP Particle

// 	newP.position = s.position
// 	newP.radius = s.radius
// 	newP.red = s.red
// 	newP.green = s.green
// 	newP.blue = s.blue

// 	return &newP
// }

// Surface method: Copy()
// Creates a deep copy of the Surface object.
func (s *Surface) Copy() *Surface {
	// create new address for newS
	var newS Surface
	newS.particles = make([]*Particle, 0)
	newS.width = s.width
	// iterate through the particles on the surface
	for _, particle := range s.particles {
		newParticle := &Particle{
			position: particle.position,
			species:  particle.species,
		}
		// append the new particle to the new surface
		newS.particles = append(newS.particles, newParticle)
	}
	return &newS
}
