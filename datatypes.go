package main

type Particle struct {
	position OrderedPair
	species  *Species // pointer to the type of species
}

type Species struct { // every type of particles should have uniform following parameters
	name             string // A, B or C
	diffusionRate    float64
	red, green, blue uint8
	radius           float64
}

type Surface struct {
	width     float64
	particles []*Particle
}

type OrderedPair struct {
	x, y float64
}
