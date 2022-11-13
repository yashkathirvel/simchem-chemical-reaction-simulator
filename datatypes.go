/*
This file contains all the classes in the simulation.
*/

package main

type OrderedPair struct {
	x, y float64
}

type Surface struct {
	particles []*Particle
	width     float64
}

type Particle struct {
	position         OrderedPair
	radius           float64
	red, green, blue uint8
}
