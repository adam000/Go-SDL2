package sdl

// #include "SDL.h"
import "C"

type Texture struct {
	t C.SDL_Texture
}

func (t *Texture) Destroy() {
	C.SDL_DestroyTexture(&t.t)
}
