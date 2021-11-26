package main

import "fmt"

type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintln(sender, "\t:", message)
	fmt.Printf("[%s's chat session] \t %s", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (c *ChatRoom) Message(source, dest, message string) {
	for _, p := range c.people {
		if p.Name == dest {
			p.Receive(source, message)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := ChatRoom{}

	rad := NewPerson("Rad")
	fate := NewPerson("Fate")

	room.Join(rad)
	room.Join(fate)

	rad.Say("Hello there")
	fate.Say("The angel from my nightmare")

	ciara := NewPerson("Ciara")
	room.Join(ciara)

	ciara.Say("The shadow in the background of the morgue")

	rad.PrivateMessage("Fate", "What a string of fate")
}
