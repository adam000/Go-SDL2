package keys

// #include <stdlib.h>
// #include "SDL.h"
import "C"

import (
	"unsafe"
)

// Scancode represents a keyboard key.
type Scancode int32

// ScancodeFromName returns the scancode for a human-readable name
// or ScancodeUnknown if name isn't recognized.
func ScancodeFromName(name string) Scancode {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Scancode(C.SDL_GetScancodeFromName(cname))
}

// String returns a human-readable name for the key.
func (code Scancode) String() string {
	return C.GoString(C.SDL_GetScancodeName(C.SDL_Scancode(code)))
}

// Code returns the key code corresponding to the scancode
// according to the current keyboard layout.
func (scode Scancode) Code() Code {
	return Code(C.SDL_GetKeyFromScancode(C.SDL_Scancode(scode)))
}

// Scan codes
const (
    ScancodeUnknown Scancode = C.SDL_SCANCODE_UNKNOWN

    ScancodeA Scancode = C.SDL_SCANCODE_A
    ScancodeB Scancode = C.SDL_SCANCODE_B
    ScancodeC Scancode = C.SDL_SCANCODE_C
    ScancodeD Scancode = C.SDL_SCANCODE_D
    ScancodeE Scancode = C.SDL_SCANCODE_E
    ScancodeF Scancode = C.SDL_SCANCODE_F
    ScancodeG Scancode = C.SDL_SCANCODE_G
    ScancodeH Scancode = C.SDL_SCANCODE_H
    ScancodeI Scancode = C.SDL_SCANCODE_I
    ScancodeJ Scancode = C.SDL_SCANCODE_J
    ScancodeK Scancode = C.SDL_SCANCODE_K
    ScancodeL Scancode = C.SDL_SCANCODE_L
    ScancodeM Scancode = C.SDL_SCANCODE_M
    ScancodeN Scancode = C.SDL_SCANCODE_N
    ScancodeO Scancode = C.SDL_SCANCODE_O
    ScancodeP Scancode = C.SDL_SCANCODE_P
    ScancodeQ Scancode = C.SDL_SCANCODE_Q
    ScancodeR Scancode = C.SDL_SCANCODE_R
    ScancodeS Scancode = C.SDL_SCANCODE_S
    ScancodeT Scancode = C.SDL_SCANCODE_T
    ScancodeU Scancode = C.SDL_SCANCODE_U
    ScancodeV Scancode = C.SDL_SCANCODE_V
    ScancodeW Scancode = C.SDL_SCANCODE_W
    ScancodeX Scancode = C.SDL_SCANCODE_X
    ScancodeY Scancode = C.SDL_SCANCODE_Y
    ScancodeZ Scancode = C.SDL_SCANCODE_Z

    Scancode1 Scancode = C.SDL_SCANCODE_1
    Scancode2 Scancode = C.SDL_SCANCODE_2
    Scancode3 Scancode = C.SDL_SCANCODE_3
    Scancode4 Scancode = C.SDL_SCANCODE_4
    Scancode5 Scancode = C.SDL_SCANCODE_5
    Scancode6 Scancode = C.SDL_SCANCODE_6
    Scancode7 Scancode = C.SDL_SCANCODE_7
    Scancode8 Scancode = C.SDL_SCANCODE_8
    Scancode9 Scancode = C.SDL_SCANCODE_9
    Scancode0 Scancode = C.SDL_SCANCODE_0

    ScancodeReturn Scancode = C.SDL_SCANCODE_RETURN
    ScancodeEscape Scancode = C.SDL_SCANCODE_ESCAPE
    ScancodeBackspace Scancode = C.SDL_SCANCODE_BACKSPACE
    ScancodeTab Scancode = C.SDL_SCANCODE_TAB
    ScancodeSpace Scancode = C.SDL_SCANCODE_SPACE

    ScancodeMinus Scancode = C.SDL_SCANCODE_MINUS
    ScancodeEquals Scancode = C.SDL_SCANCODE_EQUALS
    ScancodeLeftBracket Scancode = C.SDL_SCANCODE_LEFTBRACKET
    ScancodeRightBracket Scancode = C.SDL_SCANCODE_RIGHTBRACKET
    ScancodeBackslash Scancode = C.SDL_SCANCODE_BACKSLASH
    ScancodeNonUSHash Scancode = C.SDL_SCANCODE_NONUSHASH
    ScancodeSemicolon Scancode = C.SDL_SCANCODE_SEMICOLON
    ScancodeApostrophe Scancode = C.SDL_SCANCODE_APOSTROPHE
    ScancodeGrave Scancode = C.SDL_SCANCODE_GRAVE
    ScancodeComma Scancode = C.SDL_SCANCODE_COMMA
    ScancodePeriod Scancode = C.SDL_SCANCODE_PERIOD
    ScancodeSlash Scancode = C.SDL_SCANCODE_SLASH

    ScancodeCapsLock Scancode = C.SDL_SCANCODE_CAPSLOCK

    ScancodeF1 Scancode = C.SDL_SCANCODE_F1
    ScancodeF2 Scancode = C.SDL_SCANCODE_F2
    ScancodeF3 Scancode = C.SDL_SCANCODE_F3
    ScancodeF4 Scancode = C.SDL_SCANCODE_F4
    ScancodeF5 Scancode = C.SDL_SCANCODE_F5
    ScancodeF6 Scancode = C.SDL_SCANCODE_F6
    ScancodeF7 Scancode = C.SDL_SCANCODE_F7
    ScancodeF8 Scancode = C.SDL_SCANCODE_F8
    ScancodeF9 Scancode = C.SDL_SCANCODE_F9
    ScancodeF10 Scancode = C.SDL_SCANCODE_F10
    ScancodeF11 Scancode = C.SDL_SCANCODE_F11
    ScancodeF12 Scancode = C.SDL_SCANCODE_F12

    ScancodePrintScreen Scancode = C.SDL_SCANCODE_PRINTSCREEN
    ScancodeScrollLock Scancode = C.SDL_SCANCODE_SCROLLLOCK
    ScancodePause Scancode = C.SDL_SCANCODE_PAUSE
    ScancodeInsert Scancode = C.SDL_SCANCODE_INSERT
    ScancodeHome Scancode = C.SDL_SCANCODE_HOME
    ScancodePageUp Scancode = C.SDL_SCANCODE_PAGEUP
    ScancodeDelete Scancode = C.SDL_SCANCODE_DELETE
    ScancodeEnd Scancode = C.SDL_SCANCODE_END
    ScancodePageDown Scancode = C.SDL_SCANCODE_PAGEDOWN
    ScancodeRight Scancode = C.SDL_SCANCODE_RIGHT
    ScancodeLeft Scancode = C.SDL_SCANCODE_LEFT
    ScancodeDown Scancode = C.SDL_SCANCODE_DOWN
    ScancodeUp Scancode = C.SDL_SCANCODE_UP

    ScancodeNumlockClear Scancode = C.SDL_SCANCODE_NUMLOCKCLEAR
    ScancodeKeypadDivide Scancode = C.SDL_SCANCODE_KP_DIVIDE
    ScancodeKeypadMultiply Scancode = C.SDL_SCANCODE_KP_MULTIPLY
    ScancodeKeypadMinus Scancode = C.SDL_SCANCODE_KP_MINUS
    ScancodeKeypadPlus Scancode = C.SDL_SCANCODE_KP_PLUS
    ScancodeKeypadEnter Scancode = C.SDL_SCANCODE_KP_ENTER
    ScancodeKeypad1 Scancode = C.SDL_SCANCODE_KP_1
    ScancodeKeypad2 Scancode = C.SDL_SCANCODE_KP_2
    ScancodeKeypad3 Scancode = C.SDL_SCANCODE_KP_3
    ScancodeKeypad4 Scancode = C.SDL_SCANCODE_KP_4
    ScancodeKeypad5 Scancode = C.SDL_SCANCODE_KP_5
    ScancodeKeypad6 Scancode = C.SDL_SCANCODE_KP_6
    ScancodeKeypad7 Scancode = C.SDL_SCANCODE_KP_7
    ScancodeKeypad8 Scancode = C.SDL_SCANCODE_KP_8
    ScancodeKeypad9 Scancode = C.SDL_SCANCODE_KP_9
    ScancodeKeypad0 Scancode = C.SDL_SCANCODE_KP_0
    ScancodeKeypadPeriod Scancode = C.SDL_SCANCODE_KP_PERIOD

    ScancodeNonUSBackslash Scancode = C.SDL_SCANCODE_NONUSBACKSLASH
    ScancodeApplication Scancode = C.SDL_SCANCODE_APPLICATION
    ScancodePower Scancode = C.SDL_SCANCODE_POWER
    ScancodeKeypadEquals Scancode = C.SDL_SCANCODE_KP_EQUALS
    ScancodeF13 Scancode = C.SDL_SCANCODE_F13
    ScancodeF14 Scancode = C.SDL_SCANCODE_F14
    ScancodeF15 Scancode = C.SDL_SCANCODE_F15
    ScancodeF16 Scancode = C.SDL_SCANCODE_F16
    ScancodeF17 Scancode = C.SDL_SCANCODE_F17
    ScancodeF18 Scancode = C.SDL_SCANCODE_F18
    ScancodeF19 Scancode = C.SDL_SCANCODE_F19
    ScancodeF20 Scancode = C.SDL_SCANCODE_F20
    ScancodeF21 Scancode = C.SDL_SCANCODE_F21
    ScancodeF22 Scancode = C.SDL_SCANCODE_F22
    ScancodeF23 Scancode = C.SDL_SCANCODE_F23
    ScancodeF24 Scancode = C.SDL_SCANCODE_F24
    ScancodeExecute Scancode = C.SDL_SCANCODE_EXECUTE
    ScancodeHelp Scancode = C.SDL_SCANCODE_HELP
    ScancodeMenu Scancode = C.SDL_SCANCODE_MENU
    ScancodeSelect Scancode = C.SDL_SCANCODE_SELECT
    ScancodeStop Scancode = C.SDL_SCANCODE_STOP
    ScancodeAgain Scancode = C.SDL_SCANCODE_AGAIN
    ScancodeUndo Scancode = C.SDL_SCANCODE_UNDO
    ScancodeCut Scancode = C.SDL_SCANCODE_CUT
    ScancodeCopy Scancode = C.SDL_SCANCODE_COPY
    ScancodePaste Scancode = C.SDL_SCANCODE_PASTE
    ScancodeFind Scancode = C.SDL_SCANCODE_FIND
    ScancodeMute Scancode = C.SDL_SCANCODE_MUTE
    ScancodeVolumeUp Scancode = C.SDL_SCANCODE_VOLUMEUP
    ScancodeVolumeDown Scancode = C.SDL_SCANCODE_VOLUMEDOWN
    ScancodeKeypadComma Scancode = C.SDL_SCANCODE_KP_COMMA
    ScancodeKeypadEqualsAS400 Scancode = C.SDL_SCANCODE_KP_EQUALSAS400

    ScancodeInternational1 Scancode = C.SDL_SCANCODE_INTERNATIONAL1
    ScancodeInternational2 Scancode = C.SDL_SCANCODE_INTERNATIONAL2
    ScancodeInternational3 Scancode = C.SDL_SCANCODE_INTERNATIONAL3
    ScancodeInternational4 Scancode = C.SDL_SCANCODE_INTERNATIONAL4
    ScancodeInternational5 Scancode = C.SDL_SCANCODE_INTERNATIONAL5
    ScancodeInternational6 Scancode = C.SDL_SCANCODE_INTERNATIONAL6
    ScancodeInternational7 Scancode = C.SDL_SCANCODE_INTERNATIONAL7
    ScancodeInternational8 Scancode = C.SDL_SCANCODE_INTERNATIONAL8
    ScancodeInternational9 Scancode = C.SDL_SCANCODE_INTERNATIONAL9
    ScancodeLang1 Scancode = C.SDL_SCANCODE_LANG1
    ScancodeLang2 Scancode = C.SDL_SCANCODE_LANG2
    ScancodeLang3 Scancode = C.SDL_SCANCODE_LANG3
    ScancodeLang4 Scancode = C.SDL_SCANCODE_LANG4
    ScancodeLang5 Scancode = C.SDL_SCANCODE_LANG5
    ScancodeLang6 Scancode = C.SDL_SCANCODE_LANG6
    ScancodeLang7 Scancode = C.SDL_SCANCODE_LANG7
    ScancodeLang8 Scancode = C.SDL_SCANCODE_LANG8
    ScancodeLang9 Scancode = C.SDL_SCANCODE_LANG9

    ScancodeAltErase Scancode = C.SDL_SCANCODE_ALTERASE
    ScancodeSysReq Scancode = C.SDL_SCANCODE_SYSREQ
    ScancodeCancel Scancode = C.SDL_SCANCODE_CANCEL
    ScancodeClear Scancode = C.SDL_SCANCODE_CLEAR
    ScancodePrior Scancode = C.SDL_SCANCODE_PRIOR
    ScancodeReturn2 Scancode = C.SDL_SCANCODE_RETURN2
    ScancodeSeparator Scancode = C.SDL_SCANCODE_SEPARATOR
    ScancodeOut Scancode = C.SDL_SCANCODE_OUT
    ScancodeOper Scancode = C.SDL_SCANCODE_OPER
    ScancodeClearAgain Scancode = C.SDL_SCANCODE_CLEARAGAIN
    ScancodeCrSel Scancode = C.SDL_SCANCODE_CRSEL
    ScancodeExSel Scancode = C.SDL_SCANCODE_EXSEL

    ScancodeKeypad00 Scancode = C.SDL_SCANCODE_KP_00
    ScancodeKeypad000 Scancode = C.SDL_SCANCODE_KP_000
    ScancodeThousandsSeparator Scancode = C.SDL_SCANCODE_THOUSANDSSEPARATOR
    ScancodeDecimalSeparator Scancode = C.SDL_SCANCODE_DECIMALSEPARATOR
    ScancodeCurrencyUnit Scancode = C.SDL_SCANCODE_CURRENCYUNIT
    ScancodeCurrencysubUnit Scancode = C.SDL_SCANCODE_CURRENCYSUBUNIT
    ScancodeKeypadLeftParen Scancode = C.SDL_SCANCODE_KP_LEFTPAREN
    ScancodeKeypadRightParen Scancode = C.SDL_SCANCODE_KP_RIGHTPAREN
    ScancodeKeypadLeftBrace Scancode = C.SDL_SCANCODE_KP_LEFTBRACE
    ScancodeKeypadRightBrace Scancode = C.SDL_SCANCODE_KP_RIGHTBRACE
    ScancodeKeypadTab Scancode = C.SDL_SCANCODE_KP_TAB
    ScancodeKeypadBackspace Scancode = C.SDL_SCANCODE_KP_BACKSPACE
    ScancodeKeypadA Scancode = C.SDL_SCANCODE_KP_A
    ScancodeKeypadB Scancode = C.SDL_SCANCODE_KP_B
    ScancodeKeypadC Scancode = C.SDL_SCANCODE_KP_C
    ScancodeKeypadD Scancode = C.SDL_SCANCODE_KP_D
    ScancodeKeypadE Scancode = C.SDL_SCANCODE_KP_E
    ScancodeKeypadF Scancode = C.SDL_SCANCODE_KP_F
    ScancodeKeypadXOR Scancode = C.SDL_SCANCODE_KP_XOR
    ScancodeKeypadPower Scancode = C.SDL_SCANCODE_KP_POWER
    ScancodeKeypadPercent Scancode = C.SDL_SCANCODE_KP_PERCENT
    ScancodeKeypadLess Scancode = C.SDL_SCANCODE_KP_LESS
    ScancodeKeypadGreater Scancode = C.SDL_SCANCODE_KP_GREATER
    ScancodeKeypadAmpersand Scancode = C.SDL_SCANCODE_KP_AMPERSAND
    ScancodeKeypadDblAmpersand Scancode = C.SDL_SCANCODE_KP_DBLAMPERSAND
    ScancodeKeypadVerticalBar Scancode = C.SDL_SCANCODE_KP_VERTICALBAR
    ScancodeKeypadDblVerticalBar Scancode = C.SDL_SCANCODE_KP_DBLVERTICALBAR
    ScancodeKeypadColon Scancode = C.SDL_SCANCODE_KP_COLON
    ScancodeKeypadHash Scancode = C.SDL_SCANCODE_KP_HASH
    ScancodeKeypadSpace Scancode = C.SDL_SCANCODE_KP_SPACE
    ScancodeKeypadAt Scancode = C.SDL_SCANCODE_KP_AT
    ScancodeKeypadExclam Scancode = C.SDL_SCANCODE_KP_EXCLAM
    ScancodeKeypadMemStore Scancode = C.SDL_SCANCODE_KP_MEMSTORE
    ScancodeKeypadMemRecall Scancode = C.SDL_SCANCODE_KP_MEMRECALL
    ScancodeKeypadMemClear Scancode = C.SDL_SCANCODE_KP_MEMCLEAR
    ScancodeKeypadMemAdd Scancode = C.SDL_SCANCODE_KP_MEMADD
    ScancodeKeypadMemSubtract Scancode = C.SDL_SCANCODE_KP_MEMSUBTRACT
    ScancodeKeypadMemMultiply Scancode = C.SDL_SCANCODE_KP_MEMMULTIPLY
    ScancodeKeypadMemDivide Scancode = C.SDL_SCANCODE_KP_MEMDIVIDE
    ScancodeKeypadPlusMinus Scancode = C.SDL_SCANCODE_KP_PLUSMINUS
    ScancodeKeypadClear Scancode = C.SDL_SCANCODE_KP_CLEAR
    ScancodeKeypadClearEntry Scancode = C.SDL_SCANCODE_KP_CLEARENTRY
    ScancodeKeypadBinary Scancode = C.SDL_SCANCODE_KP_BINARY
    ScancodeKeypadOctal Scancode = C.SDL_SCANCODE_KP_OCTAL
    ScancodeKeypadDecimal Scancode = C.SDL_SCANCODE_KP_DECIMAL
    ScancodeKeypadHexadecimal Scancode = C.SDL_SCANCODE_KP_HEXADECIMAL

    ScancodeLCtrl Scancode = C.SDL_SCANCODE_LCTRL
    ScancodeLShift Scancode = C.SDL_SCANCODE_LSHIFT
    ScancodeLAlt Scancode = C.SDL_SCANCODE_LALT
    ScancodeLGUI Scancode = C.SDL_SCANCODE_LGUI
    ScancodeRCtrl Scancode = C.SDL_SCANCODE_RCTRL
    ScancodeRShift Scancode = C.SDL_SCANCODE_RSHIFT
    ScancodeRAlt Scancode = C.SDL_SCANCODE_RALT
    ScancodeRGUI Scancode = C.SDL_SCANCODE_RGUI

    ScancodeMode Scancode = C.SDL_SCANCODE_MODE

    ScancodeAudioNext Scancode = C.SDL_SCANCODE_AUDIONEXT
    ScancodeAudioPrev Scancode = C.SDL_SCANCODE_AUDIOPREV
    ScancodeAudioStop Scancode = C.SDL_SCANCODE_AUDIOSTOP
    ScancodeAudioPlay Scancode = C.SDL_SCANCODE_AUDIOPLAY
    ScancodeAudioMute Scancode = C.SDL_SCANCODE_AUDIOMUTE
    ScancodeMediaSelect Scancode = C.SDL_SCANCODE_MEDIASELECT
    ScancodeWWW Scancode = C.SDL_SCANCODE_WWW
    ScancodeMail Scancode = C.SDL_SCANCODE_MAIL
    ScancodeCalculator Scancode = C.SDL_SCANCODE_CALCULATOR
    ScancodeComputer Scancode = C.SDL_SCANCODE_COMPUTER
    ScancodeAppControlSearch Scancode = C.SDL_SCANCODE_AC_SEARCH
    ScancodeAppControlHome Scancode = C.SDL_SCANCODE_AC_HOME
    ScancodeAppControlBack Scancode = C.SDL_SCANCODE_AC_BACK
    ScancodeAppControlForward Scancode = C.SDL_SCANCODE_AC_FORWARD
    ScancodeAppControlStop Scancode = C.SDL_SCANCODE_AC_STOP
    ScancodeAppControlRefresh Scancode = C.SDL_SCANCODE_AC_REFRESH
    ScancodeAppControlBookmarks Scancode = C.SDL_SCANCODE_AC_BOOKMARKS

    ScancodeBrightnessDown Scancode = C.SDL_SCANCODE_BRIGHTNESSDOWN
    ScancodeBrightnessUp Scancode = C.SDL_SCANCODE_BRIGHTNESSUP
    ScancodeDisplaySwitch Scancode = C.SDL_SCANCODE_DISPLAYSWITCH
    ScancodeKeyboardIllumToggle Scancode = C.SDL_SCANCODE_KBDILLUMTOGGLE
    ScancodeKeyboardIllumDown Scancode = C.SDL_SCANCODE_KBDILLUMDOWN
    ScancodeKeyboardIllumUp Scancode = C.SDL_SCANCODE_KBDILLUMUP
    ScancodeEject Scancode = C.SDL_SCANCODE_EJECT
    ScancodeSleep Scancode = C.SDL_SCANCODE_SLEEP

    ScancodeApp1 Scancode = C.SDL_SCANCODE_APP1
    ScancodeApp2 Scancode = C.SDL_SCANCODE_APP2

		// Not a key, just marks the number of scancodes for array bounds.
    NumScancodes Scancode = C.SDL_NUM_SCANCODES
)
