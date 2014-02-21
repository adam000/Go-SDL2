package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
//	#include "SDL_image.h"
import "C"

import "unsafe"

func LoadImage(fileName string) (Surface, error) {
	cName := C.CString(fileName)
	surf := C.IMG_Load(cName)
	C.free(unsafe.Pointer(cName))

	if surf == nil {
		return Surface{}, getError()
	}

	return Surface{surf}, nil
}
