package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}
func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

func main() {
	drawing := GraphicObject{"Rad Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle("Black"))
	drawing.Children = append(drawing.Children, *NewCircle("Blue"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Black"))
	group.Children = append(group.Children, *NewCircle("Orange"))
	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}
