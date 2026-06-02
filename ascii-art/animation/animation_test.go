package animation

import "testing"

func TestNewAnimation(t *testing.T) {
	a := NewAnimation("TEST", 10)
	if a.text != "TEST" || a.frames != 10 {
		t.Errorf("NewAnimation failed")
	}
}
