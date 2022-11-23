package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {
	// now evolve the universe: feel free to adjust the following parameters.
	numGens := 100
	timeStep := 5.00
	//Declaring A particle
	A := Species{
		diffusionRate: 1.0,
		radius:        5,
		red:           132,
		green:         83,
		blue:          60,
	}

	// initial Surface (for testing purposes)
	initialSurface := &Surface{
		particles: []*Particle{
			{
				x:       200,
				y:       200,
				species: &A,
			},
		},
		width: 400,
	}
	for i := 0; i < 10; i++ {
		p := Particle{
			x:       200,
			y:       200,
			species: &A,
		}
		initialSurface.particles = append(initialSurface.particles, &p)
	}
	start := time.Now()
	timePoints := SimulateSurface(initialSurface, numGens, timeStep)
	elapse := time.Since(start)

	fmt.Println("Simulation took", elapse, "s. Now drawing images.")
	canvasWidth := 1000
	frequency := 1
	scalingFactor := 3.0
	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "simchem")
	fmt.Println("GIF drawn.")
	fmt.Println("Exiting normally.")
}
