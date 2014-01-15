/*
A binding of SDL2 and SDL2_image with an object-oriented twist
*/

package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

import (
	"unsafe"
)

// General

type InitFlag uint32

const (
	INIT_TIMER          InitFlag = 0x00000001
	INIT_AUDIO                   = 0x00000010
	INIT_VIDEO                   = 0x00000020 // INIT_VIDEO implies INIT_EVENTS 
	INIT_JOYSTICK                = 0x00000200 // INIT_JOYSTICK implies INIT_EVENTS 
	INIT_HAPTIC                  = 0x00001000
	INIT_GAMECONTROLLER          = 0x00002000 // INIT_GAMECONTROLLER implies INIT_JOYSTICK
	INIT_EVENTS                  = 0x00004000
	INIT_NOPARACHUTE             = 0x00100000 // Don't catch fatal signals
	INIT_EVERYTHING              = 0x00107231 // This should be the sum of all the other flags
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

func GetError() error {
	// TODO(light): check for empty string?
	return Error(C.GoString(C.SDL_GetError()))
}

// Window Handling

type WindowFlag uint32

const (
	WINDOW_FULLSCREEN WindowFlag = 1 << iota
	WINDOW_OPENGL
	WINDOW_SHOWN
	WINDOW_HIDDEN
	WINDOW_BORDERLESS
	WINDOW_RESIZABLE
	WINDOW_MINIMIZED
	WINDOW_MAXIMIZED
	WINDOW_INPUT_GRABBED
	WINDOW_INPUT_FOCUS
	WINDOW_MOUSE_FOCUS
	WINDOW_FOREIGN
	_
	WINDOW_ALLOW_HIGHDPI
	WINDOW_FULLSCREEN_DESKTOP = 0x00001001
)

const WINDOWPOS_UNDEFINED = 0x1FFF0000
const WINDOWPOS_CENTERED =   0x2FFF0000

type Window struct {
	w *C.SDL_Window
}

// Create a new window
func CreateWindow(title string, x, y, w, h int, flags WindowFlag) (Window, error) {
	if window := C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h),
		C.Uint32(flags)); window != nil {
		return Window{window}, nil
	}

	return Window{}, GetError()
}

// Get the window's surface
func (w *Window) GetSurface() *Surface {
	return (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(w.w)))
}

func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.w)
}

// Renderer

type Renderer struct {
	r *C.SDL_Renderer
}

// Create a new renderer
func CreateRenderer(w Window, index int, flags uint32) *Renderer {
	renderer := C.SDL_CreateRenderer(w.w, C.int(index), C.Uint32(flags))

	return (*Renderer)(unsafe.Pointer(renderer))
}
