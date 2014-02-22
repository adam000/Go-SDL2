package sdl

// #include "SDL.h"
import "C"

type GLContext struct {
	c C.SDL_GLContext
}

// Create a new GLContext in the given window
func NewGLContext(w Window) (GLContext, error) {
	context := C.SDL_GL_CreateContext(w.w)
	if context == nil {
		return GLContext{}, GetError()
	}
	return GLContext{context}, nil
}

// Destroy the GLContext
func (c GLContext) Destroy() {
	C.SDL_GL_DeleteContext(c.c)
}

// "Swap the buffers" for GL applications
func (w Window) GLSwap() {
	C.SDL_GL_SwapWindow(w.w)
}
