package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
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

// Separtaion of concerns

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(strings.Join(j.entries, LineSeparator), []byte(j.String()), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(strings.Join(j.entries, p.lineSeparator), []byte(j.String()), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I coded today")
	j.AddEntry("I almost died today")
	j.AddEntry("I am sad today")
	fmt.Println(j.String())

	// SaveToFile(&j, "journal.txt")

	p := Persistence{"\n"}
	p.SaveToFile(&j, "journal.txt")
}
