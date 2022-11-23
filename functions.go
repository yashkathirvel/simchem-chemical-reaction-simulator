package main

import "fmt"

func SimulateSurface(initialS *Surface, numGens int, timeStep float64) []*Surface {
	timePoints := make([]*Surface, numGens)
	// set the initial Surface object as the first time point
	timePoints[0] = initialS
	// iterate through numGens generations and update the Surface object each time.
	for i := 1; i < numGens; i++ {
		timePoints[i] = timePoints[i-1].Update(timeStep)
		fmt.Println("Generation: ", i)
	}
	return timePoints
}

// Surface method: Update()
// Updates the Surface object given a time step.
func (s *Surface) Update(timeStep float64) *Surface {
	// create a copy of the current Surface object
	newS := s.Copy()
	newS.Diffuse(timeStep)
	// update the position of each particle
	// for _, p := range newS.particles {
	// 	p.UpdatePosition(time)
	// }
	return newS
}

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
			x:       particle.x,
			y:       particle.y,
			species: particle.species,
		}
		// append the new particle to the new surface
		newS.particles = append(newS.particles, newParticle)
	}
	return &newS
}
