package effects

import "testing"

func TestSpin(t *testing.T) {
	frames := Spin("ABC", 3)
	if len(frames) != 3 {
		t.Errorf("Spin failed")
	}
}

func TestWave(t *testing.T) {
	frames := Wave("ABC", 5)
	if len(frames) != 5 {
		t.Errorf("Wave failed")
	}
}

func TestZoom(t *testing.T) {
	frames := Zoom("ABC", 6)
	if len(frames) != 6 {
		t.Errorf("Zoom failed")
	}
}
