package sdl

import "testing"

func TestInitEverything(t *testing.T) {
	if INIT_EVERYTHING != INIT_TIMER | INIT_AUDIO | INIT_VIDEO |
		INIT_JOYSTICK | INIT_HAPTIC | INIT_GAMECONTROLLER |
		INIT_EVENTS | INIT_NOPARACHUTE {
		t.Errorf("INIT_EVERYTHING is not the sum of all the other flags")
	}
}
