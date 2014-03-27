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
		return &commonEvent{
			tp:   EventType(common._type),
			time: uint32(common.timestamp),
		}
	case WindowEventType:
		ce := (*C.SDL_WindowEvent)(cEvent)
		return &WindowEvent{
			Time:     uint32(ce.timestamp),
			WindowID: uint32(ce.windowID),
			Event:    WindowEventID(ce.event),
			Data1:    int32(ce.data1),
			Data2:    int32(ce.data2),
		}
	case KeyDownEventType, KeyUpEventType:
		ce := (*C.SDL_KeyboardEvent)(cEvent)
		return &KeyboardEvent{
			Time:     uint32(ce.timestamp),
			WindowID: uint32(ce.windowID),
			Pressed:  ce.state == C.SDL_PRESSED,
			Repeat:   ce.repeat != 0,
			KeySym: KeySym{
				Scancode: keys.Scancode(ce.keysym.scancode),
				Code:     keys.Code(ce.keysym.sym),
				Mod:      keys.Mod(ce.keysym.mod),
			},
		}
	case TextEditingEventType:
		ce := (*C.SDL_TextEditingEvent)(cEvent)
		return &TextEditingEvent{
			Time:     uint32(ce.timestamp),
			WindowID: uint32(ce.windowID),
			Text:     C.GoString(&ce.text[0]),
			Start:    int(ce.start),
			Length:   int(ce.length),
		}
	case TextInputEventType:
		ce := (*C.SDL_TextInputEvent)(cEvent)
		return &TextInputEvent{
			Time:     uint32(ce.timestamp),
			WindowID: uint32(ce.windowID),
			Text:     C.GoString(&ce.text[0]),
		}
	case MouseMotionEventType:
		ce := (*C.SDL_MouseMotionEvent)(cEvent)
		return &MouseMotionEvent{
			Time:     uint32(ce.timestamp),
			WindowID: uint32(ce.windowID),
			Which:    uint32(ce.which),
			State:    uint32(ce.state),
			X:        int32(ce.x),
			Y:        int32(ce.y),
			RelX:     int32(ce.xrel),
			RelY:     int32(ce.yrel),
		}
	case MouseButtonDownEventType, MouseButtonUpEventType:
		ce := (*C.SDL_MouseButtonEvent)(cEvent)
		return &MouseButtonEvent{
			Time:     uint32(ce.timestamp),
			WindowID: uint32(ce.windowID),
			Which:    uint32(ce.which),
			Button:   MouseButton(ce.button),
			Pressed:  ce.state == C.SDL_PRESSED,
			X:        int32(ce.x),
			Y:        int32(ce.y),
		}
	case MouseWheelEventType:
		ce := (*C.SDL_MouseWheelEvent)(cEvent)
		return &MouseWheelEvent{
			Time:     uint32(ce.timestamp),
			WindowID: uint32(ce.windowID),
			Which:    uint32(ce.which),
			X:        int32(ce.x),
			Y:        int32(ce.y),
		}
	case JoyAxisMotionEventType:
		ce := (*C.SDL_JoyAxisEvent)(cEvent)
		return &JoyAxisEvent{
			Time:  uint32(ce.timestamp),
			Which: JoystickID(ce.which),
			Axis:  uint8(ce.axis),
			Value: int16(ce.value),
		}
	case JoyBallMotionEventType:
		ce := (*C.SDL_JoyBallEvent)(cEvent)
		return &JoyBallEvent{
			Time:  uint32(ce.timestamp),
			Which: JoystickID(ce.which),
			Ball:  uint8(ce.ball),
			RelX:  int16(ce.xrel),
			RelY:  int16(ce.yrel),
		}
	case JoyHatMotionEventType:
		ce := (*C.SDL_JoyHatEvent)(cEvent)
		return &JoyHatEvent{
			Time:     uint32(ce.timestamp),
			Which:    JoystickID(ce.which),
			Hat:      uint8(ce.hat),
			Position: HatPosition(ce.value),
		}
	case JoyButtonDownEventType, JoyButtonUpEventType:
		ce := (*C.SDL_JoyButtonEvent)(cEvent)
		return &JoyButtonEvent{
			Time:    uint32(ce.timestamp),
			Which:   JoystickID(ce.which),
			Button:  uint8(ce.button),
			Pressed: ce.state == C.SDL_PRESSED,
		}
	case JoyDeviceAddedEventType, JoyDeviceRemovedEventType:
		ce := (*C.SDL_JoyDeviceEvent)(cEvent)
		return &JoyDeviceEvent{
			Time:  uint32(ce.timestamp),
			Which: int32(ce.which),
			Added: EventType(ce._type) == JoyDeviceAddedEventType,
		}
	case ControllerAxisMotionEventType:
		ce := (*C.SDL_ControllerAxisEvent)(cEvent)
		return &ControllerAxisEvent{
			Time:  uint32(ce.timestamp),
			Which: JoystickID(ce.which),
			Axis:  uint8(ce.axis),
			Value: int16(ce.value),
		}
	case ControllerButtonDownEventType, ControllerButtonUpEventType:
		ce := (*C.SDL_ControllerButtonEvent)(cEvent)
		return &ControllerButtonEvent{
			Time:    uint32(ce.timestamp),
			Which:   JoystickID(ce.which),
			Button:  uint8(ce.button),
			Pressed: ce.state == C.SDL_PRESSED,
		}
	case ControllerDeviceAddedEventType, ControllerDeviceRemovedEventType, ControllerDeviceRemappedEventType:
		ce := (*C.SDL_ControllerDeviceEvent)(cEvent)
		return &ControllerDeviceEvent{
			EventType: EventType(ce._type),
			Time:      uint32(ce.timestamp),
			Which:     int32(ce.which),
		}
	}
	if EventType(common._type).IsUserEvent() {
		ce := (*C.SDL_UserEvent)(cEvent)
		return &UserEvent{
			EventType: EventType(ce._type),
			Time:      uint32(ce.timestamp),
			WindowID:  uint32(ce.windowID),
			Code:      int32(ce.code),
			Data1:     ce.data1,
			Data2:     ce.data2,
		}
	}
	fmt.Println("Unhandled event type:", EventType(common._type))
	return &commonEvent{
		tp:   EventType(common._type),
		time: uint32(common.timestamp),
	}
}

