package main

import (
	"fmt"
	"sync"
)

// CoR, Mediator, Observer, CQS

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Creature struct {
	game            *Game
	Name            string
	attack, defense int
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{game, name, attack, defense}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Deffense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintln(c.Name, "has been summonned with stats of", c.Attack(), "Atk power", c.Deffense(), "Def power")
}

type Observer interface {
	Handle(q *Query)
}
type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}
func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

type DoubleAttackModifier struct {
	CreatureModifier
}

// func (c *CreatureModifier) Handle(q *Query) {}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{g, c}}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func main() {
	game := &Game{sync.Map{}}
	chimera := NewCreature(game, "Lion Chimera", 3, 4)
	fmt.Println(chimera.String())
	{
		m := NewDoubleAttackModifier(game, chimera)
		fmt.Println(chimera.String())
		m.Close()
	}
	fmt.Println(chimera.String())

}
