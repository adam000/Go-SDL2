package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/adam000/Go-SDL2/image"
	"github.com/adam000/Go-SDL2/sdl"
	"github.com/adam000/Go-SDL2/sdl/keys"
)

var fullscreen = flag.Bool("fullscreen", false, "fullscreen window")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		sdl.Quit()
		os.Exit(2)
	}
	var err error
	go sdl.Do(func() { err = run() })
	sdl.Main()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	defer sdl.Quit()
	sdl.Init(sdl.InitVideo)

	surfaces, err := loadImages(flag.Args())
	defer func() {
		for _, s := range surfaces {
			s.Free()
		}
	}()
	if err != nil {
		return err
	}

	_, renderer, err := openWindow(flag.Arg(0), maxSize(surfaces))
	if err != nil {
		return err
	}

	textures, err := convertToTextures(renderer, surfaces)
	defer func() {
		for _, t := range textures {
			t.Destroy()
		}
	}()
	if err != nil {
		return err
	}

	mainLoop(renderer, textures)
	return nil
}

func mainLoop(renderer *sdl.Renderer, textures []*sdl.Texture) {
	currTex := 0
	for {
		// Poll for events
		for sdl.HasEvent() {
			switch ev := sdl.PollEvent(); ev.Type() {
			case sdl.QuitEventType:
				fmt.Println("QUIT")
				return
			case sdl.KeyDownEventType:
				ev := ev.(*sdl.KeyboardEvent)
				fmt.Println("KEY", ev.Code)
				switch ev.Code {
				case keys.Escape:
					return
				case keys.Left:
					if currTex > 0 {
						currTex--
					} else {
						currTex = len(textures) - 1
					}
				case keys.Right:
					if currTex < len(textures)-1 {
						currTex++
					} else {
						currTex = 0
					}
				}
			}
		}

		// Display image
		renderer.Clear()
		renderer.CopyTexture(textures[currTex], nil, nil)
		renderer.Present()

		// Wait a bit
		time.Sleep(100 * time.Millisecond)
	}
}

func openWindow(title string, size sdl.Point) (*sdl.Window, *sdl.Renderer, error) {
	var windowFlags sdl.WindowFlag
	if *fullscreen {
		windowFlags |= sdl.WindowFullscreen
	}
	window, err := sdl.NewWindow(title, sdl.Rect(0, 0, size.X, size.Y), windowFlags)
	if err != nil {
		return window, nil, err
	}
	renderer, err := window.CreateRenderer(-1)
	return window, renderer, err
}

func loadImages(names []string) ([]*sdl.Surface, error) {
	surfaces := make([]*sdl.Surface, 0, len(names))
	for _, name := range names {
		s, err := image.Load(name)
		if err != nil {
			return surfaces, &os.PathError{Op: "open", Path: name, Err: err}
		}
		surfaces = append(surfaces, s)
	}
	return surfaces, nil
}

func convertToTextures(renderer *sdl.Renderer, surfaces []*sdl.Surface) ([]*sdl.Texture, error) {
	textures := make([]*sdl.Texture, 0, len(surfaces))
	for _, s := range surfaces {
		t, err := sdl.NewTextureFromSurface(renderer, s)
		if err != nil {
			return textures, err
		}
		textures = append(textures, t)
	}
	return textures, nil
}

func maxSize(s []*sdl.Surface) sdl.Point {
	var size sdl.Point
	for _, ss := range s {
		z := ss.Size()
		if z.X > size.X {
			size.X = z.X
		}
		if z.Y > size.Y {
			size.Y = z.Y
		}
	}
	return size
}
