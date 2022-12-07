package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {
	// evolution parameters
	timeStep := 0.01
	numGens := 1000

	// construct Species types
	A := &Species{
		name:          "A",
		diffusionRate: 500.0,
		radius:        3,
		red:           255,
		green:         0,
		blue:          0,
		mass:          1.0,
	}

	B := &Species{
		name:          "B",
		diffusionRate: 500.0,
		radius:        3,
		red:           0,
		green:         0,
		blue:          255,
		mass:          1.0,
	}

	// initial Surface (for testing purposes)
	initialSurface := &Surface{
		A_particles: []*Particle{
			{
				position: OrderedPair{0.0, 0.0},
				species:  A,
			},
		},
		B_particles: []*Particle{
			{
				position: OrderedPair{200.0, 200.0},
				species:  B,
			},
		},
		width: 400,
	}

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			A_p := Particle{
				//position: OrderedPair{200, 200},
				species: A,
			}
			A_p.position.x = float64(i * 20.0)
			A_p.position.y = float64(j * 20.0)
			B_p := Particle{
				//position: OrderedPair{150, 150},
				species: B,
			}
			B_p.position.x = float64(i*20.0 + 10.0)
			B_p.position.y = float64(j*20.0 + 10.0)
			initialSurface.A_particles = append(initialSurface.A_particles, &A_p)
			initialSurface.B_particles = append(initialSurface.B_particles, &B_p)
		}
	}

	killRate := 2.0
	zerothRateConstant := 2.0
	bimolecularRateConstant := 40000.0
	// DRIVER CODE (DO NOT CHANGE!!)
	startTime := time.Now()
	timePoints := SimulateSurface(initialSurface, numGens, timeStep, A.diffusionRate, B.diffusionRate, killRate, zerothRateConstant, bimolecularRateConstant)
	//timePoints := SimulateSurfaceCollision(initialSurface, numGens, timeStep, rateConstant0, rateConstant1, A.diffusionRate, B.diffusionRate)
	elapsedTime := time.Since(startTime)

	fmt.Println("Simulation took", elapsedTime, "s. Now drawing images.")
	canvasWidth := 1000
	frequency := 10
	scalingFactor := 1.0
	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "simchem2cd")
	fmt.Println("GIF drawn.")
	fmt.Println("Exiting normally.")
}
