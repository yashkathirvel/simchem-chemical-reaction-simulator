package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func SimulateSurface(timePoints []*Surface, numGens int, timeStep float64, reactionMap map[string][]Reaction) []*Surface {
	//timePoints := make([]*Surface, numGens)
	// set the initial Surface object as the first time point
	//timePoints[0] = &initialSurface
	// iterate through numGens generations and update the Surface object each time.
	//dataA := make([]int, 0)
	//dataB := make([]int, 0)
	for i := 1; i < numGens; i++ {
		timePoints[i] = timePoints[i-1].Update(timeStep, reactionMap)
		time.Sleep(time.Nanosecond)
		//for testing purpose
		fmt.Println(len(timePoints[i].molecularMap[reactionMap["bi"][0].reactants[0]]), ",", len(timePoints[i].molecularMap[reactionMap["bi"][0].reactants[1]]), ",")
	}
	return timePoints
}

// Surface method: Update()
// Updates the Surface object given a time s
func (s *Surface) Update(timeStep float64, reactionMap map[string][]Reaction) *Surface {
	// create a copy of the current Surface object
	newS := s.Copy()
	rand.Seed(time.Now().UnixNano())
	// iterate through the particles on the surface and diffuse them
	for _, particles := range newS.molecularMap {
		for _, p := range particles {
			p.Diffuse(timeStep)
		}
	}

	//newS.LoktaVolterraReaction(bimolecularRateConstant, diffusion_cons_A, diffusion_cons_B)
	//newS.AddAParticles(zerothRateConstant, timeStep)
	//newS.KillParticles(killRate, timeStep)
	if reactionMap["zero"] != nil && len(reactionMap["zero"]) != 0 { //handling zeroth order, i.e. adding
		for _, reaction := range reactionMap["zero"] {
			newS.ZerothOrder(reaction, timeStep)
		}
	}
	if reactionMap["bi"] != nil && len(reactionMap["bi"]) != 0 { //handling bimolecular order,
		for _, reaction := range reactionMap["bi"] {
			newS.BimolecularReaction(reaction)
			//fmt.Println("ith bi reaction", i, "name", reaction.reactants[0].name, "num", len(newS.molecularMap[reaction.reactants[0]]), "name", reaction.reactants[1].name, "num", len(newS.molecularMap[reaction.reactants[1]]))
		}
	}
	if reactionMap["uni"] != nil && len(reactionMap["uni"]) != 0 { //handling uni order
		for _, reaction := range reactionMap["uni"] {
			if len(reaction.products) == 0 {
				newS.KillParticles(reaction, timeStep) //uni that only kills particles
			} else {
				newS.UnimolecularReaction(reaction, timeStep) //uni that has products
			}
			//fmt.Println("ith uni reaction", i, "name", reaction.reactants[0].name, "num", len(newS.molecularMap[reaction.reactants[0]]))
		}
	}
	return newS
}

func (newS *Surface) ZerothOrder(reaction Reaction, timeStep float64) {
	// initialize global pseudo random generator
	number := reaction.reactionConstant * timeStep
	//k0dt product molecules are formed during each time step.
	for i := 0; i < int(number); i++ {

		newParticle := Particle{
			position: OrderedPair{rand.Float64() * newS.width, rand.Float64() * newS.width},
			species:  reaction.reactants[0],
		}
		newS.molecularMap[newParticle.species] = append(newS.molecularMap[newParticle.species], &newParticle)
	}
}
func (newS *Surface) DeleteParticle(a *Particle) {
	//range through surface to find the index of the particle
	for i := 0; i < len(newS.molecularMap[a.species]); i++ {
		if newS.molecularMap[a.species][i] == a {
			newS.molecularMap[a.species] = append(newS.molecularMap[a.species][:i], newS.molecularMap[a.species][i+1:]...)
		}
	}
}

// Particle method: Copy{}
// Creates a deep copy of the particle object.
func (s *Particle) Copy() *Particle {
	// create new address for newP
	var newP Particle
	// copy the position and species of the particle
	newP.position.x = s.position.x
	newP.position.y = s.position.y
	newP.species = s.species
	return &newP
}

// Surface method: Copy()
// Creates a deep copy of the Surface object.
func (s *Surface) Copy() *Surface {
	// create new address for newS
	var newS Surface
	newS.molecularMap = make(map[*Species][]*Particle, len(s.molecularMap))
	newS.width = s.width
	// iterate through the particles on the surface
	for species, particles := range s.molecularMap {
		for _, particle := range particles {
			newParticle := particle.Copy()
			newS.molecularMap[species] = append(newS.molecularMap[species], newParticle)
		}
	}
	return &newS
}

