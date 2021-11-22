package main

import "fmt"

// OCP
// open for extension, closed for modification
// similar to specification approach
type Color int

const (
	red Color = iota
	green
	black
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}
type Filter struct {
	// ...
}

// This is the wrong approach
// ===========================

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// and so on ....

// ===========================
// This is the wrong approach

// Implementing Specification
// Below is the correct approach
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == c.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct{}

// Filter implements filtering by accepting a specification interface and returns a result array
func (f *BetterFilter) Filter(produts []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range produts {
		if spec.IsSatisfied(&v) {
			result = append(result, &produts[i])
		}
	}
	return result
}

func main() {
	shirt := Product{"Shirt", black, small}
	tree := Product{"Tree", green, large}
	car := Product{"Car", black, large}

	products := []Product{shirt, tree, car}
	fmt.Printf("Green Products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("Green Products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("Large and Black Products (new):\n")
	blackSpec := ColorSpecification{black}
	largeSpec := SizeSpecification{large}
	lbSpec := AndSpecification{blackSpec, largeSpec}
	for _, v := range bf.Filter(products, lbSpec) {
		fmt.Printf(" - %s is large and black\n", v.name)
	}

}
