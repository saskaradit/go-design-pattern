package main

import "fmt"

type Person interface {
	SayHello()
}

// hidden from outside
type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, I'm too tired")
}
func (p *person) SayHello() {
	fmt.Println("Hi! my name is", p.name, "I am", p.age, "years old")
}

func NewPerson(name string, age int) Person {
	if age > 40 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	// cannot modify the underlying attributes after created
	p := NewPerson("Rad", 21)
	t := NewPerson("Jenkins", 50)
	p.SayHello()
	t.SayHello()
}
