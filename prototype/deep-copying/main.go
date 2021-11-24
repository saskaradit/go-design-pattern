package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John", &Address{"Groove Street", "Los Santos", "USA"}}

	rad := john
	rad.Name = "Rad" // ok

	jane := john
	// Deep copying
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		jane.Address.Country,
	}
	jane.Name = "Jane"

	//rad.Address.StreetAddress = "Baker Street" -> this will change johns address
	jane.Address.StreetAddress = "Avenue" // this will not change johns address

	fmt.Println(rad, rad.Address)
	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
