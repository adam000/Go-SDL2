package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
//
// Uint32 GetType(SDL_Event* ev) {
//		return ev->type;
// }
//
// SDL_QuitEvent* ConvertToQuitEvent(SDL_Event* ev) {
//		return &ev->quit;
// }
//
// SDL_KeyboardEvent* ConvertToKeyboardEvent(SDL_Event* ev) {
//		return &ev->key;
// }
import "C"

import (
	"fmt"
)

// START EventType

type EventType uint16

const (
	FIRSTEVENT EventType = 0
)

/* Application events */
const (
	QUIT EventType = 0x100 + iota
	/* These application events have special meaning on iOS, see README-ios.txt for details */
	APP_TERMINATING
	APP_LOWMEMORY
	APP_WILLENTERBACKGROUND
	APP_DIDENTERBACKGROUND
	APP_WILLENTERFOREGROUND
	APP_DIDENTERFOREGROUND
)

/* Window events */
const (
	WINDOWEVENT EventType = 0x200 + iota
	SYSWMEVENT
)

/* Keyboard events */
const (
	KEYDOWN EventType = 0x300 + iota
	KEYUP
	TEXTEDITING
	TEXTINPUT
)

/* Mouse events */
const (
	MOUSEMOTION EventType = 0x400 + iota
	MOUSEBUTTONDOWN
	MOUSEBUTTONUP
	MOUSEWHEEL
)

/* Joystick events */
const (
	JOYAXISMOTION EventType = 0x600 + iota
	JOYBALLMOTION
	JOYHATMOTION
	JOYBUTTONDOWN
	JOYBUTTONUP
	JOYDEVICEADDED
	JOYDEVICEREMOVED
)

/* Game controller events */
const (
	CONTROLLERAXISMOTION EventType = 0x650 + iota
	CONTROLLERBUTTONDOWN
	CONTROLLERBUTTONUP
	CONTROLLERDEVICEADDED
	CONTROLLERDEVICEREMOVED
	CONTROLLERDEVICEREMAPPED
)

/* Touch events */
const (
	FINGERDOWN EventType = 0x700 + iota
	FINGERUP
	FINGERMOTION
)

/* Gesture events */
const (
	DOLLARGESTURE EventType = 0x800 + iota
	DOLLARRECORD
	MULTIGESTURE
)

/* Clipboard events */
const (
	CLIPBOARDUPDATE EventType = 0x900
)

/* Drag and drop events */
const (
	DROPFILE EventType = 0x1000
)

/*
Events ::USEREVENT through ::LASTEVENT are for your use
and should be allocated with RegisterEvents()
*/
const (
	USEREVENT EventType = 0x8000
	// This last event is only for bounding internal arrays
	LASTEVENT EventType = 0xFFFF
)

// END EventType

type Event interface {
	Timestamp() uint32
}

func PollEvent() (ev Event) {
	var cEvent C.SDL_Event
	if isEvent := C.SDL_PollEvent(&cEvent); isEvent == 0 {
		return nil
	}

	return convertEvent(&cEvent)
}

func convertEvent(cEvent *C.SDL_Event) (ev Event) {
	switch EventType(C.GetType(cEvent)) {
	case QUIT:
		return QuitEvent{*C.ConvertToQuitEvent(cEvent)}
	case KEYDOWN, KEYUP, TEXTEDITING, TEXTINPUT:
		return KeyboardEvent{*C.ConvertToKeyboardEvent(cEvent)}
	default:
		fmt.Printf("Unhandled event with int: %d\n", int(C.GetType(cEvent)));
		return nil
	}
}

// TODO make this implement Event
type CommonEvent struct {
	ev C.SDL_CommonEvent
}

// TODO make this implement Event
type WindowEvent struct {
	ev C.SDL_WindowEvent
}

type KeyboardEvent struct {
	ev C.SDL_KeyboardEvent
}

func (event KeyboardEvent) Timestamp() uint32 {
	return uint32(event.ev.timestamp)
}

func (event KeyboardEvent) WindowID() uint32 {
	return uint32(event.ev.windowID)
}

func (event KeyboardEvent) State() uint8 {
	return uint8(event.ev.state)
}

func (event KeyboardEvent) Repeat() uint8 {
	return uint8(event.ev.repeat)
}

// TODO this
/*
func (event KeyboardEvent) Keysym() Keysym {
}
*/

// TODO make this implement Event
type TextEditingEvent struct {
	ev C.SDL_TextEditingEvent
}

// TODO make this implement Event
type TextInputEvent struct {
	ev C.SDL_TextInputEvent
}

// TODO make this implement Event
type MouseMotionEvent struct {
	ev C.SDL_MouseMotionEvent
}

// TODO make this implement Event
type MouseButtonEvent struct {
	ev C.SDL_MouseButtonEvent
}

// TODO make this implement Event
type MouseWheelEvent struct {
	ev C.SDL_MouseWheelEvent
}

// TODO make this implement Event
type JoyAxisEvent struct {
	ev C.SDL_JoyAxisEvent
}

// TODO make this implement Event
type JoyBallEvent struct {
	ev C.SDL_JoyBallEvent
}

// TODO make this implement Event
type JoyHatEvent struct {
	ev C.SDL_JoyHatEvent
}

// TODO make this implement Event
type JoyButtonEvent struct {
	ev C.SDL_JoyButtonEvent
}

// TODO make this implement Event
type JoyDeviceEvent struct {
	ev C.SDL_JoyDeviceEvent
}

// TODO make this implement Event
type ControllerAxisEvent struct {
	ev C.SDL_ControllerAxisEvent
}

// TODO make this implement Event
type ControllerButtonEvent struct {
	ev C.SDL_ControllerButtonEvent
}

// TODO make this implement Event
type ControllerDeviceEvent struct {
	ev C.SDL_ControllerDeviceEvent
}

type QuitEvent struct {
	ev C.SDL_QuitEvent
}

func (event QuitEvent) Timestamp() uint32 {
	return uint32(event.ev.timestamp)
}

// TODO make this implement Event
type UserEvent struct {
	ev C.SDL_UserEvent
}

// TODO make this implement Event
type SysWMEvent struct {
	ev C.SDL_SysWMEvent
}

// TODO make this implement Event
type TouchFingerEvent struct {
	ev C.SDL_TouchFingerEvent
}

// TODO make this implement Event
type MultiGestureEvent struct {
	ev C.SDL_MultiGestureEvent
}

// TODO make this implement Event
type DollarGestureEvent struct {
	ev C.SDL_DollarGestureEvent
}

// TODO make this implement Event
type DropEvent struct {
	ev C.SDL_DropEvent
}

