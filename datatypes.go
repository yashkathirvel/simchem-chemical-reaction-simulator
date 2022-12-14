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
	name             string // A, B or C
	mass             float64
	diffusionRate    float64
	red, green, blue uint8
	radius           float64
}

type Surface struct {
	width float64
	//particles []*Particle
	A_particles []*Particle
	B_particles []*Particle
	C_particles []*Particle
}

type OrderedPair struct {
	x, y float64
}
