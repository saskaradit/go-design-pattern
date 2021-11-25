package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{fullName}
}

var allNames []string

type User2 struct {
	names []uint8 // assume only 256 unique names
}

func NewUser2(fullName string) *User2 {
	// memory savings
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		// does not store the same characcters twice
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func main() {
	rad := NewUser("Rad Saskara")

	fmt.Println(rad.FullName)

	logic := NewUser2("Robert Bryson")
	alsoLogic := NewUser2("Sir Robert")
	fmt.Println(logic.FullName())
	fmt.Println(alsoLogic.FullName())

	fmt.Println(allNames)
}
