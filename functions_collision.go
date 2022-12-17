package main

import (
	"math"
	"math/rand"
	"time"
)

func SimulateSurfaceCollision(timePoints []*Surface, numGens int, timeStep float64, reactionMap map[string][]Reaction) []*Surface {

	//fmt.Print(initialS.A_particles[50].velocity)
	// iterate through numGens generations and update the Surface object each time.
	for i := 1; i < numGens; i++ {
		timePoints[i] = timePoints[i-1].UpdateCollision(timeStep, reactionMap)
		//fmt.Println("Generation: ", i)
	}
	return timePoints
}

func (s *Surface) SetInitialVelocity(timeStep float64) {
	rand.Seed(time.Now().UnixNano())
	for _, particles := range s.molecularMap {
		for _, p := range particles {
			p.SetInitialVelocity(timeStep)
		}
	}
}

// Surface method: Update()
// Updates the Surface object given a time s
func (s *Surface) UpdateCollision(timeStep float64, reactionMap map[string][]Reaction) *Surface {
	// create a copy of the current Surface object
	newS := s.CopyCollision()
	//reset PRNG before diffusion
	rand.Seed(time.Now().UnixNano())
	// iterate through the particles on the surface and diffuse them
	newS.DiffuseCollision(timeStep)
	//reaction map recorded all zero,bimolecular and unimolecular reactions in a string-array map.

	if reactionMap["zero"] != nil && len(reactionMap["zero"]) != 0 { //handling zeroth order, i.e. adding particles
		//range through all zeroth order reaction
		for _, reaction := range reactionMap["zero"] {
			newS.ZerothOrderCollision(reaction, timeStep)
		}
	}
	if reactionMap["uni"] != nil && len(reactionMap["uni"]) != 0 { //handling uni order
		for _, reaction := range reactionMap["uni"] {
			if len(reaction.products) == 0 {
				newS.KillParticles(reaction, timeStep) //uni that only kills particles
			} else {
				newS.UnimolecularReactionCollision(reaction, timeStep) //uni that has products
			}
			//fmt.Println("ith uni reaction", i, "name", reaction.reactants[0].name, "num", len(newS.molecularMap[reaction.reactants[0]]))
		}
	}
	if reactionMap["bi"] != nil && len(reactionMap["bi"]) != 0 { //handling bimolecular order,
		for _, reaction := range reactionMap["bi"] {
			newS.BimolecularReactionCollision(reaction, timeStep)
			//fmt.Println("ith bi reaction", i, "name", reaction.reactants[0].name, "num", len(newS.molecularMap[reaction.reactants[0]]), "name", reaction.reactants[1].name, "num", len(newS.molecularMap[reaction.reactants[1]]))
		}
	}
	return newS
}

// Particle method: Copy{}
// Creates a deep copy of the particle object.
func (s *Particle) CopyCollision() *Particle {
	// create new address for newP
	var newP Particle
	// copy the position and species of the particle
	newP.position.x = s.position.x
	newP.position.y = s.position.y
	newP.velocity.x = s.velocity.x
	newP.velocity.y = s.velocity.y
	newP.species = s.species
	return &newP
}

// Surface method: Copy()
// Creates a deep copy of the Surface object.
func (s *Surface) CopyCollision() *Surface {

	// iterate through the particles on the surface
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
func (newS *Surface) UnimolecularReactionCollision(reaction Reaction, timeStep float64) {
	// initialize global pseudo random generator
	time.Sleep(time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	prob := 1.0 - math.Exp(-reaction.reactionConstant*timeStep)
	new_distDictionary := make([]OrderedPair, 0)
	new_velDictionary := make([]OrderedPair, 0)
	for _, particle := range newS.molecularMap[reaction.reactants[0]] {
		if rand.Float64() < prob {
			new_distDictionary = append(new_distDictionary, particle.position)
			new_distDictionary = append(new_distDictionary, particle.velocity)
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
			product_Particle.velocity.x = -new_velDictionary[i].x
			product_Particle.velocity.y = -new_velDictionary[i].y
			newS.molecularMap[product] = append(newS.molecularMap[product], &product_Particle)
		}
	}
}
func (newS *Surface) ZerothOrderCollision(reaction Reaction, timeStep float64) {
	// initialize global pseudo random generator
	number := reaction.reactionConstant * timeStep
	//k0dt product molecules are formed during each time step.
	std := math.Sqrt(2 * timeStep * reaction.reactants[0].diffusionRate)
	for i := 0; i < int(number); i++ {

		newParticle := Particle{
			position: OrderedPair{rand.Float64() * newS.width, rand.Float64() * newS.width},
			velocity: OrderedPair{rand.NormFloat64() * std, rand.NormFloat64() * std},
			species:  reaction.reactants[0],
		}
		newS.molecularMap[newParticle.species] = append(newS.molecularMap[newParticle.species], &newParticle)
	}
}
func (newS *Surface) BimolecularReactionCollision(reaction Reaction, timeStep float64) {
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
			p := Particle{
				position: new_distDictionary[i],
				species:  product,
			}
			p.Diffuse(timeStep)
			p.SetInitialVelocity(timeStep)
			newS.molecularMap[product] = append(newS.molecularMap[product], &p)
			//fmt.Println("how many B", len(newS.molecularMap[product]))
		}
	}
	//fmt.Println("how many products B", len(newS.molecularMap[reaction.products[1]]))
}
