package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

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

type renderer struct {
	r *C.SDL_Renderer
}

type Window struct {
	w        *C.SDL_Window
	Renderer *renderer
}

// Create a new window
func NewWindow(title string, x, y, w, h int, flags WindowFlag) (Window, error) {
	if window := C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h),
		C.Uint32(flags)); window != nil {

		return Window{window, nil}, nil
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

// Renderer stuff

func NewRenderer(window *Window, index int, flags uint32) (renderer, error) {
	if r := C.SDL_CreateRenderer(window.w, C.int(index), C.Uint32(flags)); r != nil {
		return renderer{r}, nil
	}

	return renderer{}, getError()
}

func (w *Window) GetRenderer() (renderer, error) {
	if r := C.SDL_GetRenderer(w.w); r != nil {
		return renderer{r}, nil
	}

	return renderer{}, getError()
}
