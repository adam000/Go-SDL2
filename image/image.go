// Package image provides access to popular image formats via SDL_image.
package image

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL_image.h"
import "C"

import "unsafe"

func Init() {
}

func LoadImage(fileName string) (Surface, error) {
	cName := C.CString(fileName)
	surf := C.IMG_Load(cName)
	C.free(unsafe.Pointer(cName))

	if surf == nil {
		return Surface{}, getError()
	}

	return Surface{surf}, nil
}
