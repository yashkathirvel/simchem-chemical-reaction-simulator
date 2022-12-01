package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {
	// evolution parameters
	numGens := 100
	timeStep := 5.00

	// construct Species types
	A := &Species{
		name:          "A",
		diffusionRate: 1.0,
		radius:        1,
		red:           132,
		green:         83,
		blue:          60,
	}

	// initial Surface (for testing purposes)
	initialSurface := &Surface{
		particles: []*Particle{
			{
				position: OrderedPair{200, 200},
				species:  A,
			},
		},
		width: 400,
	}
	for i := 0; i < 10; i++ {
		p := Particle{
			position: OrderedPair{200, 200},
			species:  A,
		}
		initialSurface.particles = append(initialSurface.particles, &p)
	}

	// DRIVER CODE (DO NOT CHANGE!!)
	startTime := time.Now()
	timePoints := SimulateSurface(initialSurface, numGens, timeStep)
	elapsedTime := time.Since(startTime)

	fmt.Println("Simulation took", elapsedTime, "s. Now drawing images.")
	canvasWidth := 1000
	frequency := 1
	scalingFactor := 3.0
	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "simchem")
	fmt.Println("GIF drawn.")
	fmt.Println("Exiting normally.")
}
