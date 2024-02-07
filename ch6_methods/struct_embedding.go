package main

import (
	"fmt"
	"image/color"
	"sync"
)

/* Composing Types by Struct Embedding */
// Named types also inherits the methods of embedded ones.
// This is just a mere embedding not subclassing, therefore, ColoredPoint is not a point
// It is just merely for bypassing embedded type access, the compiler explicitly constructs the whole chain
type ColoredPoint struct {
	Point
	Color color.RGBA
}

var red = color.RGBA{255, 0, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

var p = ColoredPoint{Point{1, 1}, red}
var q = ColoredPoint{Point{5, 4}, blue}

// Embedding makes it possible for unnamed struct types to have methods too
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{mapping: map[string]string{"you": "me"}}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func structEmb() {
	fmt.Println(Lookup("you"))
}
