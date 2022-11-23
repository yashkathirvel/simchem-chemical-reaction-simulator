package main

import (
	"math"
	"math/rand"
	"runtime"
	"time"
)

// calling BrownianMotion() to all particles in parallel
func (s *Surface) Diffuse(timeStep float64) {
	//cr: Prof.Phillp from now
	numParticles := len(s.particles)
	numProcs := runtime.NumCPU()
	finished := make(chan bool, numProcs)

	//split the work over numProcs processes.
	for i := 0; i < numProcs; i++ {
		//each processor gets approx. numParticles/numProcs particles
		startIndex := i * numParticles / numProcs
		var endIndex int
		if i < numProcs-1 {
			endIndex = (i + 1) * numParticles / numProcs
		} else {
			endIndex = numParticles
		}
		//don't want a race condition where all processes share
		//a single PRNG object.
		sourceX := rand.NewSource(time.Now().UnixNano())
		generatorX := rand.New(sourceX) // creates new PRNG object
		time.Sleep(time.Nanosecond)
		sourceY := rand.NewSource(time.Now().UnixNano())
		generatorY := rand.New(sourceY)
		go DiffuseOneProc(s.particles[startIndex:endIndex], generatorX, generatorY, finished, timeStep)
	}

	// we need to receive a message from all our channels that they're finished
	for i := 0; i < numProcs; i++ {
		<-finished
	}

}
func DiffuseOneProc(particles []*Particle, generatorX, generatorY *(rand.Rand), finished chan bool, timeStep float64) {
	//source := rand.NewSource(time.Now().UnixNano())
	//generator := rand.New(source)
	for _, p := range particles {
		p.BrownianMotion(generatorX, generatorY, timeStep)
	}
	//function is done. Indicate this by placing true (or false) into channel
	finished <- true
}

// Diffuse function update a Particle's displacement after 1 timeStep
func (p *Particle) BrownianMotion(generatorX, generatorY *(rand.Rand), timeStep float64) {
	std := math.Sqrt(2 * timeStep * p.species.diffusionRate)
	dx := generatorX.NormFloat64() * std
	dy := generatorY.NormFloat64() * std
	p.x += dx
	p.y += dy
	//probably need to handle off boundary senario
}
