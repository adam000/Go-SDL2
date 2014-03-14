package sdl

// #include "SDL.h"
import "C"

import (
	"fmt"
)

// MouseButton is an enumeration of mouse buttons.
type MouseButton uint8

// Common mouse buttons
const (
	LeftMouseButton   MouseButton = C.SDL_BUTTON_LEFT
	MiddleMouseButton MouseButton = C.SDL_BUTTON_MIDDLE
	RightMouseButton  MouseButton = C.SDL_BUTTON_RIGHT
	X1MouseButton     MouseButton = C.SDL_BUTTON_X1
	X2MouseButton     MouseButton = C.SDL_BUTTON_X2
)

// String returns the button's name like "LeftMouseButton".
func (mb MouseButton) String() string {
	switch mb {
	case LeftMouseButton:
		return "LeftMouseButton"
	case MiddleMouseButton:
		return "MiddleMouseButton"
	case RightMouseButton:
		return "RightMouseButton"
	case X1MouseButton:
		return "X1MouseButton"
	case X2MouseButton:
		return "X2MouseButton"
	default:
		return fmt.Sprintf("MouseButton(%d)", uint8(mb))
	}
}

// Mask returns the bitmask for checking a mouse state.
func (mb MouseButton) Mask() uint32 {
	return 1 << (uint32(mb) - 1)
}

// GetMouseState returns the current mouse (x, y) position
// relative to the focus window and buttons pressed.
func GetMouseState() (x, y int32, buttonMask uint32) {
	var cx, cy C.int
	state := C.SDL_GetMouseState(&cx, &cy)
	return int32(cx), int32(cy), uint32(state)
}
