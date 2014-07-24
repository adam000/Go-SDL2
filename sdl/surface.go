package sdl

// #include "SDL.h"
import "C"

import (
	"image"
	"image/color"
	"unsafe"
)

// Surface is a rectangular array of pixels.
type Surface struct {
	s C.SDL_Surface
}

// PixelFormat returns the surface's pixel format.
func (surface *Surface) PixelFormat() *PixelFormat {
	return &PixelFormat{
		Format:        PixelFormatEnum(surface.s.format.format),
		BitsPerPixel:  uint8(surface.s.format.BitsPerPixel),
		BytesPerPixel: uint8(surface.s.format.BytesPerPixel),
		Rmask:         uint32(surface.s.format.Rmask),
		Gmask:         uint32(surface.s.format.Gmask),
		Bmask:         uint32(surface.s.format.Bmask),
		Amask:         uint32(surface.s.format.Amask),
	}
}

// Size returns the surface's width and height.
func (surface *Surface) Size() Point {
	return Point{int(surface.s.w), int(surface.s.h)}
}

// PixelData locks the surface and returns a PixelData value,
// which can be used to access and modify the surface's pixels.
// The returned PixelData must be closed before the surface can be
// used again.
func (surface *Surface) PixelData() (PixelData, error) {
	if result := C.SDL_LockSurface(&surface.s); result < 0 {
		return PixelData{}, GetError()
	}
	return PixelData{s: &surface.s}, nil
}

// Free destroys the surface.
func (surface *Surface) Free() {
	C.SDL_FreeSurface(&surface.s)
}

// PixelData is a mutable view of a surface's pixels.  The data is only
// available while a surface is locked, so pixel data should be closed to
// allow the surface to be used again.
//
// PixelData implements the image.Image and draw.Image interfaces.
// See: http://golang.org/pkg/image/#Image and http://golang.org/pkg/image/draw/#Image.
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

// At returns the pixel at the given position.
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

// ColorModel returns the color model of the pixel data.
func (pix PixelData) ColorModel() color.Model {
	// TODO this is a guess
	return color.NRGBAModel
}

// Bounds returns a rectangle of (0,0) => (w,h).
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

// Set sets the color at an x, y position in the PixelData to a given color.
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

// Close unlocks the underlying surface.  pix should not be used after
// calling Close.
func (pix PixelData) Close() error {
	C.SDL_UnlockSurface(pix.s)
	return nil
}
