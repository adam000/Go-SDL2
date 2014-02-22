package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/adam000/Go-SDL2/image"
	"github.com/adam000/Go-SDL2/sdl"
)

var fullscreen = flag.Bool("fullscreen", false, "fullscreen window")

func main() {
	flag.Parse()
	defer sdl.Quit()

	// Open window
	var windowFlags sdl.WindowFlag
	if *fullscreen {
		windowFlags |= sdl.WindowFullscreen
	}
	// TODO(light): surface width & height
	window, err := sdl.NewWindow(flag.Arg(0), 0, 0, 640, 480, windowFlags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	renderer, err := sdl.NewRenderer(window, -1, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Load textures
	textures := make([]sdl.Texture, flag.NArg())
	for i, name := range flag.Args() {
		surf, err := image.Load(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v", name, err)
			os.Exit(1)
		}
		textures[i], err = surf.ToTexture(renderer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v", name, err)
			os.Exit(1)
		}
	}
	defer func() {
		for _, tex := range textures {
			tex.Destroy()
		}
	}()

	// Display
	currTex := 0
mainLoop:
	for {
		// Poll for events
		for {
			ev := sdl.PollEvent()
			if ev == nil {
				break
			}
			switch ev.Type() {
			case sdl.QuitEv:
				fmt.Println("QUIT")
				break mainLoop
			case sdl.KeyDownEv:
				fmt.Println("KEY")
				// TODO(light): advance texture
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
