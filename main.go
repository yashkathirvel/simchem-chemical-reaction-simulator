package main

import (
	"flag"
	"fmt"
	"gifhelper"
	"os"
	"time"
)

func main() {

	collisionPtr := flag.Bool("c", false, "Check -c if you wish to carry out collision-driven simulation.")
	if len(os.Args) < 3 {
		panic("missing input/output name.")
	}
	fileName := os.Args[1]
	outputName := os.Args[2]

	// Parse the command-line flags.
	flag.Parse()

	// evolution parameters
	surfaceWidth, timeStep, scalingFactor, numGens, canvasWidth, frequency, speciesMap, reactionMap := ReadParameters(fileName)
	// initial Surface (for testing purposes)
	initialSurface := Surface{
		width: surfaceWidth,
	}
	initialSurface.molecularMap = make(map[*Species][]*Particle, len(speciesMap))

	initialSurface.Initialization(speciesMap)

	//good parameters for L-V
	//killRate := 2.0
	//zerothRateConstant := 2.0
	//bimolecularRateConstant := 40000.0
	// DRIVER CODE (DO NOT CHANGE!!)
	timePoints := make([]*Surface, numGens)
	timePoints[0] = &initialSurface
	startTime := time.Now()
	if *collisionPtr {
		//to be implemented
		timePoints[0].SetInitialVelocity(timeStep)
		timePoints = SimulateSurfaceCollision(timePoints, numGens, timeStep, reactionMap)
	} else {
		timePoints = SimulateSurface(timePoints, numGens, timeStep, reactionMap)
	}
	//timePoints := SimulateSurfaceCollision(initialSurface, numGens, timeStep, reactionMap)
	elapsedTime := time.Since(startTime)

	fmt.Println("Simulation took", elapsedTime, "s. Now drawing images.")

	imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(imageList, outputName)
	fmt.Println("GIF drawn.")
	fmt.Println("Exiting normally.")
}
