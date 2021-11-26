package main

import (
	"container/list"
	"fmt"
)

// Observable, Observer

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(a Observer) {
	o.subs.PushBack(a)
}

func (o *Observable) Unsubscribe(a Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == a {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string // "Age", "Height"
	Value interface{}
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person) Age() int { return p.age }

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

type BirthdayManagement struct {
	o Observable
}

func (t *BirthdayManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 21 {
			fmt.Println("Congratulations! You are now legal")
			t.o.Unsubscribe(t)
		} else {
			fmt.Println("Happy Birthday!!")
		}
	}
}

func main() {
	p := NewPerson(15)
	t := &BirthdayManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 26; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
