package main

type Person struct {
	Name  string
	Age   int
	Alive bool
}

func NewPerson(name string, age int) *Person {
	if age > 20 {
		// ...
	}
	// It has a default value
	return &Person{name, age, true}
}

func main() {
	// This is the normal object creation
	_ = Person{"Rad", 21, true}

	// Creates the object without defining everything that may not need to be customised
	p := NewPerson("Saskara", 21)
	// If need for customisation
	p.Alive = false

}
