/*
A binding of SDL2 and SDL2_image with an object-oriented twist
*/

package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

// General

type InitFlag uint32

const (
	InitTimer          InitFlag = 0x00000001
	InitAudio                   = 0x00000010
	InitVideo                   = 0x00000020 // InitVideo implies InitEvents
	InitJoystick                = 0x00000200 // InitJoystick implies InitEvents
	InitHaptic                  = 0x00001000
	InitGameController          = 0x00002000 // InitGameController implies InitJoystick
	InitEvents                  = 0x00004000
	InitNoParachute             = 0x00100000 // Don't catch fatal signals
	InitEverything              = InitTimer | InitAudio | InitVideo | InitJoystick | InitHaptic |
		InitGameController | InitEvents | InitNoParachute
)

// Initialize SDL and subsystems
func Init(flags InitFlag) int {
	return int(C.SDL_Init(C.Uint32(flags)))
}

// Clean up for SDL
func Quit() {
	C.SDL_Quit()
}

// Error Handling

type Error string

func (e Error) Error() string {
	return string(e)
}

func getError() error {
	// TODO(light): check for empty string?
	return Error(C.GoString(C.SDL_GetError()))
}
