package sdl

// #include "SDL.h"
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/adam000/Go-SDL2/sdl/keys"
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
	case TextEditingEventType:
		return TextEditingEvent{(*C.SDL_TextEditingEvent)(cEvent)}
	case TextInputEventType:
		return TextInputEvent{(*C.SDL_TextInputEvent)(cEvent)}
	case WindowEventType:
		return WindowEvent{(*C.SDL_WindowEvent)(cEvent)}
	case MouseMotionEventType:
		return MouseMotionEvent{(*C.SDL_MouseMotionEvent)(cEvent)}
	case MouseButtonDownEventType, MouseButtonUpEventType:
		return MouseButtonEvent{(*C.SDL_MouseButtonEvent)(cEvent)}
	case MouseWheelEventType:
		return MouseWheelEvent{(*C.SDL_MouseWheelEvent)(cEvent)}
	case JoyAxisMotionEventType:
		return JoyAxisEvent{(*C.SDL_JoyAxisEvent)(cEvent)}
	case JoyBallMotionEventType:
		return JoyBallEvent{(*C.SDL_JoyBallEvent)(cEvent)}
	case JoyHatMotionEventType:
		return JoyHatEvent{(*C.SDL_JoyHatEvent)(cEvent)}
	case JoyButtonDownEventType, JoyButtonUpEventType:
		return JoyButtonEvent{(*C.SDL_JoyButtonEvent)(cEvent)}
	case JoyDeviceAddedEventType, JoyDeviceRemovedEventType:
		return JoyDeviceEvent{(*C.SDL_JoyDeviceEvent)(cEvent)}
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

// KeySym returns the key information from this event.
func (e KeyboardEvent) KeySym() KeySym {
	return KeySym{
		ScanCode: int32(e.ev.keysym.scancode),
		KeyCode:  int32(e.ev.keysym.sym),
		Mod:      keys.Mod(e.ev.keysym.mod),
	}
}

// }}}2 KeyboardEvent

// {{{2 TextEditingEvent

// TextEditingEvent holds a partial text input event.  See the description of TextInputEvent.
type TextEditingEvent struct {
	ev *C.SDL_TextEditingEvent
}

