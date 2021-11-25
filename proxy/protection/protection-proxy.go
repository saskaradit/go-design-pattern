package main

import "fmt"

type Drive interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		fmt.Println("You are too young")
	}
}

func main() {
	car := NewCarProxy(&Driver{20})
	car.Drive()
}
