package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
//
// Uint32 goSDL_pixelFlag(Uint32 pixelFormat) {
//   return SDL_PIXELFLAG(pixelFormat);
// }
//
// Uint32 goSDL_pixelType(Uint32 pixelFormat) {
//   return SDL_PIXELTYPE(pixelFormat);
// }
//
// Uint32 goSDL_pixelOrder(Uint32 pixelFormat) {
//   return SDL_PIXELORDER(pixelFormat);
// }
//
// Uint32 goSDL_pixelLayout(Uint32 pixelFormat) {
//   return SDL_PIXELLAYOUT(pixelFormat);
// }
//
// Uint32 goSDL_bitsPerPixel(Uint32 pixelFormat) {
//   return SDL_BITSPERPIXEL(pixelFormat);
// }
//
// Uint32 goSDL_bytesPerPixel(Uint32 pixelFormat) {
//   return SDL_BYTESPERPIXEL(pixelFormat);
// }
//
// Uint32 goSDL_isPixelFormatIndexed(Uint32 pixelFormat) {
//   return SDL_ISPIXELFORMAT_INDEXED(pixelFormat);
// }
//
// Uint32 goSDL_isPixelFormatAlpha(Uint32 pixelFormat) {
//   return SDL_ISPIXELFORMAT_ALPHA(pixelFormat);
// }
//
// Uint32 goSDL_isPixelFormatFourCC(Uint32 pixelFormat) {
//   return SDL_ISPIXELFORMAT_FOURCC(pixelFormat);
// }
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

type PixelType uint32

// Pixel types.
const (
	PixelTypeUnknown  PixelType = C.SDL_PIXELTYPE_UNKNOWN
	PixelTypeIndex1   PixelType = C.SDL_PIXELTYPE_INDEX1
	PixelTypeIndex4   PixelType = C.SDL_PIXELTYPE_INDEX4
	PixelTypeIndex8   PixelType = C.SDL_PIXELTYPE_INDEX8
	PixelTypePacked8  PixelType = C.SDL_PIXELTYPE_PACKED8
	PixelTypePacked16 PixelType = C.SDL_PIXELTYPE_PACKED16
	PixelTypePacked32 PixelType = C.SDL_PIXELTYPE_PACKED32
	PixelTypeArrayU8  PixelType = C.SDL_PIXELTYPE_ARRAYU8
	PixelTypeArrayU16 PixelType = C.SDL_PIXELTYPE_ARRAYU16
	PixelTypeArrayU32 PixelType = C.SDL_PIXELTYPE_ARRAYU32
	PixelTypeArrayF16 PixelType = C.SDL_PIXELTYPE_ARRAYF16
	PixelTypeArrayF32 PixelType = C.SDL_PIXELTYPE_ARRAYF32
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

type PixelLayout uint32

// Packed component layout.
const (
	PackedLayoutNone PixelLayout = iota
	PackedLayout332
	PackedLayout4444
	PackedLayout1555
	PackedLayout5551
	PackedLayout565
	PackedLayout8888
	PackedLayout2101010
	PackedLayout1010102
)

// PixelFormatEnum describes the method of storing pixel data.
type PixelFormatEnum uint32

func (pf PixelFormatEnum) PixelType() PixelType {
	return PixelType(C.goSDL_pixelType(C.Uint32(pf)))
}

func (pf PixelFormatEnum) PixelOrder() PixelOrder {
	return PixelOrder(C.goSDL_pixelOrder(C.Uint32(pf)))
}

func (pf PixelFormatEnum) PixelLayout() PixelLayout {
	return PixelLayout(C.goSDL_pixelLayout(C.Uint32(pf)))
}

// BitsPerPixel returns the number of significant bits in a pixel value
// stored in this format.
func (pf PixelFormatEnum) BitsPerPixel() int {
	return int(C.goSDL_bitsPerPixel(C.Uint32(pf)))
}

// BytesPerPixel returns the number of bytes required to hold a
// pixel value stored in this format.
func (pf PixelFormatEnum) BytesPerPixel() int {
	return int(C.goSDL_bytesPerPixel(C.Uint32(pf)))
}

// String returns the SDL constant name of the pixel format.
func (pf PixelFormatEnum) String() string {
	return C.GoString(C.SDL_GetPixelFormatName(C.Uint32(pf)))
}

func (pf PixelFormatEnum) IsIndexed() bool {
	return C.goSDL_isPixelFormatIndexed(C.Uint32(pf)) != 0
}

func (pf PixelFormatEnum) IsAlpha() bool {
	return C.goSDL_isPixelFormatAlpha(C.Uint32(pf)) != 0
}

func (pf PixelFormatEnum) IsFourCC() bool {
	return C.goSDL_isPixelFormatFourCC(C.Uint32(pf)) != 0
}

// Defined pixel formats.
const (
	PixelFormatUnknown     PixelFormatEnum = C.SDL_PIXELFORMAT_UNKNOWN
	PixelFormatIndex1LSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX1LSB
	PixelFormatIndex1MSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX1MSB
	PixelFormatIndex4LSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX4LSB
	PixelFormatIndex4MSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX4MSB
	PixelFormatIndex8      PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX8
	PixelFormatRGB332      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB332
	PixelFormatRGB444      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB444
	PixelFormatRGB555      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB555
	PixelFormatBGR555      PixelFormatEnum = C.SDL_PIXELFORMAT_BGR555
	PixelFormatARGB4444    PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB4444
	PixelFormatRGBA4444    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBA4444
	PixelFormatABGR4444    PixelFormatEnum = C.SDL_PIXELFORMAT_ABGR4444
	PixelFormatBGRA4444    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRA4444
	PixelFormatARGB1555    PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB1555
	PixelFormatRGBA5551    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBA5551
	PixelFormatABGR1555    PixelFormatEnum = C.SDL_PIXELFORMAT_ABGR1555
	PixelFormatBGRA5551    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRA5551
	PixelFormatRGB565      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB565
	PixelFormatBGR565      PixelFormatEnum = C.SDL_PIXELFORMAT_BGR565
	PixelFormatRGB24       PixelFormatEnum = C.SDL_PIXELFORMAT_RGB24
	PixelFormatBGR24       PixelFormatEnum = C.SDL_PIXELFORMAT_BGR24
	PixelFormatRGB888      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB888
	PixelFormatRGBX8888    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBX8888
	PixelFormatBGR888      PixelFormatEnum = C.SDL_PIXELFORMAT_BGR888
	PixelFormatBGRX8888    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRX8888
	PixelFormatARGB8888    PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB8888
	PixelFormatRGBA8888    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBA8888
	PixelFormatABGR8888    PixelFormatEnum = C.SDL_PIXELFORMAT_ABGR8888
	PixelFormatBGRA8888    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRA8888
	PixelFormatARGB2101010 PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB2101010
	PixelFormatYV12        PixelFormatEnum = C.SDL_PIXELFORMAT_YV12
	PixelFormatIYUV        PixelFormatEnum = C.SDL_PIXELFORMAT_IYUV
	PixelFormatYUY2        PixelFormatEnum = C.SDL_PIXELFORMAT_YUY2
	PixelFormatUYVY        PixelFormatEnum = C.SDL_PIXELFORMAT_UYVY
	PixelFormatYVYU        PixelFormatEnum = C.SDL_PIXELFORMAT_YVYU
)

// END Pixel Info }}}1

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
