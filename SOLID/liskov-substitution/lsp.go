package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// This is wrong, because it overrides attributes that are defined in the higher scopes
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

// This is wrong, because it overrides attributes that are defined in the higher scopes
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected area of", expectedArea, "but got", actualArea)
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
}
