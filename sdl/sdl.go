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

// Window Handling

type WindowFlag uint32

const (
	WindowFullscreen WindowFlag = 1 << iota
	WindowOpenGL
	WindowShown
	WindowHidden
	WindowBorderless
	WindowResizable
	WindowMinimized
	WindowMaximized
	WindowInputGrabbed
	WindowInputFocus
	WindowMouseFocus
	WindowForeign
	_
	WindowAllowHighDpi
	WindowFullscreenDesktop = 0x00001001
)

const WindowPosUndefined = 0x1FFF0000
const WindowPosCentered = 0x2FFF0000

type Window struct {
	w *C.SDL_Window
}

// Create a new window
func CreateWindow(title string, x, y, w, h int, flags WindowFlag) (Window, error) {
	if window := C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h),
		C.Uint32(flags)); window != nil {
		return Window{window}, nil
	}

	return Window{}, getError()
}

// Get the window's surface
func (w *Window) GetSurface() Surface {
	return Surface{C.SDL_GetWindowSurface(w.w)}
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
