package main

import "fmt"

// Circle, Square
// Raster, Vector

// RasterCircle, VectorCircle, RasterSquare, ...
// How to simplify?

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	// ...
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Draw circle of a radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixel for circle with radius of", radius)
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer, radius}
}

func main() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	circle := NewCircle(&raster, 6)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()

	ci := NewCircle(&vector, 3)
	ci.Draw()
	ci.Resize(2)
	ci.Draw()
}
