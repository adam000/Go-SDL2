package sdl

// #cgo pkg-config: sdl2
//
// #include "SDL.h"
import "C"

type GLContext struct {
	c C.SDL_GLContext
}

func NewGLContext(w Window) (GLContext, error) {
	context := C.SDL_GL_CreateContext(w.w)
	if context == nil {
		return GLContext{}, GetError()
	}
	return GLContext{context}, nil
}

func (c GLContext) Destroy() {
	C.SDL_GL_DeleteContext(c.c)
}

func (w Window) GLSwap() {
	C.SDL_GL_SwapWindow(w.w)
}
