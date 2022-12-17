package main

type Reaction struct {
	reactants        []*Species
	products         []*Species
	reactionConstant float64
}
type Particle struct {
	position OrderedPair
	velocity OrderedPair
	species  *Species // pointer to the type of species
}

type Species struct { // every type of particles should have uniform following parameters
	name          string // A, B or C
	radius        float64
	mass          float64
	diffusionRate float64
	color         string
}

type Surface struct {
	width float64
	//particles []*Particle
	molecularMap map[*Species][]*Particle
}

type OrderedPair struct {
	x, y float64
}
