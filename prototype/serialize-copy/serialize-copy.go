package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)

	return &result
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
