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

// Destroy destroys the surface.  The surface should not be used after
// a call to Destroy.
func (surface *Surface) Destroy() {
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
// TODO(adam): unit tests
func expandColor(pixel uint32, mask C.Uint32, shift, loss C.Uint8) uint8 {
	temp := pixel & uint32(mask)
	temp = temp >> uint8(shift)
	return uint8(temp << uint8(loss))
}

// pixel returns the address of the pixel at (x, y).
func (pix PixelData) pixel(x, y int) unsafe.Pointer {
	offset := uintptr(y)*uintptr(pix.s.pitch) + uintptr(x)*uintptr(pix.s.format.BytesPerPixel)
	return unsafe.Pointer(uintptr(pix.s.pixels) + offset)
}

// At returns the pixel at the given position.
func (pix PixelData) At(x, y int) color.Color {
	format := pix.s.format
	ptr := pix.pixel(x, y)
	// TODO(adam): not necesarily NRGBA (which would be an entirely different codepath)
	var col color.NRGBA

	switch format.BytesPerPixel {
	case 4:
		pixel := *(*uint32)(ptr)
		col.R = expandColor(pixel, format.Rmask, format.Rshift, format.Rloss)
		col.G = expandColor(pixel, format.Gmask, format.Gshift, format.Gloss)
		col.B = expandColor(pixel, format.Bmask, format.Bshift, format.Bloss)
		// If the alpha mask is 0, there's no alpha component, so set it opaque.
		if format.Amask == 0 {
			col.A = ^uint8(0)
		} else {
			col.A = expandColor(pixel, format.Amask, format.Ashift, format.Aloss)
		}
	default:
		// TODO(#22): handle all pixel formats
		panic("pixel format not handled")
	}
	return col
}

// ColorModel returns the color model of the pixel data.
func (pix PixelData) ColorModel() color.Model {
	// TODO(adam): this is a guess
	return color.NRGBAModel
}

// Bounds returns a rectangle of (0,0) => (w,h).
func (pix PixelData) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(pix.s.w), int(pix.s.h))
}

// collapseColor collapses a component of the color into an OR'able mask.
// TODO(adam): unit tests
func collapseColor(color uint8, shift, loss C.Uint8) uint32 {
	return uint32(color) >> loss << shift
}

// Set sets the color at an x, y position in the PixelData to a given color.
func (pix PixelData) Set(x, y int, c color.Color) {
	format := pix.s.format
	switch format.BytesPerPixel {
	case 4:
		col := pix.ColorModel().Convert(c).(color.NRGBA)

		p := (*uint32)(pix.pixel(x, y))
		*p = collapseColor(col.R, format.Rshift, format.Rloss)
		*p |= collapseColor(col.G, format.Gshift, format.Gloss)
		*p |= collapseColor(col.B, format.Bshift, format.Bloss)
		if format.Amask == 0 {
			// If the alpha mask is 0, there's no alpha component, so set it opaque.
			col.A = ^uint8(0)
		}
		*p |= collapseColor(col.A, format.Ashift, format.Aloss)
	default:
		// TODO(#22): handle all pixel formats
		panic("pixel format not handled")
	}
}

// Destroy unlocks the underlying surface.  pix should not be used after
// calling Destroy.
func (pix PixelData) Destroy() {
	C.SDL_UnlockSurface(pix.s)
}
