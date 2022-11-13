package main

import (
	"fmt"
	"gifhelper"
)

func main() {
	// now evolve the universe: feel free to adjust the following parameters.
	numGens := 1000
	time := 5.00

	// initial Surface (for testing purposes)
	initialSurface := &Surface{
		particles: []*Particle{
			{
				position: OrderedPair{200000, 200000},
				radius:   250,
				red:      132,
				green:    83,
				blue:     60,
			},
		},
		width: 400000,
	}

	timePoints := SimulateSurface(initialSurface, numGens, time)

	fmt.Println("Simulation run. Now drawing images.")
	canvasWidth := 1000
	frequency := 1000
	scalingFactor := 15.0
	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "simchem")
	fmt.Println("GIF drawn.")
	fmt.Println("Exiting normally.")
}
