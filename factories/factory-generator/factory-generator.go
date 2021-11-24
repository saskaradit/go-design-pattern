package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	// Cannot modify attributes using this approach (reccomended)
	developerFactory := NewEmployeeFactory("Developer", 20000)
	managerFactory := NewEmployeeFactory("Manager", 120000)

	developer := developerFactory("rad")
	manager := managerFactory("saskara")

	fmt.Println(developer)
	fmt.Println(manager)

	// Can modify attributes
	leadFactory := NewEmployeeFactory2("Tech Lead", 20000000)
	leadFactory.AnnualIncome = 10000000
	lead := leadFactory.Create("Rad")
	fmt.Println(lead)
}
