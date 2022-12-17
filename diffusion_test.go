package main

/**
func TestSimulatorSurface(t *testing.T) {
	numGens := 20
	timeStep := 5.00
	//Declaring A particle
	A := ParticleType{
		diffusionRate: 1200.0,
		radius:        250,
		red:           132,
		green:         83,
		blue:          60,
	}
	// initial Surface (for testing purposes)
	initialSurface := &Surface{
		particles: []*Particle{
			{
				x:       200000,
				y:       200000,
				species: &A,
			},
		},
		width: 400000,
	}
	timePoints := SimulateSurface(initialSurface, numGens, timeStep)
	fmt.Print("updated x:", timePoints[1].particles[0].x)
	fmt.Print("updated x:", timePoints[19].particles[0].x)
}
**/
/**
func TestDiffuse(t *testing.T) {
	A := Species{
		diffusionRate: 12.0,
		radius:        250,
		red:           132,
		green:         83,
		blue:          60,
	}
	particle := Particle{
		position: OrderedPair{200, 200},
		species:  &A,
	}

	var s Surface
	s.particles = append(s.particles, &particle)
	s.Diffuse(1.00)
	fmt.Print(s.particles[0].position.x)
	fmt.Print(s.particles[0].position.y)
}
**/
