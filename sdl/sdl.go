/*
Package sdl provides a binding of SDL2 with an object-oriented twist.

The Do Function

SDL is not an inherently thread-safe library, and some operating systems force
windowing calls to be made on the main thread.  Thus, most functions in this
package must be called from the Do function, and the main function must call
sdl.Main.

	package main

	import "github.com/adam000/Go-SDL2/sdl"

	func main() {
		go run()
		sdl.Main()
	}

	func run() {
		defer sdl.Quit()

		sdl.Do(func() {
			// make SDL calls here
		})
	}
*/
package sdl

// #cgo pkg-config: sdl2
//
// #include "SDL.h"
import "C"

import (
	"runtime"
)

func init() {
	// Force main to stay on main thread.
	runtime.LockOSThread()

	// Run SDL_Init as early as possible. The SDL wiki claims that creating threads
	// before calling SDL_Init is bad (but they should feel bad).
	if C.SDL_Init(0) != 0 {
		panic(GetError())
	}
}

// Main runs the main SDL service loop.
// The binary's main.main must call sdl.Main() to run this loop.
// If the binary needs to do other work, it must do it in separate goroutines.
// Main will return after calling Quit.
func Main() {
	// Taken from https://code.google.com/p/go-wiki/wiki/LockOSThread.
	for f := range mainFunc {
		f()
	}
}

var mainFunc = make(chan func())

// Do executes a function on the main thread.  Calls to Do cannot nest --
// calling Do inside of a function passed to Do will cause deadlock.
func Do(f func()) {
	// Taken from https://code.google.com/p/go-wiki/wiki/LockOSThread.
	done := make(chan bool, 1)
	mainFunc <- func() {
		f()
		done <- true
	}
	<-done
}

// An InitFlag represents a set of SDL subsystems to initialize.
type InitFlag uint32

// InitFlag masks.
const (
	InitTimer          InitFlag = C.SDL_INIT_TIMER
	InitAudio          InitFlag = C.SDL_INIT_AUDIO
	InitVideo          InitFlag = C.SDL_INIT_VIDEO    // InitVideo implies InitEvents
	InitJoystick       InitFlag = C.SDL_INIT_JOYSTICK // InitJoystick implies InitEvents
	InitHaptic         InitFlag = C.SDL_INIT_HAPTIC
	InitGameController InitFlag = C.SDL_INIT_GAMECONTROLLER // InitGameController implies InitJoystick
	InitEvents         InitFlag = C.SDL_INIT_EVENTS
	InitNoParachute    InitFlag = C.SDL_INIT_NOPARACHUTE // Don't catch fatal signals

	InitEverything InitFlag = C.SDL_INIT_EVERYTHING
)

// Init initializes SDL subsystems.  Multiple flags will be ORed together.
// You don't have to call Init unless you need a particular subsystem.
func Init(flags ...InitFlag) error {
	var f InitFlag
	for i := range flags {
		f |= flags[i]
	}
	if C.SDL_InitSubSystem(C.Uint32(f)) != 0 {
		return GetError()
	}
	return nil
}

// Quit cleans up SDL.  Main will return after calling Quit.
func Quit() {
	C.SDL_Quit()
	close(mainFunc)
}

// Error stores an SDL error.
type Error string

func (e Error) Error() string {
	return "sdl: " + string(e)
}

// GetError returns the current SDL error as a Go error value.
// This is internal to SDL but exported because it is cross-package.
func GetError() error {
	e := C.SDL_GetError()
	if *e == 0 {
		// empty string, no error.
		return nil
	}
	return Error(C.GoString(e))
}
