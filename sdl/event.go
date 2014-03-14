package sdl

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

// EventType represents the type of an application event.
type EventType uint32

// IsUserEvent reports whether t is a custom application event.
func (t EventType) IsUserEvent() bool {
	return t >= UserEventType && t <= LastEventType
}

// Special event numbers
const (
	FirstEventType EventType = C.SDL_FIRSTEVENT

	// Events UserEvent through LastEvent are for your use, and should be allocated
	// with RegisterEvents.
	UserEventType EventType = C.SDL_USEREVENT
	LastEventType EventType = C.SDL_LASTEVENT
)

// Application events
const (
	QuitEventType EventType = C.SDL_QUIT

	// These application events have special meaning on iOS, see README-ios.txt for details
	AppTerminatingEventType         EventType = C.SDL_APP_TERMINATING
	AppLowMemoryEventType           EventType = C.SDL_APP_LOWMEMORY
	AppWillEnterBackgroundEventType EventType = C.SDL_APP_WILLENTERBACKGROUND
	AppDidEnterBackgroundEventType  EventType = C.SDL_APP_DIDENTERBACKGROUND
	AppWillEnterForegroundEventType EventType = C.SDL_APP_WILLENTERFOREGROUND
	AppDidEnterForegroundEventType  EventType = C.SDL_APP_DIDENTERFOREGROUND
)

// Window events
const (
	WindowEventType EventType = C.SDL_WINDOWEVENT
	SysWMEventType  EventType = C.SDL_SYSWMEVENT
)

// Keyboard events
const (
	KeyDownEventType     EventType = C.SDL_KEYDOWN
	KeyUpEventType       EventType = C.SDL_KEYUP
	TextEditingEventType EventType = C.SDL_TEXTEDITING
	TextInputEventType   EventType = C.SDL_TEXTINPUT
)

// Mouse events
const (
	MouseMotionEventType     EventType = C.SDL_MOUSEMOTION
	MouseButtonDownEventType EventType = C.SDL_MOUSEBUTTONDOWN
	MouseButtonUpEventType   EventType = C.SDL_MOUSEBUTTONUP
	MouseWheelEventType      EventType = C.SDL_MOUSEWHEEL
)

// Joystick events
const (
	JoyAxisMotionEventType    EventType = C.SDL_JOYAXISMOTION
	JoyBallMotionEventType    EventType = C.SDL_JOYBALLMOTION
	JoyHatMotionEventType     EventType = C.SDL_JOYHATMOTION
	JoyButtonDownEventType    EventType = C.SDL_JOYBUTTONDOWN
	JoyButtonUpEventType      EventType = C.SDL_JOYBUTTONUP
	JoyDeviceAddedEventType   EventType = C.SDL_JOYDEVICEADDED
	JoyDeviceRemovedEventType EventType = C.SDL_JOYDEVICEREMOVED
)

// Game controller events
const (
	ControllerAxisMotionEventType     EventType = C.SDL_CONTROLLERAXISMOTION
	ControllerButtonDownEventType     EventType = C.SDL_CONTROLLERBUTTONDOWN
	ControllerButtonUpEventType       EventType = C.SDL_CONTROLLERBUTTONUP
	ControllerDeviceAddedEventType    EventType = C.SDL_CONTROLLERDEVICEADDED
	ControllerDeviceRemovedEventType  EventType = C.SDL_CONTROLLERDEVICEREMOVED
	ControllerDeviceRemappedEventType EventType = C.SDL_CONTROLLERDEVICEREMAPPED
)

// Touch events
const (
	FingerDownEventType   EventType = C.SDL_FINGERDOWN
	FingerUpEventType     EventType = C.SDL_FINGERUP
	FingerMotionEventType EventType = C.SDL_FINGERMOTION
)

// Gesture events
const (
	DollarGestureEventType EventType = C.SDL_DOLLARGESTURE
	DollarRecordEventType  EventType = C.SDL_DOLLARRECORD
	MultiGestureEventType  EventType = C.SDL_MULTIGESTURE
)

// Clipboard events
const (
	ClipboardUpdateEventType EventType = C.SDL_CLIPBOARDUPDATE
)

