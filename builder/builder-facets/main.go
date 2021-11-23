package main

import "fmt"

type Person struct {
	// Address
	StreetAddress, Postcode, City string

	// Job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	b.person.StreetAddress = street
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

func (b *PersonJobBuilder) WorksAt(name string) *PersonJobBuilder {
	b.person.CompanyName = name
	return b
}
func (b *PersonJobBuilder) Title(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}
func (b *PersonJobBuilder) Earning(inc int) *PersonJobBuilder {
	b.person.AnnualIncome = inc
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("Groove Street").
		In("Los Santos").
		WithPostcode("175111").
		Works().
		WorksAt("Google").
		Title("Software Engineer").
		Earning(2000000)
	person := pb.Build()
	fmt.Println(person)
}
