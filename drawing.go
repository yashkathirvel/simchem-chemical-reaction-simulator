package main

import (
	"canvas"
	//"fmt"
	"image"
)

// AnimateSystem takes a slice of Surface objects along with a canvas width
// parameter and a frequency parameter.
// Every frequency steps, it generates a slice of images corresponding to drawing each Surface
// on a canvasWidth x canvasWidth canvas.
// A scaling factor is a final input that is used to scale the particles.
func AnimateSystem(timePoints []*Surface, canvasWidth, frequency int, scalingFactor float64) []image.Image {
	images := make([]image.Image, 0)

	if len(timePoints) == 0 {
		panic("Error: no Surface objects present in AnimateSystem.")
	}

	// for every universe, draw to canvas and grab the image
	for i := range timePoints {
		if i%frequency == 0 {
			// fmt.Println(i)
			images = append(images, timePoints[i].DrawToCanvas(canvasWidth, scalingFactor))
		}
		//fmt.Println("Frame:", i)
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a Surface
// object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels.
// A scaling factor allows us to scale the particle size on the canvas.
func (s *Surface) DrawToCanvas(canvasWidth int, scalingFactor float64) image.Image {
	if s == nil {
		panic("Can't draw a nil surface.")
	}

	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the particles and draw them.
	for _, a := range s.A_particles {
		c.SetFillColor(canvas.MakeColor(a.species.red, a.species.green, a.species.blue))
		cx := (a.position.x / s.width) * float64(canvasWidth) //
		cy := (a.position.y / s.width) * float64(canvasWidth)
		r := scalingFactor * (a.species.radius / s.width) * float64(canvasWidth)
		c.Circle(cx, cy, r)
		c.Fill()
	}
	for _, b := range s.B_particles {
		c.SetFillColor(canvas.MakeColor(b.species.red, b.species.green, b.species.blue))
		cx := (b.position.x / s.width) * float64(canvasWidth) //
		cy := (b.position.y / s.width) * float64(canvasWidth)
		r := scalingFactor * (b.species.radius / s.width) * float64(canvasWidth)
		c.Circle(cx, cy, r)
		c.Fill()
	}
	for _, C := range s.C_particles {
		c.SetFillColor(canvas.MakeColor(C.species.red, C.species.green, C.species.blue))
		cx := (C.position.x / s.width) * float64(canvasWidth) //
		cy := (C.position.y / s.width) * float64(canvasWidth)
		r := scalingFactor * (C.species.radius / s.width) * float64(canvasWidth)
		c.Circle(cx, cy, r)
		c.Fill()
	}

	// we want to return an image!
	return canvas.GetImage(c)
}
