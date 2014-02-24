package sdl

// #include "SDL.h"
import "C"

type GLContext struct {
	c C.SDL_GLContext
}

// NewGLContext creates a new GLContext in the given window
func NewGLContext(w Window) (GLContext, error) {
	context := C.SDL_GL_CreateContext(w.w)
	if context == nil {
		return GLContext{}, GetError()
	}
	return GLContext{context}, nil
}

// Destroy deletes the GLContext
func (c GLContext) Destroy() {
	C.SDL_GL_DeleteContext(c.c)
}

// GLSwap updates a window with an OpenGL rendering. (Used in double-buffering environments, which
// are the default)
func (w Window) GLSwap() {
	C.SDL_GL_SwapWindow(w.w)
}
