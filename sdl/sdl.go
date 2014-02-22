// Package sdl provides a binding of SDL2 and SDL2_image with an object-oriented twist.
package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

// An InitFlag represents a set of SDL subsystems to initialize.
type InitFlag uint32

// InitFlag masks.
const (
	InitTimer          InitFlag = C.SDL_INIT_TIMER
	InitAudio          InitFlag = C.SDL_INIT_AUDIO
	InitVideo          InitFlag = C.SDL_INIT_VIDEO    // InitVideo implies InitEvents
	InitJoystick       InitFlag = C.SDL_INIT_JOYSTICK // InitJoystick implies InitEvents
	InitHaptic         InitFlag = C.SDL_INIT_HAPTIC
	InitGameController InitFlag = C.SDL_INIT_GAMECONTROLLER // InitGameController implies InitJoystick
	InitEvents         InitFlag = C.SDL_INIT_EVENTS
	InitNoParachute    InitFlag = C.SDL_INIT_NOPARACHUTE // Don't catch fatal signals

	InitEverything InitFlag = C.SDL_INIT_EVERYTHING
)

// Init initializes SDL and its subsystems.  Multiple flags will be ORed together.
// Init must be called before calling other functions in this package.
func Init(flags ...InitFlag) error {
	var f InitFlag
	for i := range flags {
		f |= flags[i]
	}
	if C.SDL_Init(C.Uint32(f)) != 0 {
		return getError()
	}
	return nil
}

// Quit cleans up SDL.
func Quit() {
	C.SDL_Quit()
}

// Error stores an SDL error.
type Error string

func (e Error) Error() string {
	return "sdl: " + string(e)
}

// GetError returns the current SDL error as a Go error value.
// This is internal to SDL but exported because it is cross-package.
func GetError() error {
	// TODO(light): synchronize access?
	e := C.SDL_GetError()
	if *e == 0 {
		// empty string, no error.
		return nil
	}
	return Error(C.GoString(e))
}
