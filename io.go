package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Read parameters from .txt file
// return: width, timeStep, generation, species list, reaction list
func ReadParameters(filename string) (surfaceWidth, timeStep, canvasWidth, scalingFactor float64, generation, frequency int, speciesList []Species, reactionList []Reaction) {
	surfaceWidth = 200
	timeStep = 1.0
	generation = 10

	file, err := os.Open(filename + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// Set the delimiter to a comma
	scanner.Split(bufio.ScanLines)

	// Loop through the scanner
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()

		// inital parameters
		if strings.HasPrefix(line, "#") {
			//eat prefix
			line = line[1:]
			// Split the line on the delimiter
			fields := strings.Split(line, ",")

			// read width
			surfaceWidth, err = strconv.ParseFloat(fields[0], 64)
			if err != nil {
				panic("surface width is not a float64")
			}
			// read timeStep
			timeStep, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				panic("timeStep is not a float64")
			}
			// read canvasWidth
			canvasWidth, err = strconv.ParseFloat(fields[2], 64)
			if err != nil {
				panic("canvas width is not a float64")
			}
			scalingFactor, err = strconv.ParseFloat(fields[3], 64)
			if err != nil {
				panic("scaling factor is not a float64")
			}
			// read generation
			generation, err = strconv.Atoi(fields[4])
			if err != nil {
				panic("generation is not an int")
			}
			frequency, err = strconv.Atoi(fields[5])
			if err != nil {
				panic("frequency is not an int")
			}

		}
		//declaration of species
		if strings.HasPrefix(line, ">") {

		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return surfaceWidth, timeStep, canvasWidth, scalingFactor, generation, frequency, speciesList, reactionList
}

// Order: Name, Diffusion rate,
func ReadSpecies(string) (A Species) {

	return A
}
