package main

import (
	"fmt"
	"gifhelper"
	"time"
)

func main() {
	// evolution parameters
	surfaceWidth, timeStep, scalingFactor, numGens, canvasWidth, frequency, speciesMap, reactionMap := ReadParameters("input")
	// initial Surface (for testing purposes)
	initialSurface := Surface{
		width: surfaceWidth,
	}
	initialSurface.molecularMap = make(map[*Species][]*Particle, len(speciesMap))
	//fmt.Println("speciesMap:", speciesMap)
	initialSurface.Initialization(speciesMap)
	fmt.Println("The map", len(initialSurface.molecularMap))
	for species := range initialSurface.molecularMap {
		fmt.Println("species:", species)
		fmt.Println("molecularMap", len(initialSurface.molecularMap[species]))
	}
	//good parameters for L-V
	//killRate := 2.0
	//zerothRateConstant := 2.0
	//bimolecularRateConstant := 40000.0
	// DRIVER CODE (DO NOT CHANGE!!)
	startTime := time.Now()
	timePoints := SimulateSurface(initialSurface, numGens, timeStep, reactionMap)
	//timePoints := SimulateSurfaceCollision(initialSurface, numGens, timeStep, reactionMap)
	elapsedTime := time.Since(startTime)

	fmt.Println("Simulation took", elapsedTime, "s. Now drawing images.")

	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, "collision")
	fmt.Println("GIF drawn.")
	fmt.Println("Exiting normally.")
}
