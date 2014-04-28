package sdl

// #include "SDL.h"
import "C"

// JoystickID is a transient joystick ID.
type JoystickID int32

// HatPosition is a joystick hat position.  Cardinal directions are
// OR'd together to describe diagonals.
type HatPosition uint8

// Hat positions.
const (
	HatCentered HatPosition = C.SDL_HAT_CENTERED

	HatUp    HatPosition = C.SDL_HAT_UP
	HatRight HatPosition = C.SDL_HAT_RIGHT
	HatDown  HatPosition = C.SDL_HAT_DOWN
	HatLeft  HatPosition = C.SDL_HAT_LEFT

	// OR'd combinations of the above.
	HatRightUp   HatPosition = C.SDL_HAT_RIGHTUP
	HatRightDown HatPosition = C.SDL_HAT_RIGHTDOWN
	HatLeftUp    HatPosition = C.SDL_HAT_LEFTUP
	HatLeftDown  HatPosition = C.SDL_HAT_LEFTDOWN
)
