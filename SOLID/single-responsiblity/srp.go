package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// Separation of concerns

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// Create a new object so that if other objects have the same purpose, we can use this object
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I coded today")
	j.AddEntry("I almost died today")
	j.AddEntry("I am sad today")
	fmt.Println(j.String())

	// SaveToFile(&j, "journal.txt")

	// using the new persistence so that if another object needs to save to a file
	p := Persistence{"\n"}
	p.SaveToFile(&j, "./journal.txt")
}
