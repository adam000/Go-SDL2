package sdl

// #cgo pkg-config: sdl2
//
// #include "SDL.h"
//
// static Uint32 goSDL_pixelFlag(Uint32 pixelFormat) {
//   return SDL_PIXELFLAG(pixelFormat);
// }
//
// static Uint32 goSDL_pixelType(Uint32 pixelFormat) {
//   return SDL_PIXELTYPE(pixelFormat);
// }
//
// static Uint32 goSDL_pixelOrder(Uint32 pixelFormat) {
//   return SDL_PIXELORDER(pixelFormat);
// }
//
// static Uint32 goSDL_pixelLayout(Uint32 pixelFormat) {
//   return SDL_PIXELLAYOUT(pixelFormat);
// }
//
// static Uint32 goSDL_bitsPerPixel(Uint32 pixelFormat) {
//   return SDL_BITSPERPIXEL(pixelFormat);
// }
//
// static Uint32 goSDL_bytesPerPixel(Uint32 pixelFormat) {
//   return SDL_BYTESPERPIXEL(pixelFormat);
// }
//
// static Uint32 goSDL_isPixelFormatIndexed(Uint32 pixelFormat) {
//   return SDL_ISPIXELFORMAT_INDEXED(pixelFormat);
// }
//
// static Uint32 goSDL_isPixelFormatAlpha(Uint32 pixelFormat) {
//   return SDL_ISPIXELFORMAT_ALPHA(pixelFormat);
// }
//
// static Uint32 goSDL_isPixelFormatFourCC(Uint32 pixelFormat) {
//   return SDL_ISPIXELFORMAT_FOURCC(pixelFormat);
// }
import "C"

// PixelFormatEnum describes the method of storing pixel data.
type PixelFormatEnum uint32

// PixelType returns the data type used for the pixel format.
func (pf PixelFormatEnum) PixelType() PixelType {
	return PixelType(C.goSDL_pixelType(C.Uint32(pf)))
}

// PixelOrder returns the ordering of channels in the pixel format.
func (pf PixelFormatEnum) PixelOrder() PixelOrder {
	return PixelOrder(C.goSDL_pixelOrder(C.Uint32(pf)))
}

// PixelLayout returns the layout of channels in a packed pixel format.
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

// IsIndexed reports whether the pixel format has a palette.
func (pf PixelFormatEnum) IsIndexed() bool {
	return C.goSDL_isPixelFormatIndexed(C.Uint32(pf)) != 0
}

// IsAlpha reports whether the pixel format has an alpha channel.
func (pf PixelFormatEnum) IsAlpha() bool {
	return C.goSDL_isPixelFormatAlpha(C.Uint32(pf)) != 0
}

// IsFourCC reports whether the pixel format is a four-character code,
// like YUV.
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

// PixelType is a pixel format's data type.
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

// PixelOrder is a pixel format's channel order.
type PixelOrder uint32

// Bitmap pixel order, high bit -> low bit.
const (
	BitmapOrderNone PixelOrder = C.SDL_BITMAPORDER_NONE
	BitmapOrder4321 PixelOrder = C.SDL_BITMAPORDER_4321
	BitmapOrder1234 PixelOrder = C.SDL_BITMAPORDER_1234
)

// Packed component order, high bit -> low bit.
const (
	PackedOrderNone PixelOrder = C.SDL_PACKEDORDER_NONE
	PackedOrderXRGB PixelOrder = C.SDL_PACKEDORDER_XRGB
	PackedOrderRGBX PixelOrder = C.SDL_PACKEDORDER_RGBX
	PackedOrderARGB PixelOrder = C.SDL_PACKEDORDER_ARGB
	PackedOrderRGBA PixelOrder = C.SDL_PACKEDORDER_RGBA
	PackedOrderXBGR PixelOrder = C.SDL_PACKEDORDER_XBGR
	PackedOrderBGRX PixelOrder = C.SDL_PACKEDORDER_BGRX
	PackedOrderABGR PixelOrder = C.SDL_PACKEDORDER_ABGR
	PackedOrderBGRA PixelOrder = C.SDL_PACKEDORDER_BGRA
)

// Array component order, low byte -> high byte.
const (
	ArrayOrderNone PixelOrder = C.SDL_ARRAYORDER_NONE
	ArrayOrderRGB  PixelOrder = C.SDL_ARRAYORDER_RGB
	ArrayOrderRGBA PixelOrder = C.SDL_ARRAYORDER_RGBA
	ArrayOrderARGB PixelOrder = C.SDL_ARRAYORDER_ARGB
	ArrayOrderBGR  PixelOrder = C.SDL_ARRAYORDER_BGR
	ArrayOrderBGRA PixelOrder = C.SDL_ARRAYORDER_BGRA
	ArrayOrderABGR PixelOrder = C.SDL_ARRAYORDER_ABGR
)

// PixelLayout is a packed pixel format's channel bit layout.
type PixelLayout uint32

// Packed channel layouts.
const (
	PackedLayoutNone    PixelLayout = C.SDL_PACKEDLAYOUT_NONE
	PackedLayout332     PixelLayout = C.SDL_PACKEDLAYOUT_332
	PackedLayout4444    PixelLayout = C.SDL_PACKEDLAYOUT_4444
	PackedLayout1555    PixelLayout = C.SDL_PACKEDLAYOUT_1555
	PackedLayout5551    PixelLayout = C.SDL_PACKEDLAYOUT_5551
	PackedLayout565     PixelLayout = C.SDL_PACKEDLAYOUT_565
	PackedLayout8888    PixelLayout = C.SDL_PACKEDLAYOUT_8888
	PackedLayout2101010 PixelLayout = C.SDL_PACKEDLAYOUT_2101010
	PackedLayout1010102 PixelLayout = C.SDL_PACKEDLAYOUT_1010102
)
