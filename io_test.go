package main

import (
	"fmt"
	"testing"
)

func TestReadParameters(t *testing.T) {
	surfaceWidth, timeStep, canvasWidth, scalingFactor, generation, frequency, speciesList, reactionList := ReadParameters("input")
	fmt.Println(surfaceWidth)
	fmt.Println(timeStep)
	fmt.Println(canvasWidth)
	fmt.Println(scalingFactor)
	fmt.Println(generation)
	fmt.Println(frequency)
	fmt.Println(speciesList)
	fmt.Println(reactionList)
}