// {{{1 Event Structs

// {{{2 CommonEvent

// commonEvent holds fields common to all events.  It is not exported because it
// doesn't provide anything useful outside of the Event interface.
type commonEvent struct {
	tp   EventType
	time uint32
}

func (e *commonEvent) Type() EventType {
	return e.tp
}

func (e *commonEvent) Timestamp() uint32 {
	return e.time
}

// }}}2 CommonEvent

// {{{2 WindowEvent

// WindowEvent holds window state change event data.
type WindowEvent struct {
	Time     uint32
	WindowID uint32
	Event    WindowEventID

	// For move events, this is the new (x, y) position of the window.
	// For resize events, this is the new window size.
	Data1, Data2 int32
}

// EventType returns WindowEventType.
func (e *WindowEvent) Type() EventType {
	return WindowEventType
}

func (e *WindowEvent) Timestamp() uint32 {
	return e.Time
}

// Window returns the ID of the window that this event occurred in.
func (e *WindowEvent) Window() uint32 {
	return e.WindowID
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

// KeyboardEvent holds a key press or key release event.
type KeyboardEvent struct {
	Time     uint32
	WindowID uint32
	Pressed  bool
	Repeat   bool
	KeySym
}

func (e *KeyboardEvent) Type() EventType {
	if e.Pressed {
		return KeyDownEventType
	} else {
		return KeyUpEventType
	}
}

func (e *KeyboardEvent) Timestamp() uint32 {
	return e.Time
}

func (e *KeyboardEvent) Window() uint32 {
	return e.WindowID
}

// KeySym holds the keyboard information from a keyboard event.
type KeySym struct {
	Scancode keys.Scancode
	Code     keys.Code
	Mod      keys.Mod
}

// }}}2 KeyboardEvent

// {{{2 TextEditingEvent

// TextEditingEvent holds a partial text input event.  See the description of TextInputEvent.
type TextEditingEvent struct {
	Time     uint32
	WindowID uint32
	Text     string
	Start    int // location to begin editing from
	Length   int // number of characters to edit
}

// Type returns TextEditingEventType.
func (e *TextEditingEvent) Type() EventType {
	return TextEditingEventType
}

func (e *TextEditingEvent) Timestamp() uint32 {
	return e.Time
}

// Window returns the window with keyboard focus or zero.
func (e *TextEditingEvent) Window() uint32 {
	return e.WindowID
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
	Time     uint32
	WindowID uint32
	Text     string
}

// Type returns TextInputEventType.
func (e *TextInputEvent) Type() EventType {
	return TextInputEventType
}

func (e *TextInputEvent) Timestamp() uint32 {
	return e.Time
}

// Window returns the window with keyboard focus or zero.
func (e *TextInputEvent) Window() uint32 {
	return e.WindowID
}

// }}}2 TextInputEvent

// {{{2 MouseMotionEvent

// MouseMotionEvent holds a mouse movement event.
type MouseMotionEvent struct {
	Time       uint32
	WindowID   uint32
	Which      uint32 // mouse that triggered the event
	State      uint32
	X, Y       int32
	RelX, RelY int32
}

// Type returns MouseMotionEventType.
func (e *MouseMotionEvent) Type() EventType {
	return MouseMotionEventType
}

func (e *MouseMotionEvent) Timestamp() uint32 {
	return e.Time
}

// WindowID returns the window with mouse focus, or zero if no window has focus.
func (e *MouseMotionEvent) Window() uint32 {
	return e.WindowID
}

// }}}2 MouseMotionEvent

// {{{2 MouseButtonEvent

// MouseButtonEvent holds a mouse button press or release event.
type MouseButtonEvent struct {
	Time     uint32
	WindowID uint32
	Which    uint32
	Button   MouseButton
	Pressed  bool
	// TODO(light): this is only available in SDL 2.0.2 and above
	// Clicks uint8 // number of clicks in sequence: 1 for single-click, 2 for double-click, etc.
	X, Y int32
}

// Type returns either MouseButtonDownEventType or MouseButtonUpEventType.
func (e *MouseButtonEvent) Type() EventType {
	if e.Pressed {
		return MouseButtonDownEventType
	} else {
		return MouseButtonUpEventType
	}
}

func (e *MouseButtonEvent) Timestamp() uint32 {
	return e.Time
}

// Window returns the window with mouse focus, or zero if no window has focus.
func (e *MouseButtonEvent) Window() uint32 {
	return e.WindowID
}

// IsTouch reports whether this event was generated by a touch input device.
func (e *MouseButtonEvent) IsTouch() bool {
	return e.Which == C.SDL_TOUCH_MOUSEID
}

// }}}2 MouseButtonEvent

// {{{2 MouseWheelEvent

// MouseWheelEvent holds a mouse wheel movement event.
type MouseWheelEvent struct {
	Time     uint32
	WindowID uint32
	Which    uint32
	X, Y     int32 // Scroll delta. The axes increase right and up.
}

// Type returns MouseWheelEventType.
func (e *MouseWheelEvent) Type() EventType {
	return MouseWheelEventType
}

func (e *MouseWheelEvent) Timestamp() uint32 {
	return e.Time
}

// Window returns the window with mouse focus, or zero if no window has focus.
func (e *MouseWheelEvent) Window() uint32 {
	return e.WindowID
}

// }}}2 MouseWheelEvent

// {{{2 JoyAxisEvent

// JoyAxisEvent holds a joystick axis movement event.
type JoyAxisEvent struct {
	Time  uint32
	Which JoystickID
	Axis  uint8
	Value int16
}

// Type returns JoyAxisMotionEventType.
func (e *JoyAxisEvent) Type() EventType {
	return JoyAxisMotionEventType
}

func (e *JoyAxisEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 JoyAxisEvent

// {{{2 JoyBallEvent

// JoyBallEvent holds a joystick trackball motion event.
type JoyBallEvent struct {
	Time       uint32
	Which      JoystickID
	Ball       uint8
	RelX, RelY int16
}

// Type returns JoyBallMotionEventType.
func (e *JoyBallEvent) Type() EventType {
	return JoyBallMotionEventType
}

func (e *JoyBallEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 JoyBallEvent

// {{{2 JoyHatEvent

// JoyHatEvent holds a joystick hat movement event.
type JoyHatEvent struct {
	Time     uint32
	Which    JoystickID
	Hat      uint8
	Position HatPosition
}

// Type returns JoyHatMotionEventType.
func (e *JoyHatEvent) Type() EventType {
	return JoyHatMotionEventType
}

func (e *JoyHatEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 JoyHatEvent

// {{{2 JoyButtonEvent

// JoyButtonEvent holds a EVENT
type JoyButtonEvent struct {
	Time    uint32
	Which   JoystickID
	Button  uint8
	Pressed bool
}

// Type returns either JoyButtonDownEventType or JoyButtonUpEventType.
func (e *JoyButtonEvent) Type() EventType {
	if e.Pressed {
		return JoyButtonDownEventType
	} else {
		return JoyButtonUpEventType
	}
}

func (e *JoyButtonEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 JoyButtonEvent

// {{{2 JoyDeviceEvent

// JoyDeviceEvent holds a joystick connection or disconnection event.
type JoyDeviceEvent struct {
	Time  uint32
	Which int32 // joystick device index for an added event or instance ID for a removal event.
	Added bool
}

// Type returns either JoyDeviceAddedEventType or JoyDeviceRemovedEventType.
func (e *JoyDeviceEvent) Type() EventType {
	if e.Added {
		return JoyDeviceAddedEventType
	} else {
		return JoyDeviceRemovedEventType
	}
}

func (e *JoyDeviceEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 JoyDeviceEvent

// {{{2 ControllerAxisEvent

// ControllerAxisEvent holds a controller axis movement event.
type ControllerAxisEvent struct {
	Time  uint32
	Which JoystickID
	Axis  uint8 // TODO(light): GameControllerAxis
	Value int16
}

// Type returns ControllerAxisMotionEventType.
func (e *ControllerAxisEvent) Type() EventType {
	return ControllerAxisMotionEventType
}

func (e *ControllerAxisEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 ControllerAxisEvent

// {{{2 ControllerButtonEvent

// ControllerButtonEvent holds a game controller button event.
type ControllerButtonEvent struct {
	Time    uint32
	Which   JoystickID
	Button  uint8 // TODO(light): GameControllerButton
	Pressed bool
}

// Type returns ControllerButtonDownEventType or ControllerButtonUpEventType.
func (e *ControllerButtonEvent) Type() EventType {
	if e.Pressed {
		return ControllerButtonDownEventType
	} else {
		return ControllerButtonUpEventType
	}
}

func (e *ControllerButtonEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 ControllerButtonEvent

// {{{2 ControllerDeviceEvent

// ControllerDeviceEvent holds a game controller device change event.
type ControllerDeviceEvent struct {
	EventType EventType
	Time      uint32
	Which     int32 // the device index for add events, otherwise the instance ID
}

// Type returns one of ControllerDeviceAddedEventType, ControllerDeviceRemovedEventType, or ControllerDeviceRemappedEventType.
func (e *ControllerDeviceEvent) Type() EventType {
	return e.EventType
}

func (e *ControllerDeviceEvent) Timestamp() uint32 {
	return e.Time
}

// }}}2 ControllerDeviceEvent

// {{{2 UserEvent

// UserEvent holds a user-defined event.
type UserEvent struct {
	EventType    EventType
	Time         uint32
	WindowID     uint32
	Code         int32
	Data1, Data2 unsafe.Pointer
}

// Type returns the event's type.
func (e *UserEvent) Type() EventType {
	return e.EventType
}

func (e *UserEvent) Timestamp() uint32 {
	return e.Time
}

// Window returns the associated window ID or zero.
func (e *UserEvent) Window() uint32 {
	return e.WindowID
}

// }}}2 UserEvent

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
