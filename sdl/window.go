package sdl

// #include "SDL.h"
import "C"

import (
	"unsafe"
)

// WindowFlag is a window creation option.
type WindowFlag uint32

// Window creation options.
const (
	WindowFullscreen        WindowFlag = C.SDL_WINDOW_FULLSCREEN
	WindowOpenGL            WindowFlag = C.SDL_WINDOW_OPENGL
	WindowShown             WindowFlag = C.SDL_WINDOW_SHOWN
	WindowHidden            WindowFlag = C.SDL_WINDOW_HIDDEN
	WindowBorderless        WindowFlag = C.SDL_WINDOW_BORDERLESS
	WindowResizable         WindowFlag = C.SDL_WINDOW_RESIZABLE
	WindowMinimized         WindowFlag = C.SDL_WINDOW_MINIMIZED
	WindowMaximized         WindowFlag = C.SDL_WINDOW_MAXIMIZED
	WindowInputGrabbed      WindowFlag = C.SDL_WINDOW_INPUT_GRABBED
	WindowInputFocus        WindowFlag = C.SDL_WINDOW_INPUT_FOCUS
	WindowMouseFocus        WindowFlag = C.SDL_WINDOW_MOUSE_FOCUS
	WindowForeign           WindowFlag = C.SDL_WINDOW_FOREIGN
	WindowAllowHighDPI      WindowFlag = C.SDL_WINDOW_ALLOW_HIGHDPI
	WindowFullscreenDesktop WindowFlag = C.SDL_WINDOW_FULLSCREEN_DESKTOP
)

// Window positions.  Used for the origin in NewWindow.
const (
	WindowPosUndefined = C.SDL_WINDOWPOS_UNDEFINED
	WindowPosCentered  = C.SDL_WINDOWPOS_CENTERED
)

// RendererInfo describes the capabilities of a Renderer.
type RendererInfo struct {
	Name             string
	Flags            RendererFlag
	TextureFormats   []PixelFormatEnum
	MaxTextureWidth  int
	MaxTextureHeight int
}

// Window is a window in a GUI environment.
type Window struct {
	w C.SDL_Window
}

// NewWindow creates a new window.  r.x or r.y may also be
// WindowPosCentered or WindowPosUndefined.  Multiple flags will be
// ORed together.
func NewWindow(title string, r Rectangle, flags ...WindowFlag) (*Window, error) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	var f WindowFlag
	for i := range flags {
		f |= flags[i]
	}
	w := C.SDL_CreateWindow(
		ctitle,
		C.int(r.Origin.X),
		C.int(r.Origin.Y),
		C.int(r.W),
		C.int(r.H),
		C.Uint32(f))
	if w == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(w)), nil
}

// Surface returns the window's surface.
func (w *Window) Surface() *Surface {
	return (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(&w.w)))
}

// CreateRenderer creates a 2D rendering context for a window.
// driverIndex is the index of the rendering driver to initialize, or -1
// to initialize the first one that supports the requested
// configuration.  Multiple flags will be ORed together.
func (w *Window) CreateRenderer(driverIndex int, flags ...RendererFlag) (*Renderer, error) {
	var f RendererFlag
	for i := range flags {
		f |= flags[i]
	}
	r := C.SDL_CreateRenderer(&w.w, C.int(driverIndex), C.Uint32(f))
	if r == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(r)), nil
}

// Renderer returns the window's renderer or nil if it doesn't have one.
func (w *Window) Renderer() *Renderer {
	return (*Renderer)(unsafe.Pointer(C.SDL_GetRenderer(&w.w)))
}

// Destroy destroys a window.  It is not safe to use the window after
// calling Destroy.
func (w *Window) Destroy() {
	C.SDL_DestroyWindow(&w.w)
}

// A Renderer represents the rendering state.
type Renderer struct {
	r C.SDL_Renderer
}

// A RendererFlag is an option for creating a renderer.
type RendererFlag uint32

// Renderer creation flags.
const (
	// Software fallback
	RendererSoftware RendererFlag = C.SDL_RENDERER_SOFTWARE
	// Hardware accelerated
	RendererAccelerated RendererFlag = C.SDL_RENDERER_ACCELERATED
	// Present is synchronized with the refresh rate
	RendererPresentVSync RendererFlag = C.SDL_RENDERER_PRESENTVSYNC
	// Render to texture support
	RendererTargetTexture RendererFlag = C.SDL_RENDERER_TARGETTEXTURE
)

// Info returns the renderer's capabilities.
func (r *Renderer) Info() (*RendererInfo, error) {
	var info C.SDL_RendererInfo
	if C.SDL_GetRendererInfo(&r.r, &info) != 0 {
		return nil, GetError()
	}

	formats := make([]PixelFormatEnum, info.num_texture_formats)
	for i := range formats {
		formats[i] = PixelFormatEnum(info.texture_formats[i])
	}
	return &RendererInfo{
		Name:             C.GoString(info.name),
		Flags:            RendererFlag(info.flags),
		TextureFormats:   formats,
		MaxTextureWidth:  int(info.max_texture_width),
		MaxTextureHeight: int(info.max_texture_height),
	}, nil
}

// CopyTexture copies a portion of the texture to the current rendering context.
func (r *Renderer) CopyTexture(texture *Texture, srcRect, destRect *Rectangle) error {
	if C.SDL_RenderCopy(&r.r, &texture.t, srcRect.toCRect(), destRect.toCRect()) != 0 {
		return GetError()
	}
	return nil
}

// Clear clears the current rendering target with the drawing color.
func (r *Renderer) Clear() error {
	if C.SDL_RenderClear(&r.r) != 0 {
		return GetError()
	}
	return nil
}

// Present updates the screen with rendering performed.
func (r *Renderer) Present() {
	C.SDL_RenderPresent(&r.r)
}

// Destroy destroys the renderer.  The renderer should not be used after
// calling Destroy.
func (r *Renderer) Destroy() {
	C.SDL_DestroyRenderer(&r.r)
}
