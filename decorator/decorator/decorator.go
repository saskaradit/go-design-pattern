package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Cirlce of radius %v", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Its size is %v x %v", s.Side, s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%v has the color %v", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%v has the transparency of %v%%", t.Shape.Render(), t.Transparency*100)
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.55}
	fmt.Println(rhsCircle.Render())

}
