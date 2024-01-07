package main

import (
	"image"
)

func FindLinks(url string) ([]string, error) {
	return nil, nil // or a multi-valued function call that returns a tuple of these two types
}

// Well-chosen names can document the significance of a function's results.
// Names are particularly valuable when a function returns multiple results of the same type, like
func Size(rect image.Rectangle) (width, height int) {
	return 0, 0
}

// In a function with named results, the operands of a return statement may be omitted. This is called a bare return.
func CountWordsAndImages(url string) (words, images int, err error) {
	words, images = cWaI("")
	err = nil
	return // we did a bare return because we already assigned the named results
}

func cWaI(n string) (words, images int) {
	return 0, 0
}
