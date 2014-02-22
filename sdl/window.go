package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

import (
	"reflect"
	"unsafe"
)

//////////////////////////////////////////
// Contains Window and Renderer methods /
////////////////////////////////////////

// Window Handling

type WindowFlag uint32

const (
	WindowFullscreen WindowFlag = 1 << iota
	WindowOpenGL
	WindowShown
	WindowHidden
	WindowBorderless
	WindowResizable
	WindowMinimized
	WindowMaximized
	WindowInputGrabbed
	WindowInputFocus
	WindowMouseFocus
	WindowForeign
	_
	WindowAllowHighDpi
	WindowFullscreenDesktop = 0x00001001
)

const WindowPosUndefined = 0x1FFF0000
const WindowPosCentered = 0x2FFF0000

type Renderer struct {
	r *C.SDL_Renderer
}

type RendererFlags uint32

const (
	RendererSoftware RendererFlags = 1 << iota
	RendererAccelerated
	RendererPresentVsync
	RendererTargetTexture
)

const TextureFormatsSize int = 16

type RendererInfo struct {
	Name              string
	Flags             RendererFlags
	NumTextureFormats uint32
	TextureFormats    []PixelFormatEnum
	MaxTextureWidth   int
	MaxTextureHeight  int
}

type Window struct {
	w *C.SDL_Window
	r *Renderer
}

// Create a new window
func NewWindow(title string, x, y, w, h int, flags WindowFlag) (Window, error) {
	if window := C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h),
		C.Uint32(flags)); window != nil {

		return Window{window, nil}, nil
	}

	return Window{}, getError()
}

// Get the window's surface
func (w Window) GetSurface() Surface {
	return Surface{C.SDL_GetWindowSurface(w.w)}
}

func (w Window) Destroy() {
	C.SDL_DestroyWindow(w.w)
}

// Renderer functions

func NewRenderer(window Window, index int, flags uint32) (Renderer, error) {
	if r := C.SDL_CreateRenderer(window.w, C.int(index), C.Uint32(flags)); r != nil {
		return Renderer{r}, nil
	}

	return Renderer{}, getError()
}

func (w Window) GetRenderer() (Renderer, error) {
	if r := C.SDL_GetRenderer(w.w); r != nil {
		return Renderer{r}, nil
	}

	return Renderer{}, getError()
}

func (r Renderer) GetRendererInfo() (i *RendererInfo, e error) {
	var info C.SDL_RendererInfo
	if retCode := C.SDL_GetRendererInfo(r.r, &info); retCode != 0 {
		return &RendererInfo{}, getError()
	}

	var textureFormats []PixelFormatEnum
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&textureFormats)))
	sliceHeader.Cap = TextureFormatsSize
	sliceHeader.Len = TextureFormatsSize
	sliceHeader.Data = uintptr(unsafe.Pointer(&info.texture_formats[0]))

	return &RendererInfo{C.GoString(info.name), RendererFlags(info.flags),
		uint32(info.num_texture_formats), textureFormats, int(info.max_texture_width),
		int(info.max_texture_height)}, nil
}

func (r Renderer) Destroy() {
	C.SDL_DestroyRenderer(r.r)
}

// SDL_RenderCopy
func (r Renderer) CopyTexture(texture Texture, srcRect *Rect, destRect *Rect) error {
	var src *C.SDL_Rect
	var dest *C.SDL_Rect
	if srcRect != nil {
		src = srcRect.toCRect()
	}
	if destRect != nil {
		dest = destRect.toCRect()
	}

	if C.SDL_RenderCopy(r.r, texture.t, src, dest) != 0 {
		return getError()
	}
	return nil
}

// SDL_RenderClear
func (r Renderer) Clear() error {
	if retCode := C.SDL_RenderClear(r.r); retCode != 0 {
		return getError()
	}
	return nil
}

func (r Renderer) Present() {
	C.SDL_RenderPresent(r.r)
}
