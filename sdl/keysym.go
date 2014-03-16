package sdl

// #include "SDL.h"
import "C"

// KeySym holds the keyboard information from a keyboard event.
type KeySym struct {
	ScanCode int32 // TODO(light): add type
	KeyCode  int32 // TODO(light): add type
	Mod      KeyMod
}

// KeyMod holds a bitmask of keyboard modifiers (e.g. shift, ctrl, etc.).
type KeyMod uint16

// HasShift reports whether the shift key (either left or right) is present.
func (mod KeyMod) HasShift() bool {
	return mod&KeyModShift != 0
}

// HasCtrl reports whether the control key (either left or right) is present.
func (mod KeyMod) HasCtrl() bool {
	return mod&KeyModCtrl != 0
}

// HasAlt reports whether the alt key (either left or right) is present.
func (mod KeyMod) HasAlt() bool {
	return mod&KeyModAlt != 0
}

// HasGUI reports whether the GUI key (either left or right) is present.
func (mod KeyMod) HasGUI() bool {
	return mod&KeyModGUI != 0
}

// Keyboard modifiers
const (
	KeyModNone     KeyMod = C.KMOD_NONE
	KeyModLShift   KeyMod = C.KMOD_LSHIFT // left Shift key
	KeyModRShift   KeyMod = C.KMOD_RSHIFT // right Shift key
	KeyModLCtrl    KeyMod = C.KMOD_LCTRL  // left Control key
	KeyModRCtrl    KeyMod = C.KMOD_RCTRL  // right Control key
	KeyModLAlt     KeyMod = C.KMOD_LALT   // left Alt key
	KeyModRAlt     KeyMod = C.KMOD_RALT   // right Alt key
	KeyModLGUI     KeyMod = C.KMOD_LGUI   // left GUI key (often the Windows key)
	KeyModRGUI     KeyMod = C.KMOD_RGUI   // right GUI key (often the Windows key)
	KeyModNum      KeyMod = C.KMOD_NUM    // Num Lock key
	KeyModCaps     KeyMod = C.KMOD_CAPS   // Caps Lock key
	KeyModMode     KeyMod = C.KMOD_MODE   // AltGr key
	KeyModReserved KeyMod = C.KMOD_RESERVED

	KeyModCtrl  KeyMod = C.KMOD_CTRL
	KeyModShift KeyMod = C.KMOD_SHIFT
	KeyModAlt   KeyMod = C.KMOD_ALT
	KeyModGUI   KeyMod = C.KMOD_GUI
)