// Drag and drop events
const (
	DropFileEventType EventType = C.SDL_DROPFILE
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
	case QuitEventType:
		// Quit events don't hold any data beyond the common events.
		return commonEvent{common}
	case KeyDownEventType, KeyUpEventType:
		return KeyboardEvent{(*C.SDL_KeyboardEvent)(cEvent)}
	case WindowEventType:
		return WindowEvent{(*C.SDL_WindowEvent)(cEvent)}
	case MouseMotionEventType:
		return MouseMotionEvent{(*C.SDL_MouseMotionEvent)(cEvent)}
	case MouseButtonDownEventType, MouseButtonUpEventType:
		return MouseButtonEvent{(*C.SDL_MouseButtonEvent)(cEvent)}
	case MouseWheelEventType:
		return MouseWheelEvent{(*C.SDL_MouseWheelEvent)(cEvent)}
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

// WindowEvent holds window state change event data.
type WindowEvent struct {
	ev *C.SDL_WindowEvent
}

// Type returns WindowEventType.
func (e WindowEvent) Type() EventType {
	return WindowEventType
}

func (e WindowEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// WindowID returns the ID of the window that this event occurred in.
func (e WindowEvent) WindowID() uint32 {
	return uint32(e.ev.windowID)
}

// WindowEvent returns the specific state change that occurred on the window.
func (e WindowEvent) WindowEvent() WindowEventID {
	return WindowEventID(e.ev.event)
}

// Data returns the event-dependent data.
// For move events, this is the new (x, y) position of the window.
// For resize events, this is the new window size.
func (e WindowEvent) Data() (data1, data2 int32) {
	return int32(e.ev.data1), int32(e.ev.data2)
}

// WindowEventID is a window event subtype.
type WindowEventID uint8

// Window event subtypes
const (
	WindowEventShown       WindowEventID = C.SDL_WINDOWEVENT_SHOWN
	WindowEventHidden      WindowEventID = C.SDL_WINDOWEVENT_HIDDEN
	WindowEventExposed     WindowEventID = C.SDL_WINDOWEVENT_EXPOSED
	WindowEventMoved       WindowEventID = C.SDL_WINDOWEVENT_MOVED
	WindowEventResized     WindowEventID = C.SDL_WINDOWEVENT_RESIZED
	WindowEventSizeChanged WindowEventID = C.SDL_WINDOWEVENT_SIZE_CHANGED
	WindowEventMinimized   WindowEventID = C.SDL_WINDOWEVENT_MINIMIZED
	WindowEventMaximized   WindowEventID = C.SDL_WINDOWEVENT_MAXIMIZED
	WindowEventRestored    WindowEventID = C.SDL_WINDOWEVENT_RESTORED
	WindowEventEnter       WindowEventID = C.SDL_WINDOWEVENT_ENTER
	WindowEventLeave       WindowEventID = C.SDL_WINDOWEVENT_LEAVE
	WindowEventFocusGained WindowEventID = C.SDL_WINDOWEVENT_FOCUS_GAINED
	WindowEventFocusLost   WindowEventID = C.SDL_WINDOWEVENT_FOCUS_LOST
	WindowEventClose       WindowEventID = C.SDL_WINDOWEVENT_CLOSE
)

// }}}2 WindowEvent

// {{{2 KeyboardEvent
type KeyboardEvent struct {
	ev *C.SDL_KeyboardEvent
}

func (e KeyboardEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

func (e KeyboardEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e KeyboardEvent) WindowID() uint32 {
	return uint32(e.ev.windowID)
}

func (e KeyboardEvent) State() uint8 {
	return uint8(e.ev.state)
}

// IsRepeat reports whether this event is a repeating key.
func (e KeyboardEvent) IsRepeat() bool {
	return e.ev.repeat != 0
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

// MouseMotionEvent holds a mouse movement event.
type MouseMotionEvent struct {
	ev *C.SDL_MouseMotionEvent
}

// Type returns MouseMotionEventType
func (e MouseMotionEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e MouseMotionEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// WindowID returns the window with mouse focus, or zero if no window has focus.
func (e MouseMotionEvent) WindowID() uint32 {
	return uint32(e.ev.windowID)
}

// Which returns the mouse which triggered the event.
func (e MouseMotionEvent) Which() uint32 {
	return uint32(e.ev.which)
}

// ButtonState returns a bitmask of the current button state.  Use MouseButton.Mask
func (e MouseMotionEvent) ButtonState() uint32 {
	return uint32(e.ev.state)
}

// Position returns the new (x, y) coordinate of the event, relative to the window.
func (e MouseMotionEvent) Position() (x, y int32) {
	return int32(e.ev.x), int32(e.ev.y)
}

// Delta returns the relative (x, y) motion captured by this event.
func (e MouseMotionEvent) Delta() (dx, dy int32) {
	return int32(e.ev.xrel), int32(e.ev.yrel)
}

// }}}2 MouseMotionEvent

// {{{2 MouseButtonEvent

// MouseButtonEvent holds a mouse button press or release event.
type MouseButtonEvent struct {
	ev *C.SDL_MouseButtonEvent
}

// Type returns either MouseButtonDownEventType or MouseButtonUpEventType.
func (e MouseButtonEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e MouseButtonEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// WindowID returns the window with mouse focus, or zero if no window has focus.
func (e MouseButtonEvent) WindowID() uint32 {
	return uint32(e.ev.windowID)
}

// Which returns the mouse which triggered the event.
func (e MouseButtonEvent) Which() uint32 {
	return uint32(e.ev.which)
}

// IsTouch reports whether this event was generated by a touch input device.
func (e MouseButtonEvent) IsTouch() bool {
	return e.ev.which == C.SDL_TOUCH_MOUSEID
}

// Button returns the button that changed.
func (e MouseButtonEvent) Button() MouseButton {
	return MouseButton(e.ev.button)
}

// IsPressed reports whether the button is pressed.
func (e MouseButtonEvent) IsPressed() bool {
	return e.ev.state == C.SDL_PRESSED
}

// TODO(light): this is only available in SDL 2.0.2 and above
// Clicks returns the number of clicks in sequence: 1 for single-click,
// 2 for double-click, etc.
//func (e MouseButtonEvent) Clicks() int {
//	return int(e.ev.clicks)
//}

// Position returns the (x, y) coordinate of the event, relative to the window.
func (e MouseButtonEvent) Position() (x, y int32) {
	return int32(e.ev.x), int32(e.ev.y)
}

// }}}2 MouseButtonEvent

// {{{2 MouseWheelEvent

// MouseWheelEvent holds a mouse wheel movement event.
type MouseWheelEvent struct {
	ev *C.SDL_MouseWheelEvent
}

// Type returns MouseWheelEventType.
func (e MouseWheelEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e MouseWheelEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// WindowID returns the window with mouse focus, or zero if no window has focus.
func (e MouseWheelEvent) WindowID() uint32 {
	return uint32(e.ev.windowID)
}

// Which returns the mouse which triggered the event.
func (e MouseWheelEvent) Which() uint32 {
	return uint32(e.ev.which)
}

// Scroll returns the amount scrolled in X and Y.  The axes increase right and up.
func (e MouseWheelEvent) Scroll() (x, y int32) {
	return int32(e.ev.x), int32(e.ev.y)
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
