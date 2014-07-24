package sdl

// #include "SDL.h"
import "C"

import (
	"unsafe"
)

// Texture is an efficient driver-specific representation of pixel data.
type Texture struct {
	t C.SDL_Texture
}

// NewTextureFromSurface creates a new texture from an existing surface.
func NewTextureFromSurface(renderer *Renderer, surface *Surface) (*Texture, error) {
	tex := C.SDL_CreateTextureFromSurface(&renderer.r, &surface.s)
	if tex == nil {
		return nil, GetError()
	}
	return (*Texture)(unsafe.Pointer(tex)), nil
}

// Destroy destroys the texture.
func (t *Texture) Destroy() {
	C.SDL_DestroyTexture(&t.t)
}
