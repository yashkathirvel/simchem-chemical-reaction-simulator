package main

import (
	"math/rand"
)

func (s Surface) Initialization(speciesList map[*Species]int) {
	//random placement
	for species, number := range speciesList {
		//fmt.Println(species.name, number)
		for i := 0; i < number; i++ {
			p := &Particle{
				position: OrderedPair{rand.Float64() * s.width, rand.Float64() * s.width},
				species:  species, // pointer to the type of species
			}
			s.molecularMap[species] = append(s.molecularMap[species], p)
		}
		//fmt.Println(species.name, len(s.molecularMap[species]))
	}
	//fmt.Println("The map", s.molecularMap)
}
