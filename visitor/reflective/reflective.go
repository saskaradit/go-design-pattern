package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	// nothing here!
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func Print(e Expression, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		sb.WriteString("(")
		Print(ae.left, sb)
		sb.WriteString("+")
		Print(ae.right, sb)
		sb.WriteString(")")
	}
	// need to add more if need to extend
	// breaks OCP
	// will work incorrectly on missing case
}

func main() {
	// 1 + (2+3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	// sb is the visitor
	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}
