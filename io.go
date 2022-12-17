package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Read parameters from .txt file
// return: width, timeStep, generation, species list, reaction list
func ReadParameters(filename string) (surfaceWidth, timeStep, scalingFactor float64, generation, canvasWidth, frequency int, speciesList map[*Species]int, reactionMap map[string][]Reaction) {
	surfaceWidth = 200
	timeStep = 1.0
	generation = 10
	//holds the input of initial number of all molecules
	speciesList = make(map[*Species]int, 10)
	//for reference in this function
	speciesMap := make(map[string]*Species, 10)
	//a map of "zero","uni","bi" and a slice of their corresponding reactions
	reactionMap = make(map[string][]Reaction, 3)

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
		//focus on lines with params only, skipping the rest
		if !(strings.HasPrefix(line, "#") || strings.HasPrefix(line, "$") || strings.HasPrefix(line, "u>") || strings.HasPrefix(line, "b>") || strings.HasPrefix(line, "z>")) {
			continue
		}

		//global parameters
		if strings.HasPrefix(line, "$") {
			//eat prefix
			line = line[1:]
			// Split the line on the delimiter
			fields := strings.Split(line, " ")

			// read surface width
			surfaceWidth, err = strconv.ParseFloat(fields[0], 64)
			if err != nil {
				panic("surface width is not a float64")
			}
			// read timeStep
			timeStep, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				panic("timeStep is not a float64")
			}
			// read generation
			generation, err = strconv.Atoi(fields[2])
			if err != nil {
				panic("generation is not an int.")
			}
			scalingFactor, err = strconv.ParseFloat(fields[3], 64)
			if err != nil {
				panic("scaling factor is not a float64")
			}
			// read generation
			canvasWidth, err = strconv.Atoi(fields[4])
			if err != nil {
				panic("canvas width is not an int. It should be pixels of canvas")
			}
			frequency, err = strconv.Atoi(fields[5])
			if err != nil {
				panic("frequency is not an int")
			}

		}
		//declaration of species
		if strings.HasPrefix(line, "#") {
			//eat prefix
			line = line[1:]
			// Split the line on the delimiter
			fields := strings.Split(line, " ")
			A, num := ReadSpecies(fields)
			speciesList[&A] = num
			speciesMap[A.name] = &A
		}
		//fmt.Println(speciesMap["A"])
		if strings.HasPrefix(line, "z>") {
			//eat prefix
			line = line[2:]
			// Split the line on the delimiter
			fields := strings.Split(line, " ")
			R := ReadZeroReaction(fields, speciesMap)
			reactionMap["zeroth"] = append(reactionMap["zeroth"], R)
		}
		if strings.HasPrefix(line, "u>") {
			//eat prefix
			line = line[2:]
			// Split the line on the delimiter
			fields := strings.Split(line, " ")
			R := ReadUniReaction(fields, speciesMap)
			//store it in the map
			reactionMap["uni"] = append(reactionMap["uni"], R)
		}
		if strings.HasPrefix(line, "b>") {
			//eat prefix
			line = line[2:]
			// Split the line on the delimiter
			fields := strings.Split(line, " ")
			R := ReadBiReaction(fields, speciesMap)
			reactionMap["bi"] = append(reactionMap["bi"], R)
		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	//surfaceWidth, timeStep, scalingFactor float64, generation, canvasWidth, frequency int, speciesList map[Species]int, reactionMap map[string][]Reaction
	return surfaceWidth, timeStep, scalingFactor, generation, canvasWidth, frequency, speciesList, reactionMap
}

// Order: name  radius  mass  diffusion-rate color initial-number
// returns The species object and how many of it
func ReadSpecies(fields []string) (A Species, num int) {
	var err error
	A.name = fields[0]
	A.radius, err = strconv.ParseFloat(fields[1], 64)
	if err != nil {
		panic("radius of molecule is not a float64")
	}
	A.mass, err = strconv.ParseFloat(fields[2], 64)
	if err != nil {
		panic("mass of molecule is not a float64")
	}
	A.diffusionRate, err = strconv.ParseFloat(fields[3], 64)
	if err != nil {
		panic("diffusionRate is not a float64")
	}
	A.color = fields[4]
	num, err = strconv.Atoi(fields[5])
	if err != nil {
		num = 0
		//panic("number of molecules is not an int")
	}

	return A, num
}

// input: the line descrbing reaction and the map of all species
// output: reaction objects
func ReadZeroReaction(fields []string, speciesMap map[string]*Species) Reaction {
	var reaction Reaction
	var err error
	speciesName := fields[0]
	reaction.reactants = append(reaction.reactants, speciesMap[speciesName])
	reaction.reactionConstant, err = strconv.ParseFloat(fields[len(fields)-1], 64)
	if err != nil {
		panic("reaction constant of zeroth order is not a float64")
	}
	return reaction
}

func ReadUniReaction(fields []string, speciesMap map[string]*Species) Reaction {
	var reaction Reaction
	var err error
	speciesName := fields[0]
	reaction.reactants = append(reaction.reactants, speciesMap[speciesName])
	if len(fields) > 2 {
		for i := 1; i < len(fields)-1; i++ {
			speciesName := fields[i]
			reaction.products = append(reaction.products, speciesMap[speciesName])
		}
	} //otherwise it vanishes
	reaction.reactionConstant, err = strconv.ParseFloat(fields[len(fields)-1], 64)
	if err != nil {
		panic("reaction constant of unimolecular reaction  is not a float64")
	}
	return reaction
}

func ReadBiReaction(fields []string, speciesMap map[string]*Species) Reaction {
	var reaction Reaction
	var err error
	if len(fields) > 3 { //make sure it is a bimolecular reaction
		AName := fields[0]
		reaction.reactants = append(reaction.reactants, speciesMap[AName])
		BName := fields[1]
		reaction.reactants = append(reaction.reactants, speciesMap[BName])
		for i := 2; i < len(fields)-1; i++ {
			speciesName := fields[i]
			reaction.products = append(reaction.products, speciesMap[speciesName])
		}
	} else {
		panic("bimolecular reaction not legit")
	}
	reaction.reactionConstant, err = strconv.ParseFloat(fields[len(fields)-1], 64)
	if err != nil {
		panic("reaction constant of bimolecular reaction is not a float64")
	}
	return reaction
}
