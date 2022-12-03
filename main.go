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
	diffusion_cons_A := 1.0
	diffusion_cons_B := 1.0

	// construct Species types
	A := &Species{
		name:          "A",
		diffusionRate: 1.0,
		radius:        1,
		red:           132,
		green:         83,
		blue:          60,
	}

	B := &Species{
		name:          "B",
		diffusionRate: 1.0,
		radius:        1,
		red:           255,
		green:         0,
		blue:          255,
	}

	// initial Surface (for testing purposes)
	initialSurface := &Surface{
		A_particles: []*Particle{
			{
				position: OrderedPair{200, 200},
				species:  A,
			},
		},
		B_particles: []*Particle{
			{
				position: OrderedPair{150, 150},
				species: B,
			},
		},
		width: 400,
	}

	for i := 0; i < 10; i++ {
		A_p := Particle{
			position: OrderedPair{200, 200},
			species:  A,
		}
		B_p := Particle{
			position: OrderedPair{150, 150},
			species:  B,
		}
		initialSurface.A_particles = append(initialSurface.A_particles, &A_p)
    initialSurface.B_particles = append(initialSurface.B_particles, &B_p)
		
	}

	// DRIVER CODE (DO NOT CHANGE!!)
	startTime := time.Now()
	timePoints := SimulateSurface(initialSurface, numGens, timeStep, diffusion_cons_A, diffusion_cons_B)
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
