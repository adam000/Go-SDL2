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

// BEGIN Pixel Info {{{1

// The following helpers are modified directly from SDL 2.0 source code:
// include/SDL_pixels.h

type PixelType uint32

const (
	PIXELTYPE_UNKNOWN PixelType = iota
	PIXELTYPE_INDEX1
	PIXELTYPE_INDEX4
	PIXELTYPE_INDEX8
	PIXELTYPE_PACKED8
	PIXELTYPE_PACKED16
	PIXELTYPE_PACKED32
	PIXELTYPE_ARRAYU8
	PIXELTYPE_ARRAYU16
	PIXELTYPE_ARRAYU32
	PIXELTYPE_ARRAYF16
	PIXELTYPE_ARRAYF32
)

type PixelOrder uint32

// Bitmap pixel order, high bit -> low bit.
const (
	BITMAPORDER_NONE PixelOrder = iota
	BITMAPORDER_4321
	BITMAPORDER_1234
)

// Packed component order, high bit -> low bit.
const (
	PACKEDORDER_NONE PixelOrder = iota
	PACKEDORDER_XRGB
	PACKEDORDER_RGBX
	PACKEDORDER_ARGB
	PACKEDORDER_RGBA
	PACKEDORDER_XBGR
	PACKEDORDER_BGRX
	PACKEDORDER_ABGR
	PACKEDORDER_BGRA
)

// Array component order, low byte -> high byte.
const (
	ARRAYORDER_NONE PixelOrder = iota
	ARRAYORDER_RGB
	ARRAYORDER_RGBA
	ARRAYORDER_ARGB
	ARRAYORDER_BGR
	ARRAYORDER_BGRA
	ARRAYORDER_ABGR
)

type PackedLayout uint32

// Packed component layout.
const (
	PACKEDLAYOUT_NONE PackedLayout = iota
	PACKEDLAYOUT_332
	PACKEDLAYOUT_4444
	PACKEDLAYOUT_1555
	PACKEDLAYOUT_5551
	PACKEDLAYOUT_565
	PACKEDLAYOUT_8888
	PACKEDLAYOUT_2101010
	PACKEDLAYOUT_1010102
)

func DEFINE_PIXELFOURCC(A, B, C, D uint32) PixelFormat {
	return PixelFormat((A << 0) | (B << 8) | (C << 16) | (D << 24))
}

func DEFINE_PIXELFORMAT(typ PixelType, order PixelOrder, layout PackedLayout, bits, bytes uint32) PixelFormat {
	return PixelFormat((1 << 28) | uint32(typ<<24) | uint32(order<<20) | uint32(layout<<16) | uint32(bits<<8) | uint32(bytes<<0))
}

func PIXELFLAG(X PixelFormat) PixelFormat {
	return (X >> 28) & 0x0F
}

func PIXELTYPE(X PixelFormat) PixelType {
	return PixelType((X >> 24) & 0x0F)
}

func PIXELORDER(X PixelFormat) PixelOrder {
	return PixelOrder((X >> 20) & 0x0F)
}

func PIXELLAYOUT(X PixelFormat) PackedLayout {
	return PackedLayout((X >> 16) & 0x0F)
}

func BITSPERPIXEL(X PixelFormat) uint32 {
	return uint32(X>>8) & 0xFF
}

func BYTESPERPIXEL(X PixelFormat) PixelFormat {
	if ISPIXELFORMAT_FOURCC(X) {
		if X == PIXELFORMAT_YUY2 || X == PIXELFORMAT_UYVY || X == PIXELFORMAT_YVYU {
			return 2
		}
		return 1
	}
	return (X >> 0) & 0xFF
}

func ISPIXELFORMAT_INDEXED(format PixelFormat) bool {
	pixelType := PIXELTYPE(format)
	return !ISPIXELFORMAT_FOURCC(format) &&
		((pixelType == PIXELTYPE_INDEX1) ||
			(pixelType == PIXELTYPE_INDEX4) ||
			(pixelType == PIXELTYPE_INDEX8))
}

func ISPIXELFORMAT_ALPHA(format PixelFormat) bool {
	pixelOrder := PIXELORDER(format)
	return !ISPIXELFORMAT_FOURCC(format) &&
		((pixelOrder == PACKEDORDER_ARGB) ||
			(pixelOrder == PACKEDORDER_RGBA) ||
			(pixelOrder == PACKEDORDER_ABGR) ||
			(pixelOrder == PACKEDORDER_BGRA))
}

// The flag is set to 1 because 0x1? is not in the printable ASCII range
func ISPIXELFORMAT_FOURCC(format PixelFormat) bool {
	return format != 0 && PIXELFLAG(format) != 1
}

type PixelFormat uint32

// TODO constantize this on the file it's based on; I can use macros!
// Original Note: If you modify this list, update SDL_GetPixelFormatName()
// Note: this is backed by tests that verify the value the same way SDL does. If the SDL macros are
// modified, the corresponding test here will need to be updated and run
const (
	PIXELFORMAT_UNKNOWN     PixelFormat = 0
	PIXELFORMAT_INDEX1LSB               = 286261504
	PIXELFORMAT_INDEX1MSB               = 287310080
	PIXELFORMAT_INDEX4LSB               = 303039488
	PIXELFORMAT_INDEX4MSB               = 304088064
	PIXELFORMAT_INDEX8                  = 318769153
	PIXELFORMAT_RGB332                  = 336660481
	PIXELFORMAT_RGB444                  = 353504258
	PIXELFORMAT_RGB555                  = 353570562
	PIXELFORMAT_BGR555                  = 357764866
	PIXELFORMAT_ARGB4444                = 355602434
	PIXELFORMAT_RGBA4444                = 356651010
	PIXELFORMAT_ABGR4444                = 359796738
	PIXELFORMAT_BGRA4444                = 360845314
	PIXELFORMAT_ARGB1555                = 355667970
	PIXELFORMAT_RGBA5551                = 356782082
	PIXELFORMAT_ABGR1555                = 359862274
	PIXELFORMAT_BGRA5551                = 360976386
	PIXELFORMAT_RGB565                  = 353701890
	PIXELFORMAT_BGR565                  = 357896194
	PIXELFORMAT_RGB24                   = 386930691
	PIXELFORMAT_BGR24                   = 390076419
	PIXELFORMAT_RGB888                  = 370546692
	PIXELFORMAT_RGBX8888                = 371595268
	PIXELFORMAT_BGR888                  = 374740996
	PIXELFORMAT_BGRX8888                = 375789572
	PIXELFORMAT_ARGB8888                = 372645892
	PIXELFORMAT_RGBA8888                = 373694468
	PIXELFORMAT_ABGR8888                = 376840196
	PIXELFORMAT_BGRA8888                = 377888772
	PIXELFORMAT_ARGB2101010             = 372711428
	PIXELFORMAT_YV12                    = 842094169
	PIXELFORMAT_IYUV                    = 1448433993
	PIXELFORMAT_YUY2                    = 844715353
	PIXELFORMAT_UYVY                    = 1498831189
	PIXELFORMAT_YVYU                    = 1431918169
)

// END Pixel Info }}}1

func (surface Surface) PixelData() (PixelData, error) {
	if result := C.SDL_LockSurface(surface.s); result < 0 {
		return PixelData{}, getError()
	}
	return PixelData{s: surface.s}, nil
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
	col := color.NRGBA{}

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
