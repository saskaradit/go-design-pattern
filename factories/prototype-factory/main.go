package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Lead
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Lead:
		return &Employee{"", "Developer", 80000}
	default:
		panic("Unsupported role")
	}
}

func main() {
	m := NewEmployee(Lead)
	m.Name = "rad"
	fmt.Println(m)
}
