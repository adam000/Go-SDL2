package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

// Represents the SDL_Rect type
type Rect struct {
	x int
	y int
	w int
	h int
}

func (r Rect) toCRect() *C.SDL_Rect {
	return &C.SDL_Rect{C.int(r.x), C.int(r.y), C.int(r.w), C.int(r.h)}
}
