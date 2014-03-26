package sdl

// #include "SDL.h"
import "C"

import (
	"github.com/adam000/Go-SDL2/sdl/keys"
)

// KeySym holds the keyboard information from a keyboard event.
type KeySym struct {
	ScanCode int32 // TODO(light): add type
	KeyCode  int32 // TODO(light): add type
	Mod      keys.Mod
}
