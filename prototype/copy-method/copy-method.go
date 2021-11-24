package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	john := Person{"John", &Address{"Groove Street", "Los Santos", "USA"}, []string{"Rad", "Doe"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker street"
	jane.Friends = append(jane.Friends, "Kei")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
