package sdl

// TODO put these C functions in their own event.c file

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
	"unsafe"
)

//
// The only way to keep this file sane is with vim folds. `:help fold` for more info
//

// START EventType {{{1

type EventType uint16

const (
	FirstEvent EventType = 0
)

/* Application events */
const (
	QuitEv EventType = 0x100 + iota
	/* These application events have special meaning on iOS, see README-ios.txt for details */
	AppTerminatingEv
	AppLowMemoryEv
	AppWillEnterBackgroundEv
	AppDidEnterBackgroundEv
	AppWillEnterForegroundEv
	AppDidEnterForegroundEv
)

/* Window events */
const (
	WindowEv EventType = 0x200 + iota
	SysWmEv
)

/* Keyboard events */
const (
	KeyDownEv EventType = 0x300 + iota
	KeyUpEv
	TextEditingEv
	TextInputEv
)

/* Mouse events */
const (
	MouseMotionEv EventType = 0x400 + iota
	MouseButtonDownEv
	MouseButtonUpEv
	MouseWheelEv
)

/* Joystick events */
const (
	JoyAxisMotionEv EventType = 0x600 + iota
	JoyBallMotionEv
	JoyHatMotionEv
	JoyButtonDownEv
	JoyButtonUpEv
	JoyDeviceAddedEv
	JoyDeviceRemovedEv
)

/* Game controller events */
const (
	ControllerAxisMotionEv EventType = 0x650 + iota
	ControllerButtonDownEv
	ControllerButtonUpEv
	ControllerDeviceAddedEv
	ControllerDeviceRemovedEv
	ControllerDeviceRemappedEv
)

/* Touch events */
const (
	FingerDownEv EventType = 0x700 + iota
	FingerUpEv
	FingerMotionEv
)

/* Gesture events */
const (
	DollarGestureEv EventType = 0x800 + iota
	DollarRecordEv
	MultiGestureEv
)

/* Clipboard events */
const (
	ClipboardUpdateEv EventType = 0x900
)

/* Drag and drop events */
const (
	DropFileEv EventType = 0x1000
)

/*
Events ::USEREVENT through ::LASTEVENT are for your use
and should be allocated with RegisterEvents()
*/
const (
	UserEv EventType = 0x8000
	// This last event is only for bounding internal arrays
	LastEv EventType = 0xFFFF
)

// END EventType }}}1

type Event interface {
	Timestamp() uint32
	Type() EventType
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
	case QuitEv:
		return QuitEvent{C.ConvertToQuitEvent(cEvent)}
	case KeyDownEv, KeyUpEv:
		return KeyboardEvent{C.ConvertToKeyboardEvent(cEvent)}
	default:
		fmt.Printf("Unhandled event with int: %d\n", int(C.GetType(cEvent)))
		return nil
	}
}

// {{{1 Event Structs

// {{{2 CommonEvent
// TODO make this implement Event
type CommonEvent struct {
	ev C.SDL_CommonEvent
}
// }}}2 CommonEvent

// {{{2 WindowEvent
// TODO make this implement Event
type WindowEvent struct {
	ev C.SDL_WindowEvent
}
// }}}2 WindowEvent

// {{{2 KeyboardEvent
type KeyboardEvent struct {
	e *C.SDL_KeyboardEvent
}

func (e KeyboardEvent) Timestamp() uint32 {
	return uint32(e.e.timestamp)
}

func (e KeyboardEvent) Type() EventType {
	// Have to do this icky thing to get what type of keyboard event it is
	return EventType(C.GetType((*C.SDL_Event)(unsafe.Pointer(e.e))))
}

func (e KeyboardEvent) WindowID() uint32 {
	return uint32(e.e.windowID)
}

func (e KeyboardEvent) State() uint8 {
	return uint8(e.e.state)
}

func (e KeyboardEvent) Repeat() uint8 {
	return uint8(e.e.repeat)
}

// TODO this
/*
func (e KeyboardEvent) Keysym() Keysym {
}
*/
// }}}2 KeyboardEvent

