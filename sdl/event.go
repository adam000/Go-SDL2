package sdl

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_image
//
// #include "SDL.h"
import "C"

import (
	"fmt"
	"unsafe"
)

//
// The only way to keep this file sane is with vim folds. `:help fold` for more info
//

// START EventType {{{1

// TODO(adam): Suffix with Event instead of Ev?

// EventType represents the type of an application event.
type EventType uint32

// IsUserEvent reports whether t is a custom application event.
func (t EventType) IsUserEvent() bool {
	return t >= UserEvent && t <= LastEvent
}

// Special event numbers
const (
	FirstEvent EventType = C.SDL_FIRSTEVENT

	// Events UserEvent through LastEvent are for your use, and should be allocated
	// with RegisterEvents.
	UserEvent EventType = C.SDL_USEREVENT
	LastEvent EventType = C.SDL_LASTEVENT
)

// Application events
const (
	QuitEv EventType = C.SDL_QUIT

	// These application events have special meaning on iOS, see README-ios.txt for details
	AppTerminatingEv         EventType = C.SDL_APP_TERMINATING
	AppLowMemoryEv           EventType = C.SDL_APP_LOWMEMORY
	AppWillEnterBackgroundEv EventType = C.SDL_APP_WILLENTERBACKGROUND
	AppDidEnterBackgroundEv  EventType = C.SDL_APP_DIDENTERBACKGROUND
	AppWillEnterForegroundEv EventType = C.SDL_APP_WILLENTERFOREGROUND
	AppDidEnterForegroundEv  EventType = C.SDL_APP_DIDENTERFOREGROUND
)

// Window events
const (
	WindowEv EventType = C.SDL_WINDOWEVENT
	SysWmEv  EventType = C.SDL_SYSWMEVENT
)

// Keyboard events
const (
	KeyDownEv     EventType = C.SDL_KEYDOWN
	KeyUpEv       EventType = C.SDL_KEYUP
	TextEditingEv EventType = C.SDL_TEXTEDITING
	TextInputEv   EventType = C.SDL_TEXTINPUT
)

// Mouse events
const (
	MouseMotionEv     EventType = C.SDL_MOUSEMOTION
	MouseButtonDownEv EventType = C.SDL_MOUSEBUTTONDOWN
	MouseButtonUpEv   EventType = C.SDL_MOUSEBUTTONUP
	MouseWheelEv      EventType = C.SDL_MOUSEWHEEL
)

// Joystick events
const (
	JoyAxisMotionEv    EventType = C.SDL_JOYAXISMOTION
	JoyBallMotionEv    EventType = C.SDL_JOYBALLMOTION
	JoyHatMotionEv     EventType = C.SDL_JOYHATMOTION
	JoyButtonDownEv    EventType = C.SDL_JOYBUTTONDOWN
	JoyButtonUpEv      EventType = C.SDL_JOYBUTTONUP
	JoyDeviceAddedEv   EventType = C.SDL_JOYDEVICEADDED
	JoyDeviceRemovedEv EventType = C.SDL_JOYDEVICEREMOVED
)

// Game controller events
const (
	ControllerAxisMotionEv     EventType = C.SDL_CONTROLLERAXISMOTION
	ControllerButtonDownEv     EventType = C.SDL_CONTROLLERBUTTONDOWN
	ControllerButtonUpEv       EventType = C.SDL_CONTROLLERBUTTONUP
	ControllerDeviceAddedEv    EventType = C.SDL_CONTROLLERDEVICEADDED
	ControllerDeviceRemovedEv  EventType = C.SDL_CONTROLLERDEVICEREMOVED
	ControllerDeviceRemappedEv EventType = C.SDL_CONTROLLERDEVICEREMAPPED
)

// Touch events
const (
	FingerDownEv   EventType = C.SDL_FINGERDOWN
	FingerUpEv     EventType = C.SDL_FINGERUP
	FingerMotionEv EventType = C.SDL_FINGERMOTION
)

// Gesture events
const (
	DollarGestureEv EventType = C.SDL_DOLLARGESTURE
	DollarRecordEv  EventType = C.SDL_DOLLARRECORD
	MultiGestureEv  EventType = C.SDL_MULTIGESTURE
)

// Clipboard events
const (
	ClipboardUpdateEv EventType = C.SDL_CLIPBOARDUPDATE
)

// Drag and drop events
const (
	DropFileEv EventType = C.SDL_DROPFILE
)

// END EventType }}}1

// Event is implemented by all SDL events.
type Event interface {
	// Type returns the event's type.
	Type() EventType

	// Timestamp returns the number of milliseconds since the SDL library initialization.
	Timestamp() uint32
}

// PollEvent returns the next available event, or nil if there is no event pending.
func PollEvent() Event {
	var cEvent C.SDL_Event
	if C.SDL_PollEvent(&cEvent) == 0 {
		return nil
	}
	return convertEvent(unsafe.Pointer(&cEvent))
}

// HasEvent returns whether there is a pending event available.
func HasEvent() bool {
	return C.SDL_PollEvent(nil) != 0
}

func convertEvent(cEvent unsafe.Pointer) Event {
	common := (*C.SDL_CommonEvent)(cEvent)
	switch EventType(common._type) {
	case QuitEv:
		// Quit events don't hold any data beyond the common events.
		return commonEvent{common}
	case KeyDownEv, KeyUpEv:
		return KeyboardEvent{(*C.SDL_KeyboardEvent)(cEvent)}
	default:
		fmt.Println("Unhandled event with int:", int(common._type))
		return commonEvent{common}
	}
}

// {{{1 Event Structs

// {{{2 CommonEvent

// commonEvent holds fields common to all events.  It is not exported because it
// doesn't provide anything useful outside of the Event interface.
type commonEvent struct {
	ev *C.SDL_CommonEvent
}

func (e commonEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e commonEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
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
	return EventType(e.e._type)
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
