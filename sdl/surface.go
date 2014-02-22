package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

import (
	"image"
	"image/color"
	"unsafe"
)

type Surface struct {
	s *C.SDL_Surface
}

func (surface Surface) PixelData() (PixelData, error) {
	if result := C.SDL_LockSurface(surface.s); result < 0 {
		return PixelData{}, getError()
	}
	return PixelData{s: surface.s}, nil
}

// SDL_CreateTextureFromSurface
func (surface Surface) ToTexture(renderer Renderer) (Texture, error) {
	txt := C.SDL_CreateTextureFromSurface(renderer.r, surface.s)
	if txt == nil {
		return Texture{}, getError()
	}
	return Texture{txt}, nil
}

func (surface Surface) Free() {
	C.SDL_FreeSurface(surface.s)
}

// Implements the image.Image and draw.Image interfaces
// See: http://golang.org/pkg/image/#Image http://golang.org/pkg/image/draw/#Image
type PixelData struct {
	s *C.SDL_Surface
}

// Expand out a component of a color from a pixel representing the color.
// TODO unit tests
func expandColor(pixel uint32, mask C.Uint32, shift, loss C.Uint8) uint8 {
	temp := pixel & uint32(mask)
	temp = temp >> uint8(shift)
	return uint8(temp << uint8(loss))
}

// Takes a uintptr to the pixel data from an SDL Surface and returns a pointer to the particular
// pixel referenced by parameters.
// Pitch is the number of bytes that make up one horizontal row (same as the field in SDL_Surface)
func getPixelPointer(pixels uintptr, x, y, bytesPerPixel, pitch int) unsafe.Pointer {
	offset := x*bytesPerPixel + y*pitch
	return unsafe.Pointer(pixels + uintptr(offset))
}

func (pix PixelData) At(x, y int) color.Color {
	format := pix.s.format
	bytesPerPixel := int(format.BytesPerPixel)

	ptr := getPixelPointer(uintptr(pix.s.pixels), x, y, bytesPerPixel, int(pix.s.pitch))
	pixel := *(*uint32)(ptr)

	// TODO not necesarily NRGBA (which would be an entirely different codepath)
	var col color.NRGBA

	switch bytesPerPixel {
	case 1:
		// TODO look up the color in color palette
	case 2, 3, 4:
		col.R = expandColor(pixel, format.Rmask, format.Rshift, format.Rloss)
		col.G = expandColor(pixel, format.Gmask, format.Gshift, format.Gloss)
		col.B = expandColor(pixel, format.Bmask, format.Bshift, format.Bloss)
		// If the alpha mask is 0, there's no alpha component, so set it to 1
		if format.Amask == 0 {
			col.A = 1
		} else {
			col.A = expandColor(pixel, format.Amask, format.Ashift, format.Aloss)
		}
	}

	return col
}

func (pix PixelData) ColorModel() color.Model {
	// TODO this is a guess
	return color.NRGBAModel
}

// PixelData doesn't actually know where it is on the screen, so return (0,0) => (w,h)
func (pix PixelData) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, 0}, image.Point{int(pix.s.w), int(pix.s.h)}}
}

// Collapse a component of the color into a pointer at a pixel representing the color.
// TODO unit tests
func collapseColor(pixel *uint32, color uint8, shift, loss C.Uint8) {
	temp := uint32(color >> uint8(loss))
	temp = temp << uint8(shift)
	*pixel = *pixel & temp
}

// Set the color at an x, y position in the PixelData to a given color.
func (pix PixelData) Set(x, y int, c color.Color) {
	format := pix.s.format
	bytesPerPixel := int(format.BytesPerPixel)

	switch bytesPerPixel {
	case 1:
		// TODO look up the color in color palette
	case 2, 3, 4:
		// if necessary, convert color model to NRGBA
		col := pix.ColorModel().Convert(c).(color.NRGBA)

		// put that in a uint32 that I can slap into the void* of pixel data (also helper function?)
		var pixel *uint32
		collapseColor(pixel, col.R, format.Rshift, format.Rloss)
		collapseColor(pixel, col.G, format.Gshift, format.Gloss)
		collapseColor(pixel, col.B, format.Bshift, format.Bloss)
		// If the alpha mask is 0, there's no alpha component, so set it to 1
		if format.Amask == 0 {
			col.A = 1
		}
		collapseColor(pixel, col.A, format.Ashift, format.Aloss)

		// get pixel offset that's (x, y)
		ptr := getPixelPointer(uintptr(pix.s.pixels), x, y, bytesPerPixel, int(pix.s.pitch))

		// slap it in
		*(*uint32)(ptr) = *pixel
	}
}

func (pix PixelData) Close() error {
	C.SDL_UnlockSurface(pix.s)
	return nil
}
