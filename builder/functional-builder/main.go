package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) WorksAs(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, v := range b.actions {
		v(&p)
	}
	return &p
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Rad").WorksAs("Software Engineer").Build()
	fmt.Println(p)
}
