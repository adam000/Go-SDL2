package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

//////////////////////////////////////////
// Contains Window and Renderer methods /
////////////////////////////////////////

// Window Handling

type WindowFlag uint32

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

const WindowPosUndefined = C.SDL_WINDOWPOS_UNDEFINED
const WindowPosCentered = C.SDL_WINDOWPOS_CENTERED

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
	Name             string
	Flags            RendererFlags
	TextureFormats   []PixelFormatEnum
	MaxTextureWidth  int
	MaxTextureHeight int
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

func (w Window) Renderer() (Renderer, error) {
	r := C.SDL_GetRenderer(w.w)
	if r == nil {
		return Renderer{}, getError()
	}
	return Renderer{r}, nil
}

func (r Renderer) Info() (*RendererInfo, error) {
	var info C.SDL_RendererInfo
	if C.SDL_GetRendererInfo(r.r, &info) != 0 {
		return nil, getError()
	}

	formats := make([]PixelFormatEnum, info.num_texture_formats)
	for i := range formats {
		formats[i] = PixelFormatEnum(info.texture_formats[i])
	}
	return &RendererInfo{
		Name:             C.GoString(info.name),
		Flags:            RendererFlags(info.flags),
		TextureFormats:   formats,
		MaxTextureWidth:  int(info.max_texture_width),
		MaxTextureHeight: int(info.max_texture_height),
	}, nil
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

// Clear clears the current rendering target with the drawing color.
func (r Renderer) Clear() error {
	if C.SDL_RenderClear(r.r) != 0 {
		return getError()
	}
	return nil
}

func (r Renderer) Present() {
	C.SDL_RenderPresent(r.r)
}
