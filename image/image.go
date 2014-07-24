// Package image provides access to popular image formats via SDL_image.
package image

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL_image.h"
import "C"

import (
	"unsafe"

	"github.com/adam000/Go-SDL2/sdl"
)

// An InitFlag represents a set of image libraries to load.
type InitFlag uint32

// Initialization flags.
const (
	InitJPG  InitFlag = C.IMG_INIT_JPG
	InitPNG  InitFlag = C.IMG_INIT_PNG
	InitTIF  InitFlag = C.IMG_INIT_TIF
	InitWEBP InitFlag = C.IMG_INIT_WEBP
)

// Init loads dynamic libraries.  Multiple flags will be ORed together.
// Init returns the libraries succesfully initialized.  It is not
// required to call Init before using other functions in this package.
func Init(flags ...InitFlag) InitFlag {
	var f InitFlag
	for i := range flags {
		f |= flags[i]
	}
	return InitFlag(C.IMG_Init(C.int(f)))
}

// Quit unloads libraries loaded with Init.
func Quit() {
	C.IMG_Quit()
}

// Load reads an image from a file.
func Load(file string) (*sdl.Surface, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	surf := C.IMG_Load(cfile)

	if surf == nil {
		return nil, sdl.GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(surf)), nil
}
