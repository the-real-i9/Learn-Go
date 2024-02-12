package main

import (
	"fmt"
	"strings"
	"time"
)

var joinString = strings.Join // method value
var jnd = joinString([]string{"k", "e", "n"}, " | ")

type Rocket struct{}

func (r *Rocket) Launch() {
	fmt.Println("Rocket launched")
}

var sr = Rocket{}

var distance = Point.Distance // method expression
var scale = (*Point).ScaleBy

type Path []Point

func (path Path) TranslatedBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func methValsExp() {
	fmt.Println(jnd)
	time.AfterFunc(10*time.Second, sr.Launch)
	// time.Sleep(12 * time.Second)

	p := Point{1, 2}
	q := Point{4, 6}

	scale(&p, 2)
	scale(&q, 2)

	fmt.Println(distance(p, q))
}
