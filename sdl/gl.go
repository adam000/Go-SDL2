package sdl

// #include "SDL.h"
import "C"

// A GLContext is an opaque handle to an OpenGL context.
type GLContext struct {
	c C.SDL_GLContext
}

// NewGLContext creates a new GLContext for use with w and makes it the current context.
func NewGLContext(w *Window) (GLContext, error) {
	context := C.SDL_GL_CreateContext(&w.w)
	if context == nil {
		return GLContext{}, GetError()
	}
	return GLContext{context}, nil
}

// MakeCurrent makes the context the current context and associates it with w.
// w must be a compatible window.
func (ctx GLContext) MakeCurrent(w *Window) error {
	if C.SDL_GL_MakeCurrent(&w.w, ctx.c) != 0 {
		return GetError()
	}
	return nil
}

// Destroy deletes the GLContext
func (ctx GLContext) Destroy() {
	C.SDL_GL_DeleteContext(ctx.c)
}

// GLSwap updates a window with an OpenGL rendering. (Used in double-buffering environments, which
// are the default)
func (w *Window) GLSwap() {
	C.SDL_GL_SwapWindow(&w.w)
}
