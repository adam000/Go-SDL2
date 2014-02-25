package sdl

// #include "SDL.h"
import "C"

import (
	"fmt"
)

// Point is a two-dimensional point.  The axes increase right and down.
type Point struct {
	X, Y int
}

// Pt is shorthand for Point{x, y}.
func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}

// String returns a string representation of p like "(3, 4)".
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) toCPoint() *C.SDL_Point {
	if p == nil {
		return nil
	}
	return &C.SDL_Point{
		x: C.int(p.X),
		y: C.int(p.Y),
	}
}

// Rectangle is a rectangle, with the origin at the upper left.
type Rectangle struct {
	Origin Point
	W, H   int
}

// Rect is shorthand for Rectangle{Pt(x, y), w, h}
func Rect(x, y, w, h int) Rectangle {
	return Rectangle{Origin: Pt(x, y), W: w, H: h}
}

// String returns a string representation of r like "(3, 4) 5x7".
func (r Rectangle) String() string {
	return fmt.Sprintf("%v %dx%d", r.Origin, r.W, r.H)
}

func (r *Rectangle) toCRect() *C.SDL_Rect {
	if r == nil {
		return nil
	}
	return &C.SDL_Rect{
		x: C.int(r.Origin.X),
		y: C.int(r.Origin.Y),
		w: C.int(r.W),
		h: C.int(r.H),
	}
}
