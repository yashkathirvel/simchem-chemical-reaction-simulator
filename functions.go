package main

import (
	"math"
	"math/rand"
	"time"
	"fmt"
)

func SimulateSurface(initialS *Surface, numGens int, time float64) []*Surface {
	timePoints := make([]*Surface, numGens)
	// set the initial Surface object as the first time point
	timePoints[0] = initialS
	// iterate through numGens generations and update the Surface object each time.
	for i := 1; i < numGens; i++ {
		timePoints[i] = timePoints[i-1].Update(time)
		fmt.Println("Generation: ", i)
	}
	return timePoints
}

// Surface method: Update()
// Updates the Surface object given a time step.
func (s *Surface) Update(time float64) *Surface {
	// create a copy of the current Surface object
	newS := s.Copy()
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
			position: OrderedPair{particle.position.x, particle.position.y},
			radius:   particle.radius,
			red:      particle.red,
			green:    particle.green,
			blue:     particle.blue,
		}
		// append the new particle to the new surface
		newS.particles = append(newS.particles, newParticle)
	}
	return &newS
}

// calling BrownianMotion() to all particles in parallel
func (s *Surface) Diffuse(timeStep float64) {
​
	for _, p := range s.particles {
		//allocate a new PRNG object for every object
		sourceX := rand.NewSource(time.Now().UnixNano())
		generatorX := rand.New(sourceX)
		time.Sleep(time.Nanosecond) //To generate a different PRNG
		sourceY := rand.NewSource(time.Now().UnixNano())
		generatorY := rand.New(sourceY)
		p.BrownianMotion(generatorX, generatorY, timeStep)
		//runing too fast that seeds being the same?
	}
}
​
// Diffuse function update a Particle's displacement after 1 time
func (p *Particle) BrownianMotion(generatorX, generatorY *(rand.Rand), timeStep float64) {
	std := math.Sqrt(2 * timeStep * p.species.diffusionRate)
	dx := generatorX.NormFloat64() * std
	dy := generatorY.NormFloat64() * std
	p.x += dx
	p.y += dy
	//probably need to handle off boundary senario
}
