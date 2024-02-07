package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// methods with a pointer receiver, in case you want to modify the receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type IntList struct {
	Value int
	Tail  *IntList
}

// nil is a valid receiver value
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func methDecl() {
	fmt.Println(p.Distance(q.Point))
}
