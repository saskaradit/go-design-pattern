package main

// Interface segregation

type Document struct {
}

// This interface needs to be break down, because it has too many attributes
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d Document) {}
func (m *MultiFunctionPrinter) Fax(d Document)   {}
func (m *MultiFunctionPrinter) Scan(d Document)  {}

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d Document) {
	// prints something
}
func (o *OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}
func (o *OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP
// this interface works well
type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}

type RadPrinter struct {
}

func (r *RadPrinter) Print(d Document) {
	// print something
}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	//
}
func (p Photocopier) Scan(d Document) {
	//
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// Decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func main() {

}
