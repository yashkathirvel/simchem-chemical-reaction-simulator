package main

type Particle struct {
	x       float64
	y       float64
	species *Species //should be A,B OR C
}

type Species struct { //every type of particles should have uniform following parameters
	name             string //A,B or C
	diffusionRate    float64
	red, green, blue uint8
	radius           float64
}

type Surface struct {
	width     float64
	particles []*Particle
}