// Type returns TextEditingEventType.
func (e TextEditingEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e TextEditingEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// WindowID returns the window with keyboard focus
func (e TextEditingEvent) WindowID() uint32 {
	return uint32(e.ev.windowID)
}

// Text returns the partial text.
func (e TextEditingEvent) Text() string {
	return C.GoString(&e.ev.text[0])
}

// Cursor returns the range of characters to edit.
func (e TextEditingEvent) Cursor() (start, n int) {
	return int(e.ev.start), int(e.ev.length)
}

// }}}2 TextEditingEvent

// {{{2 TextInputEvent

// TextInputEvent holds a complete text input event.
//
// For every text input, there are one or more text editing events followed by
// one text input event.  An input method may require multiple key presses to
// input a single character.  The text editing events allow an application to
// render feedback of receiving the characters before inputting the final
// character.
type TextInputEvent struct {
	ev *C.SDL_TextInputEvent
}

// Type returns TextInputEventType.
func (e TextInputEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e TextInputEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// WindowID returns the window with keyboard focus
func (e TextInputEvent) WindowID() uint32 {
	return uint32(e.ev.windowID)
}

// Text returns the inputted text.
func (e TextInputEvent) Text() string {
	return C.GoString(&e.ev.text[0])
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

// JoyAxisEvent holds a joystick axis movement event.
type JoyAxisEvent struct {
	ev *C.SDL_JoyAxisEvent
}

// Type returns JoyAxisMotionEventType.
func (e JoyAxisEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e JoyAxisEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// Which returns the joystick which triggered the event.
func (e JoyAxisEvent) Which() JoystickID {
	return JoystickID(e.ev.which)
}

// Axis returns the index of the axis that changed.
func (e JoyAxisEvent) Axis() uint8 {
	return uint8(e.ev.axis)
}

// Value returns the current position of the axis.
func (e JoyAxisEvent) Value() int16 {
	return int16(e.ev.value)
}

// }}}2 JoyAxisEvent

// {{{2 JoyBallEvent

// JoyBallEvent holds a joystick trackball motion event.
type JoyBallEvent struct {
	ev *C.SDL_JoyBallEvent
}

// Type returns JoyBallMotionEventType.
func (e JoyBallEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e JoyBallEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// Which returns the joystick which triggered the event.
func (e JoyBallEvent) Which() JoystickID {
	return JoystickID(e.ev.which)
}

// Ball returns the index of the trackball that changed.
func (e JoyBallEvent) Ball() uint8 {
	return uint8(e.ev.ball)
}

// Delta returns the relative (x, y) motion captured by this event.
func (e JoyBallEvent) Delta() (dx, dy int16) {
	return int16(e.ev.xrel), int16(e.ev.yrel)
}

// }}}2 JoyBallEvent

// {{{2 JoyHatEvent

// JoyHatEvent holds a EVENT
type JoyHatEvent struct {
	ev *C.SDL_JoyHatEvent
}

// Type returns JoyHatMotionEventType.
func (e JoyHatEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e JoyHatEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// Which returns the joystick which triggered the event.
func (e JoyHatEvent) Which() JoystickID {
	return JoystickID(e.ev.which)
}

// Hat returns the index of the hat that changed.
func (e JoyHatEvent) Hat() uint8 {
	return uint8(e.ev.hat)
}

// Position returns the new position of the hat.
func (e JoyHatEvent) Position() HatPosition {
	return HatPosition(e.ev.value)
}

// }}}2 JoyHatEvent

// {{{2 JoyButtonEvent

// JoyButtonEvent holds a EVENT
type JoyButtonEvent struct {
	ev *C.SDL_JoyButtonEvent
}

// Type returns either JoyButtonDownEventType or JoyButtonUpEventType.
func (e JoyButtonEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e JoyButtonEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// Which returns the joystick which triggered the event.
func (e JoyButtonEvent) Which() JoystickID {
	return JoystickID(e.ev.which)
}

// Button returns the index of the button that changed.
func (e JoyButtonEvent) Button() uint8 {
	return uint8(e.ev.button)
}

// IsPressed reports whether the button is pressed.
func (e JoyButtonEvent) IsPressed() bool {
	return e.ev.state == C.SDL_PRESSED
}

// }}}2 JoyButtonEvent

// {{{2 JoyDeviceEvent

// JoyDeviceEvent holds a joystick connection or disconnection event.
type JoyDeviceEvent struct {
	ev *C.SDL_JoyDeviceEvent
}

// Type returns either JoyDeviceAddedEventType or JoyDeviceRemovedEventType.
func (e JoyDeviceEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e JoyDeviceEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// Which returns the joystick device index for an added event or the instance ID for a removal event.
func (e JoyDeviceEvent) Which() int32 {
	return int32(e.ev.which)
}

// }}}2 JoyDeviceEvent

// {{{2 ControllerAxisEvent

// ControllerAxisEvent holds a EVENT
type ControllerAxisEvent struct {
	ev *C.SDL_ControllerAxisEvent
}

// Type returns...
func (e ControllerAxisEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e ControllerAxisEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 ControllerAxisEvent

// {{{2 ControllerButtonEvent

// ControllerButtonEvent holds a EVENT
type ControllerButtonEvent struct {
	ev *C.SDL_ControllerButtonEvent
}

// Type returns...
func (e ControllerButtonEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e ControllerButtonEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 ControllerButtonEvent

// {{{2 ControllerDeviceEvent

// ControllerDeviceEvent holds a EVENT
type ControllerDeviceEvent struct {
	ev *C.SDL_ControllerDeviceEvent
}

// Type returns...
func (e ControllerDeviceEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e ControllerDeviceEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 ControllerDeviceEvent

// {{{2 UserEvent

// UserEvent holds a EVENT
type UserEvent struct {
	ev *C.SDL_UserEvent
}

// Type returns...
func (e UserEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e UserEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 UserEvent

// {{{2 SysWMEvent

// SysWMEvent holds a EVENT
type SysWMEvent struct {
	ev *C.SDL_SysWMEvent
}

// Type returns...
func (e SysWMEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e SysWMEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 SysWMEvent

// {{{2 TouchFingerEvent

// TouchFingerEvent holds a EVENT
type TouchFingerEvent struct {
	ev *C.SDL_TouchFingerEvent
}

// Type returns...
func (e TouchFingerEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e TouchFingerEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 TouchFingerEvent

// {{{2 MultiGestureEvent

// MultiGestureEvent holds a EVENT
type MultiGestureEvent struct {
	ev *C.SDL_MultiGestureEvent
}

// Type returns...
func (e MultiGestureEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e MultiGestureEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 MultiGestureEvent

// {{{2 DollarGestureEvent

// DollarGestureEvent holds a EVENT
type DollarGestureEvent struct {
	ev *C.SDL_DollarGestureEvent
}

// Type returns...
func (e DollarGestureEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e DollarGestureEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 DollarGestureEvent

// {{{2 DropEvent

// DropEvent holds a EVENT
type DropEvent struct {
	ev *C.SDL_DropEvent
}

// Type returns...
func (e DropEvent) Type() EventType {
	return EventType(e.ev._type)
}

func (e DropEvent) Timestamp() uint32 {
	return uint32(e.ev.timestamp)
}

// TODO(light)

// }}}2 DropEvent

// }}}1 Event Structs
