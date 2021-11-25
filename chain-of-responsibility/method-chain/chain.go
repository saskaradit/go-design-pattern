package main

import "fmt"

type Creature struct {
	Name             string
	Attack, Deffense int
}

func (c *Creature) Summon() string {
	return fmt.Sprintln(c.Name, "has been summonned with stats of", c.Attack, "Atk power", c.Deffense, "Def power")
}

func NewCreature(name string, attack int, deffense int) *Creature {
	return &Creature{name, attack, deffense}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (c *CreatureModifier) Add(mod Modifier) {
	if c.next != nil {
		c.next.Add(mod)
	} else {
		c.next = mod
	}
}
func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "attack!!")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type IncreaseDeffenseModifier struct {
	CreatureModifier
}

func NewIncreaseDeffenseModifier(c *Creature) *IncreaseDeffenseModifier {
	return &IncreaseDeffenseModifier{CreatureModifier{creature: c}}
}

func (i *IncreaseDeffenseModifier) Handle() {
	fmt.Println("Increasing", i.creature.Name, "deffense!!")
	i.creature.Deffense = i.creature.Deffense + i.creature.Attack/2
	i.CreatureModifier.Handle()
}

type RemoveBoostsModifier struct {
	CreatureModifier
}

func NewRemoveBoostsModifier(c *Creature) *RemoveBoostsModifier {
	return &RemoveBoostsModifier{CreatureModifier{creature: c}}
}

func (r *RemoveBoostsModifier) Handle() {
	fmt.Println("Removing boosts")
}

func main() {
	chimera := NewCreature("Chimera", 20, 20)
	fmt.Println(chimera.Summon())

	root := NewCreatureModifier(chimera)
	// This modifier breaks the chain of command
	root.Add(NewRemoveBoostsModifier(chimera))

	root.Add(NewDoubleAttackModifier(chimera))
	root.Add(NewIncreaseDeffenseModifier(chimera))
	root.Handle()
	fmt.Println(chimera.Summon())
}
