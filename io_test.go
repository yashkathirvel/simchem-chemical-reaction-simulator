package main

import (
	"fmt"
	"testing"
)

func TestReadParameters(t *testing.T) {
	surfaceWidth, timeStep, canvasWidth, scalingFactor, generation, frequency, speciesList, reactionMap := ReadParameters("input")
	fmt.Println(surfaceWidth)
	fmt.Println(timeStep)
	fmt.Println(canvasWidth)
	fmt.Println(scalingFactor)
	fmt.Println(generation)
	fmt.Println(frequency)
	fmt.Println(speciesList)
	fmt.Println(reactionMap["zeroth"][0].reactants[0].name)
	fmt.Println(reactionMap["uni"][0].reactants[0].name)
	fmt.Println(reactionMap["bi"][0].reactants[0].name)
}
