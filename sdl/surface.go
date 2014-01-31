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
	PixelTypeUnknown PixelType = iota
	PixelTypeIndex1
	PixelTypeIndex4
	PixelTypeIndex8
	PixelTypePacked8
	PixelTypePacked16
	PixelTypePacked32
	PixelTypeArrayU8
	PixelTypeArrayU16
	PixelTypeArrayU32
	PixelTypeArrayF16
	PixelTypeArrayF32
)

type PixelOrder uint32

// Bitmap pixel order, high bit -> low bit.
const (
	BitmapOrderNone PixelOrder = iota
	BitmapOrder4321
	BitmapOrder1234
)

// Packed component order, high bit -> low bit.
const (
	PackedOrderNone PixelOrder = iota
	PackedOrderXRGB
	PackedOrderRGBX
	PackedOrderARGB
	PackedOrderRGBA
	PackedOrderXBGR
	PackedOrderBGRX
	PackedOrderABGR
	PackedOrderBGRA
)

// Array component order, low byte -> high byte.
const (
	ArrayOrderNone PixelOrder = iota
	ArrayOrderRGB
	ArrayOrderRGBA
	ArrayOrderARGB
	ArrayOrderBGR
	ArrayOrderBGRA
	ArrayOrderABGR
)

type PackedLayout uint32

// Packed component layout.
const (
	PackedLayoutNone PackedLayout = iota
	PackedLayout332
	PackedLayout4444
	PackedLayout1555
	PackedLayout5551
	PackedLayout565
	PackedLayout8888
	PackedLayout2101010
	PackedLayout1010102
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
		if X == PixelFormatYUY2 || X == PixelFormatUYVY || X == PixelFormatYVYU {
			return 2
		}
		return 1
	}
	return (X >> 0) & 0xFF
}

func ISPIXELFORMAT_INDEXED(format PixelFormat) bool {
	pixelType := PIXELTYPE(format)
	return !ISPIXELFORMAT_FOURCC(format) &&
		((pixelType == PixelTypeIndex1) ||
			(pixelType == PixelTypeIndex4) ||
			(pixelType == PixelTypeIndex8))
}

func ISPIXELFORMAT_ALPHA(format PixelFormat) bool {
	pixelOrder := PIXELORDER(format)
	return !ISPIXELFORMAT_FOURCC(format) &&
		((pixelOrder == PackedOrderARGB) ||
			(pixelOrder == PackedOrderRGBA) ||
			(pixelOrder == PackedOrderABGR) ||
			(pixelOrder == PackedOrderBGRA))
}

// The flag is set to 1 because 0x1? is not in the printable ASCII range
func ISPIXELFORMAT_FOURCC(format PixelFormat) bool {
	return format != 0 && PIXELFLAG(format) != 1
}

type PixelFormat uint32

const (
	PixelFormatUnknown     PixelFormat = PixelFormat(C.SDL_PIXELFORMAT_UNKNOWN)
	PixelFormatIndex1LSB               = PixelFormat(C.SDL_PIXELFORMAT_INDEX1LSB)
	PixelFormatIndex1MSB               = PixelFormat(C.SDL_PIXELFORMAT_INDEX1MSB)
	PixelFormatIndex4LSB               = PixelFormat(C.SDL_PIXELFORMAT_INDEX4LSB)
	PixelFormatIndex4MSB               = PixelFormat(C.SDL_PIXELFORMAT_INDEX4MSB)
	PixelFormatIndex8                  = PixelFormat(C.SDL_PIXELFORMAT_INDEX8)
	PixelFormatRGB332                  = PixelFormat(C.SDL_PIXELFORMAT_RGB332)
	PixelFormatRGB444                  = PixelFormat(C.SDL_PIXELFORMAT_RGB444)
	PixelFormatRGB555                  = PixelFormat(C.SDL_PIXELFORMAT_RGB555)
	PixelFormatBGR555                  = PixelFormat(C.SDL_PIXELFORMAT_BGR555)
	PixelFormatARGB4444                = PixelFormat(C.SDL_PIXELFORMAT_ARGB4444)
	PixelFormatRGBA4444                = PixelFormat(C.SDL_PIXELFORMAT_RGBA4444)
	PixelFormatABGR4444                = PixelFormat(C.SDL_PIXELFORMAT_ABGR4444)
	PixelFormatBGRA4444                = PixelFormat(C.SDL_PIXELFORMAT_BGRA4444)
	PixelFormatARGB1555                = PixelFormat(C.SDL_PIXELFORMAT_ARGB1555)
	PixelFormatRGBA5551                = PixelFormat(C.SDL_PIXELFORMAT_RGBA5551)
	PixelFormatABGR1555                = PixelFormat(C.SDL_PIXELFORMAT_ABGR1555)
	PixelFormatBGRA5551                = PixelFormat(C.SDL_PIXELFORMAT_BGRA5551)
	PixelFormatRGB565                  = PixelFormat(C.SDL_PIXELFORMAT_RGB565)
	PixelFormatBGR565                  = PixelFormat(C.SDL_PIXELFORMAT_BGR565)
	PixelFormatRGB24                   = PixelFormat(C.SDL_PIXELFORMAT_RGB24)
	PixelFormatBGR24                   = PixelFormat(C.SDL_PIXELFORMAT_BGR24)
	PixelFormatRGB888                  = PixelFormat(C.SDL_PIXELFORMAT_RGB888)
	PixelFormatRGBX8888                = PixelFormat(C.SDL_PIXELFORMAT_RGBX8888)
	PixelFormatBGR888                  = PixelFormat(C.SDL_PIXELFORMAT_BGR888)
	PixelFormatBGRX8888                = PixelFormat(C.SDL_PIXELFORMAT_BGRX8888)
	PixelFormatARGB8888                = PixelFormat(C.SDL_PIXELFORMAT_ARGB8888)
	PixelFormatRGBA8888                = PixelFormat(C.SDL_PIXELFORMAT_RGBA8888)
	PixelFormatABGR8888                = PixelFormat(C.SDL_PIXELFORMAT_ABGR8888)
	PixelFormatBGRA8888                = PixelFormat(C.SDL_PIXELFORMAT_BGRA8888)
	PixelFormatARGB2101010             = PixelFormat(C.SDL_PIXELFORMAT_ARGB2101010)
	PixelFormatYV12                    = PixelFormat(C.SDL_PIXELFORMAT_YV12)
	PixelFormatIYUV                    = PixelFormat(C.SDL_PIXELFORMAT_IYUV)
	PixelFormatYUY2                    = PixelFormat(C.SDL_PIXELFORMAT_YUY2)
	PixelFormatUYVY                    = PixelFormat(C.SDL_PIXELFORMAT_UYVY)
	PixelFormatYVYU                    = PixelFormat(C.SDL_PIXELFORMAT_YVYU)
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