// {{{2 TextEditingEvent
// TODO make this implement Event
type TextEditingEvent struct {
	ev C.SDL_TextEditingEvent
}
// }}}2 TextEditingEvent

// {{{2 TextInputEvent
// TODO make this implement Event
type TextInputEvent struct {
	ev C.SDL_TextInputEvent
}
// }}}2 TextInputEvent

// {{{2 MouseMotionEvent
// TODO make this implement Event
type MouseMotionEvent struct {
	ev C.SDL_MouseMotionEvent
}
// }}}2 MouseMotionEvent

// {{{2 MouseButtonEvent
// TODO make this implement Event
type MouseButtonEvent struct {
	ev C.SDL_MouseButtonEvent
}
// }}}2 MouseButtonEvent

// {{{2 MouseWheelEvent
// TODO make this implement Event
type MouseWheelEvent struct {
	ev C.SDL_MouseWheelEvent
}
// }}}2 MouseWheelEvent

// {{{2 JoyAxisEvent
// TODO make this implement Event
type JoyAxisEvent struct {
	ev C.SDL_JoyAxisEvent
}
// }}}2 JoyAxisEvent

// {{{2 JoyBallEvent
// TODO make this implement Event
type JoyBallEvent struct {
	ev C.SDL_JoyBallEvent
}
// }}}2 JoyBallEvent

// {{{2 JoyHatEvent
// TODO make this implement Event
type JoyHatEvent struct {
	ev C.SDL_JoyHatEvent
}
// }}}2 JoyHatEvent

// {{{2 JoyButtonEvent
// TODO make this implement Event
type JoyButtonEvent struct {
	ev C.SDL_JoyButtonEvent
}
// }}}2 JoyButtonEvent

// {{{2 JoyDeviceEvent
// TODO make this implement Event
type JoyDeviceEvent struct {
	ev C.SDL_JoyDeviceEvent
}
// }}}2 JoyDeviceEvent

// {{{2 ControllerAxisEvent
// TODO make this implement Event
type ControllerAxisEvent struct {
	ev C.SDL_ControllerAxisEvent
}
// }}}2 ControllerAxisEvent

// {{{2 ControllerButtonEvent
// TODO make this implement Event
type ControllerButtonEvent struct {
	ev C.SDL_ControllerButtonEvent
}
// }}}2 ControllerButtonEvent

// {{{2 ControllerDeviceEvent
// TODO make this implement Event
type ControllerDeviceEvent struct {
	ev C.SDL_ControllerDeviceEvent
}
// }}}2 ControllerDeviceEvent

// {{{2 QuitEvent
type QuitEvent struct {
	e *C.SDL_QuitEvent
}

func (e QuitEvent) Timestamp() uint32 {
	return uint32(e.e.timestamp)
}

func (e QuitEvent) Type() EventType {
	return QuitEv
}
// }}}2 QuitEvent

// {{{2 UserEvent
// TODO make this implement Event
type UserEvent struct {
	ev C.SDL_UserEvent
}
// }}}2 UserEvent

// {{{2 SysWMEvent
// TODO make this implement Event
type SysWMEvent struct {
	ev C.SDL_SysWMEvent
}
// }}}2 SysWMEvent

// {{{2 TouchFingerEvent
// TODO make this implement Event
type TouchFingerEvent struct {
	ev C.SDL_TouchFingerEvent
}
// }}}2 TouchFingerEvent

// {{{2 MultiGestureEvent
// TODO make this implement Event
type MultiGestureEvent struct {
	ev C.SDL_MultiGestureEvent
}
// }}}2 MultiGestureEvent

// {{{2 DollarGestureEvent
// TODO make this implement Event
type DollarGestureEvent struct {
	ev C.SDL_DollarGestureEvent
}
// }}}2 DollarGestureEvent

// {{{2 DropEvent
// TODO make this implement Event
type DropEvent struct {
	ev C.SDL_DropEvent
}
// }}}2 DropEvent

// }}}1 Event Structs