// Particle method: SurfaceReaction(), this method takes into account the interaction of the particles with the surface
// the simulation is kept simple by defining boundaries, we simulate an inert permeable boundary
// this function takes the witdth of the surface and reflects particles back into the medium when they hit the surface
func (p *Particle) SurfaceReaction(width float64) {
	if p.position.x > width {
		p.position.x = p.position.x - (p.position.x - width)
	} else {
		if p.position.x < 0 {
			p.position.x = width - (p.position.x * (width / p.position.x))
		}
	}

	if p.position.y > width {
		p.position.y = p.position.y - (p.position.y - width)
	} else {
		if p.position.y < 0 {
			p.position.y = width - (p.position.y * (width / p.position.y))
		}
	}
}

// this function simulates the bimolecular reaction
// input: takes the rate constant of the reaction, calculates a binding radius from it which determines how far
// two species need to be from each other to initiate collision and consequently a chemical reaction
func (newS *Surface) BimolecularReaction(reaction Reaction) {
	//kSi = 4πDσb.
	binding_radius := reaction.reactionConstant / (4 * math.Pi * (reaction.reactants[0].diffusionRate + reaction.reactants[1].diffusionRate))
	//fmt.Println("radius:", binding_radius, reaction.reactants[0].name, reaction.reactants[0].diffusionRate, reaction.reactants[1].name, reaction.reactants[1].diffusionRate)
	new_distDictionary := make([]OrderedPair, 0)

	//making sure each pair of reactants just collide for 1 time
	flagsOfA := make([]bool, len(newS.molecularMap[reaction.reactants[0]]))
	flagsOfB := make([]bool, len(newS.molecularMap[reaction.reactants[1]]))
	//range through a and compare it's distance with the B particles
	//if the distance between them is less than the binding radius, make C_particles
	for i, b_particle := range newS.molecularMap[reaction.reactants[1]] {
		for j, a_particle := range newS.molecularMap[reaction.reactants[0]] {
			particle_dist := Distance(a_particle.position, b_particle.position)
			if particle_dist < binding_radius && !flagsOfA[j] && !flagsOfB[i] { //reaction happens
				//location of product to be added
				new_dist := Average_pos(a_particle.position, b_particle.position)
				new_distDictionary = append(new_distDictionary, new_dist)
				newS.DeleteParticle(a_particle)
				newS.DeleteParticle(b_particle)
				flagsOfA[j] = true
				flagsOfB[i] = true
				//newS.C_particles = append(newS.C_particles, &C_p)
				continue
			}
		}
	}
	//fmt.Println("how many bi", len(new_distDictionary))
	for i := range new_distDictionary {
		//fmt.Println("how many products", len(reaction.products))
		for _, product := range reaction.products {
			product_Particle := Particle{
				position: new_distDictionary[i],
				species:  product,
			}
			newS.molecularMap[product] = append(newS.molecularMap[product], &product_Particle)
			//fmt.Println("how many B", len(newS.molecularMap[product]))
		}
	}
	//fmt.Println("how many products B", len(newS.molecularMap[reaction.products[1]]))
}

func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return (math.Sqrt(deltaX*deltaX + deltaY*deltaY))
}

func Average_pos(p1, p2 OrderedPair) OrderedPair {
	//calculates avergae of the two positions
	var dist OrderedPair
	deltaX := (p1.x + p2.x) / 2
	deltaY := (p1.y + p2.y) / 2

	dist.x = deltaX
	dist.y = deltaY

	return dist
}

// Surface method: Delete random B particles
func (s *Surface) DeleteRandomParticle(species *Species, i int) {
	//range through surface to find the index of the particle
	s.molecularMap[species] = append(s.molecularMap[species][:i], s.molecularMap[species][i+1:]...)

}

// i.e. unimolecular reaction that removes 1 species
func (newS *Surface) KillParticles(reaction Reaction, timeStep float64) {
	// initialize global pseudo random generator
	time.Sleep(time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	prob := 1.0 - math.Exp(-reaction.reactionConstant*timeStep)
	deathList := make([]int, 0)
	for i := range newS.molecularMap[reaction.reactants[0]] {
		dice := rand.Float64()
		//fmt.Println("prob", prob)
		if dice < prob {
			deathList = append(deathList, i)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(deathList)))
	for i := range deathList {
		newS.DeleteRandomParticle(reaction.reactants[0], deathList[i])
	}
}

func (newS *Surface) UnimolecularReaction(reaction Reaction, timeStep float64) {
	// initialize global pseudo random generator
	time.Sleep(time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	prob := 1.0 - math.Exp(-reaction.reactionConstant*timeStep)
	new_distDictionary := make([]OrderedPair, 0)
	for _, particle := range newS.molecularMap[reaction.reactants[0]] {
		if rand.Float64() < prob {
			new_distDictionary = append(new_distDictionary, particle.position)
			newS.DeleteParticle(particle)
		}
	}
	//adding products
	for i := range new_distDictionary {
		for _, product := range reaction.products {
			product_Particle := Particle{
				position: new_distDictionary[i],
				species:  product,
			}
			newS.molecularMap[product] = append(newS.molecularMap[product], &product_Particle)
		}
	}
}
