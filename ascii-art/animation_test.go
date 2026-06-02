package animation

import (
	"strings"
	"testing"
)

func TestNewAnimation(t *testing.T) {
	a := NewAnimation("HELLO", 8)
	if a.Text != "HELLO" {
		t.Errorf("expected Text=HELLO, got %q", a.Text)
	}
	if a.Frames != 8 {
		t.Errorf("expected Frames=8, got %d", a.Frames)
	}
}

func TestNewAnimationMinFrames(t *testing.T) {
	a := NewAnimation("X", 0)
	if a.Frames != 1 {
		t.Errorf("expected at least 1 frame, got %d", a.Frames)
	}
}

func frameCount(s string) int {
	return strings.Count(s, "\n")/FrameHeight + 1
}

// eachFrameIs10Lines checks every frame in the frames slice is exactly 10 lines.
func eachFrameIs10Lines(t *testing.T, frames []string) {
	t.Helper()
	for i, f := range frames {
		lines := strings.Split(strings.TrimRight(f, "\n"), "\n")
		if len(lines) != FrameHeight {
			t.Errorf("frame %d: expected %d lines, got %d", i, FrameHeight, len(lines))
		}
	}
}

// --- Spin ---

func TestGenerateSpinFramesCount(t *testing.T) {
	a := NewAnimation("GO", 12)
	a.GenerateSpinFrames()
	if got := len(a.frames); got != 12 {
		t.Errorf("expected 12 spin frames, got %d", got)
	}
}

func TestGenerateSpinFramesHeight(t *testing.T) {
	a := NewAnimation("GO", 8)
	a.GenerateSpinFrames()
	eachFrameIs10Lines(t, a.frames)
}

// --- Wave ---

func TestGenerateWaveFramesCount(t *testing.T) {
	a := NewAnimation("WAVE", 16)
	a.GenerateWaveFrames()
	if got := len(a.frames); got != 16 {
		t.Errorf("expected 16 wave frames, got %d", got)
	}
}

func TestGenerateWaveFramesHeight(t *testing.T) {
	a := NewAnimation("WAVE", 8)
	a.GenerateWaveFrames()
	eachFrameIs10Lines(t, a.frames)
}

// --- Zoom ---

func TestGenerateZoomFramesCount(t *testing.T) {
	a := NewAnimation("ZOOM", 10)
	a.GenerateZoomFrames()
	if got := len(a.frames); got != 10 {
		t.Errorf("expected 10 zoom frames, got %d", got)
	}
}

func TestGenerateZoomFramesHeight(t *testing.T) {
	a := NewAnimation("ZOOM", 6)
	a.GenerateZoomFrames()
	eachFrameIs10Lines(t, a.frames)
}

// --- GetFrame ---

func TestGetFrameWrap(t *testing.T) {
	a := NewAnimation("LOOP", 4)
	a.GenerateWaveFrames()
	f0 := a.GetFrame(0)
	fN := a.GetFrame(4) // should wrap to 0
	if f0 != fN {
		t.Error("GetFrame should wrap seamlessly: frame 0 != frame N")
	}
}

func TestGetFrameEmpty(t *testing.T) {
	a := NewAnimation("X", 4)
	// no Generate* called
	if got := a.GetFrame(0); got != "" {
		t.Errorf("expected empty string for uninitialised animation, got %q", got)
	}
}

// --- Play ---

func TestPlayContainsAllFrames(t *testing.T) {
	a := NewAnimation("PLAY", 6)
	a.GenerateSpinFrames()
	out := a.Play()
	// Delimiter count should be frames-1
	delimiters := strings.Count(out, "--- FRAME")
	if delimiters != 5 {
		t.Errorf("expected 5 frame delimiters, got %d", delimiters)
	}
}

func TestPlayEmptyAnimation(t *testing.T) {
	a := NewAnimation("X", 4)
	if got := a.Play(); got != "" {
		t.Errorf("expected empty Play() output, got %q", got)
	}
}
