package keys

// #include <stdlib.h>
// #include "SDL.h"
import "C"

import (
	"unsafe"
)

// Code represents keyboard keys using the current layout of the keyboard.
// These values include Unicode values representing the unmodified character
// that would be generated by pressing the key or a constant for keys that
// don't generate characters.
type Code rune

// CodeFromName returns the key code for a human-readable name
// or Unknown if name isn't recognized.
func CodeFromName(name string) Code {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Code(C.SDL_GetKeyFromName(cname))
}

// String returns a human-readable name for the key.
func (code Code) String() string {
	return C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(code)))
}

// Scancode returns the scancode corresponding to the key code
// according to the current keyboard layout.
func (code Code) Scancode() Scancode {
	return Scancode(C.SDL_GetScancodeFromKey(C.SDL_Keycode(code)))
}

// Key codes
const (
	Unknown Code = C.SDLK_UNKNOWN

	Return     Code = C.SDLK_RETURN
	Escape     Code = C.SDLK_ESCAPE
	Backspace  Code = C.SDLK_BACKSPACE
	Tab        Code = C.SDLK_TAB
	Space      Code = C.SDLK_SPACE
	Exclaim    Code = C.SDLK_EXCLAIM
	QuoteDbl   Code = C.SDLK_QUOTEDBL
	Hash       Code = C.SDLK_HASH
	Percent    Code = C.SDLK_PERCENT
	Dollar     Code = C.SDLK_DOLLAR
	Ampersand  Code = C.SDLK_AMPERSAND
	Quote      Code = C.SDLK_QUOTE
	LeftParen  Code = C.SDLK_LEFTPAREN
	RightParen Code = C.SDLK_RIGHTPAREN
	Asterisk   Code = C.SDLK_ASTERISK
	Plus       Code = C.SDLK_PLUS
	Comma      Code = C.SDLK_COMMA
	Minus      Code = C.SDLK_MINUS
	Period     Code = C.SDLK_PERIOD
	Slash      Code = C.SDLK_SLASH

	// TODO(light): digits
	K0 Code = C.SDLK_0
	K1 Code = C.SDLK_1
	K2 Code = C.SDLK_2
	K3 Code = C.SDLK_3
	K4 Code = C.SDLK_4
	K5 Code = C.SDLK_5
	K6 Code = C.SDLK_6
	K7 Code = C.SDLK_7
	K8 Code = C.SDLK_8
	K9 Code = C.SDLK_9

	Colon        Code = C.SDLK_COLON
	Semicolon    Code = C.SDLK_SEMICOLON
	Less         Code = C.SDLK_LESS
	Equals       Code = C.SDLK_EQUALS
	Greater      Code = C.SDLK_GREATER
	Question     Code = C.SDLK_QUESTION
	At           Code = C.SDLK_AT
	LeftBracket  Code = C.SDLK_LEFTBRACKET
	Backslash    Code = C.SDLK_BACKSLASH
	RightBracket Code = C.SDLK_RIGHTBRACKET
	Caret        Code = C.SDLK_CARET
	Underscore   Code = C.SDLK_UNDERSCORE
	Backquote    Code = C.SDLK_BACKQUOTE
	A            Code = C.SDLK_a
	B            Code = C.SDLK_b
	C            Code = C.SDLK_c
	D            Code = C.SDLK_d
	E            Code = C.SDLK_e
	F            Code = C.SDLK_f
	G            Code = C.SDLK_g
	H            Code = C.SDLK_h
	I            Code = C.SDLK_i
	J            Code = C.SDLK_j
	K            Code = C.SDLK_k
	L            Code = C.SDLK_l
	M            Code = C.SDLK_m
	N            Code = C.SDLK_n
	O            Code = C.SDLK_o
	P            Code = C.SDLK_p
	Q            Code = C.SDLK_q
	R            Code = C.SDLK_r
	S            Code = C.SDLK_s
	T            Code = C.SDLK_t
	U            Code = C.SDLK_u
	V            Code = C.SDLK_v
	W            Code = C.SDLK_w
	X            Code = C.SDLK_x
	Y            Code = C.SDLK_y
	Z            Code = C.SDLK_z

	CapsLock Code = C.SDLK_CAPSLOCK

	F1  Code = C.SDLK_F1
	F2  Code = C.SDLK_F2
	F3  Code = C.SDLK_F3
	F4  Code = C.SDLK_F4
	F5  Code = C.SDLK_F5
	F6  Code = C.SDLK_F6
	F7  Code = C.SDLK_F7
	F8  Code = C.SDLK_F8
	F9  Code = C.SDLK_F9
	F10 Code = C.SDLK_F10
	F11 Code = C.SDLK_F11
	F12 Code = C.SDLK_F12

	PrintScreen Code = C.SDLK_PRINTSCREEN
	ScrollLock  Code = C.SDLK_SCROLLLOCK
	Pause       Code = C.SDLK_PAUSE
	Insert      Code = C.SDLK_INSERT
	Home        Code = C.SDLK_HOME
	PageUp      Code = C.SDLK_PAGEUP
	Delete      Code = C.SDLK_DELETE
	End         Code = C.SDLK_END
	PageDown    Code = C.SDLK_PAGEDOWN
	Right       Code = C.SDLK_RIGHT
	Left        Code = C.SDLK_LEFT
	Down        Code = C.SDLK_DOWN
	Up          Code = C.SDLK_UP

	NumlockClear   Code = C.SDLK_NUMLOCKCLEAR
	KeypadDivide   Code = C.SDLK_KP_DIVIDE
	KeypadMultiply Code = C.SDLK_KP_MULTIPLY
	KeypadMinus    Code = C.SDLK_KP_MINUS
	KeypadPlus     Code = C.SDLK_KP_PLUS
	KeypadEnter    Code = C.SDLK_KP_ENTER
	Keypad1        Code = C.SDLK_KP_1
	Keypad2        Code = C.SDLK_KP_2
	Keypad3        Code = C.SDLK_KP_3
	Keypad4        Code = C.SDLK_KP_4
	Keypad5        Code = C.SDLK_KP_5
	Keypad6        Code = C.SDLK_KP_6
	Keypad7        Code = C.SDLK_KP_7
	Keypad8        Code = C.SDLK_KP_8
	Keypad9        Code = C.SDLK_KP_9
	Keypad0        Code = C.SDLK_KP_0
	KeypadPeriod   Code = C.SDLK_KP_PERIOD

	Application       Code = C.SDLK_APPLICATION
	Power             Code = C.SDLK_POWER
	KeypadEquals      Code = C.SDLK_KP_EQUALS
	F13               Code = C.SDLK_F13
	F14               Code = C.SDLK_F14
	F15               Code = C.SDLK_F15
	F16               Code = C.SDLK_F16
	F17               Code = C.SDLK_F17
	F18               Code = C.SDLK_F18
	F19               Code = C.SDLK_F19
	F20               Code = C.SDLK_F20
	F21               Code = C.SDLK_F21
	F22               Code = C.SDLK_F22
	F23               Code = C.SDLK_F23
	F24               Code = C.SDLK_F24
	Execute           Code = C.SDLK_EXECUTE
	Help              Code = C.SDLK_HELP
	Menu              Code = C.SDLK_MENU
	Select            Code = C.SDLK_SELECT
	Stop              Code = C.SDLK_STOP
	Again             Code = C.SDLK_AGAIN
	Undo              Code = C.SDLK_UNDO
	Cut               Code = C.SDLK_CUT
	Copy              Code = C.SDLK_COPY
	Paste             Code = C.SDLK_PASTE
	Find              Code = C.SDLK_FIND
	Mute              Code = C.SDLK_MUTE
	VolumeUp          Code = C.SDLK_VOLUMEUP
	VolumeDown        Code = C.SDLK_VOLUMEDOWN
	KeypadComma       Code = C.SDLK_KP_COMMA
	KeypadEqualsAS400 Code = C.SDLK_KP_EQUALSAS400

	AltErase   Code = C.SDLK_ALTERASE
	SysReq     Code = C.SDLK_SYSREQ
	Cancel     Code = C.SDLK_CANCEL
	Clear      Code = C.SDLK_CLEAR
	Prior      Code = C.SDLK_PRIOR
	Return2    Code = C.SDLK_RETURN2
	Separator  Code = C.SDLK_SEPARATOR
	Out        Code = C.SDLK_OUT
	Oper       Code = C.SDLK_OPER
	ClearAgain Code = C.SDLK_CLEARAGAIN
	CrSel      Code = C.SDLK_CRSEL
	ExSel      Code = C.SDLK_EXSEL

	Keypad00             Code = C.SDLK_KP_00
	Keypad000            Code = C.SDLK_KP_000
	ThousandsSeparator   Code = C.SDLK_THOUSANDSSEPARATOR
	DecimalSeparator     Code = C.SDLK_DECIMALSEPARATOR
	CurrencyUnit         Code = C.SDLK_CURRENCYUNIT
	CurrencySubUnit      Code = C.SDLK_CURRENCYSUBUNIT
	KeypadLeftParen      Code = C.SDLK_KP_LEFTPAREN
	KeypadRightParen     Code = C.SDLK_KP_RIGHTPAREN
	KeypadLeftBrace      Code = C.SDLK_KP_LEFTBRACE
	KeypadRightBrace     Code = C.SDLK_KP_RIGHTBRACE
	KeypadTab            Code = C.SDLK_KP_TAB
	KeypadBackspace      Code = C.SDLK_KP_BACKSPACE
	KeypadA              Code = C.SDLK_KP_A
	KeypadB              Code = C.SDLK_KP_B
	KeypadC              Code = C.SDLK_KP_C
	KeypadD              Code = C.SDLK_KP_D
	KeypadE              Code = C.SDLK_KP_E
	KeypadF              Code = C.SDLK_KP_F
	KeypadXOR            Code = C.SDLK_KP_XOR
	KeypadPower          Code = C.SDLK_KP_POWER
	KeypadPercent        Code = C.SDLK_KP_PERCENT
	KeypadLess           Code = C.SDLK_KP_LESS
	KeypadGreater        Code = C.SDLK_KP_GREATER
	KeypadAmpersand      Code = C.SDLK_KP_AMPERSAND
	KeypadDblAmpersand   Code = C.SDLK_KP_DBLAMPERSAND
	KeypadVerticalBar    Code = C.SDLK_KP_VERTICALBAR
	KeypadDblVerticalBar Code = C.SDLK_KP_DBLVERTICALBAR
	KeypadColon          Code = C.SDLK_KP_COLON
	KeypadHash           Code = C.SDLK_KP_HASH
	KeypadSpace          Code = C.SDLK_KP_SPACE
	KeypadAt             Code = C.SDLK_KP_AT
	KeypadExclam         Code = C.SDLK_KP_EXCLAM
	KeypadMemStore       Code = C.SDLK_KP_MEMSTORE
	KeypadMemRecall      Code = C.SDLK_KP_MEMRECALL
	KeypadMemClear       Code = C.SDLK_KP_MEMCLEAR
	KeypadMemAdd         Code = C.SDLK_KP_MEMADD
	KeypadMemSubtract    Code = C.SDLK_KP_MEMSUBTRACT
	KeypadMemMultiply    Code = C.SDLK_KP_MEMMULTIPLY
	KeypadMemDivide      Code = C.SDLK_KP_MEMDIVIDE
	KeypadPlusMinus      Code = C.SDLK_KP_PLUSMINUS
	KeypadClear          Code = C.SDLK_KP_CLEAR
	KeypadClearEntry     Code = C.SDLK_KP_CLEARENTRY
	KeypadBinary         Code = C.SDLK_KP_BINARY
	KeypadOctal          Code = C.SDLK_KP_OCTAL
	KeypadDecimal        Code = C.SDLK_KP_DECIMAL
	KeypadHexadecimal    Code = C.SDLK_KP_HEXADECIMAL

	LCtrl  Code = C.SDLK_LCTRL
	LShift Code = C.SDLK_LSHIFT
	LAlt   Code = C.SDLK_LALT
	LGUI   Code = C.SDLK_LGUI
	RCtrl  Code = C.SDLK_RCTRL
	RShift Code = C.SDLK_RSHIFT
	RAlt   Code = C.SDLK_RALT
	RGUI   Code = C.SDLK_RGUI

	Mode Code = C.SDLK_MODE

	AudioNext           Code = C.SDLK_AUDIONEXT
	AudioPrev           Code = C.SDLK_AUDIOPREV
	AudioStop           Code = C.SDLK_AUDIOSTOP
	AudioPlay           Code = C.SDLK_AUDIOPLAY
	AudioMute           Code = C.SDLK_AUDIOMUTE
	MediaSelect         Code = C.SDLK_MEDIASELECT
	WWW                 Code = C.SDLK_WWW
	Mail                Code = C.SDLK_MAIL
	Calculator          Code = C.SDLK_CALCULATOR
	Computer            Code = C.SDLK_COMPUTER
	AppControlSearch    Code = C.SDLK_AC_SEARCH
	AppControlHome      Code = C.SDLK_AC_HOME
	AppControlBack      Code = C.SDLK_AC_BACK
	AppControlForward   Code = C.SDLK_AC_FORWARD
	AppControlStop      Code = C.SDLK_AC_STOP
	AppControlRefresh   Code = C.SDLK_AC_REFRESH
	AppControlBookmarks Code = C.SDLK_AC_BOOKMARKS

	BrightnessDown      Code = C.SDLK_BRIGHTNESSDOWN
	BrightnessUp        Code = C.SDLK_BRIGHTNESSUP
	DisplaySwitch       Code = C.SDLK_DISPLAYSWITCH
	KeyboardIllumToggle Code = C.SDLK_KBDILLUMTOGGLE
	KeyboardIllumDown   Code = C.SDLK_KBDILLUMDOWN
	KeyboardIllumUp     Code = C.SDLK_KBDILLUMUP
	Eject               Code = C.SDLK_EJECT
	Sleep               Code = C.SDLK_SLEEP
)
