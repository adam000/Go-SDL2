package keys

// #include "SDL.h"
import "C"

import (
	"strings"
)

// Mod holds a bitmask of keyboard modifiers (e.g. shift, ctrl, etc.).
type Mod uint16

// HasShift reports whether the shift key (either left or right) is present.
func (mod Mod) HasShift() bool {
	return mod&ModShift != 0
}

// HasCtrl reports whether the control key (either left or right) is present.
func (mod Mod) HasCtrl() bool {
	return mod&ModCtrl != 0
}

// HasAlt reports whether the alt key (either left or right) is present.
func (mod Mod) HasAlt() bool {
	return mod&ModAlt != 0
}

// HasGUI reports whether the GUI key (either left or right) is present.
func (mod Mod) HasGUI() bool {
	return mod&ModGUI != 0
}

// String returns a string like "LShift|LCtrl".
func (mod Mod) String() string {
	if mod == ModNone {
		return "None"
	}
	parts := make([]string, 0, len(modMaskNames))
	for _, mn := range modMaskNames {
		if mod&mn.mask != 0 {
			parts = append(parts, mn.name)
		}
	}
	return strings.Join(parts, "|")
}

// Keyboard modifiers
const (
	ModNone     Mod = C.KMOD_NONE
	ModLShift   Mod = C.KMOD_LSHIFT // left Shift key
	ModRShift   Mod = C.KMOD_RSHIFT // right Shift key
	ModLCtrl    Mod = C.KMOD_LCTRL  // left Control key
	ModRCtrl    Mod = C.KMOD_RCTRL  // right Control key
	ModLAlt     Mod = C.KMOD_LALT   // left Alt key
	ModRAlt     Mod = C.KMOD_RALT   // right Alt key
	ModLGUI     Mod = C.KMOD_LGUI   // left GUI key (often the Windows key)
	ModRGUI     Mod = C.KMOD_RGUI   // right GUI key (often the Windows key)
	ModNum      Mod = C.KMOD_NUM    // Num Lock key
	ModCaps     Mod = C.KMOD_CAPS   // Caps Lock key
	ModMode     Mod = C.KMOD_MODE   // AltGr key
	ModReserved Mod = C.KMOD_RESERVED

	ModCtrl  Mod = C.KMOD_CTRL
	ModShift Mod = C.KMOD_SHIFT
	ModAlt   Mod = C.KMOD_ALT
	ModGUI   Mod = C.KMOD_GUI
)

// modMaskNames is an ordered map of modifier mask to name for Mod.String.
var modMaskNames = [...]struct {
	mask Mod
	name string
}{
	{ModLShift, "LShift"},
	{ModRShift, "RShift"},
	{ModLCtrl, "LCtrl"},
	{ModRCtrl, "RCtrl"},
	{ModLAlt, "LAlt"},
	{ModRAlt, "RAlt"},
	{ModLGUI, "LGUI"},
	{ModRGUI, "RGUI"},
	{ModNum, "Num"},
	{ModCaps, "Caps"},
	{ModMode, "Mode"},
}
