package effects

import (
	"strings"
	"testing"
)

const (
	testWidth  = 40
	testHeight = 10
)

// countLines returns the number of lines in a frame string.
func countLines(frame string) int {
	return len(strings.Split(strings.TrimRight(frame, "\n"), "\n"))
}

// --- GenerateSpin ---

func TestGenerateSpinLength(t *testing.T) {
	frames := GenerateSpin("TEST", 8, testWidth, testHeight)
	if len(frames) != 8 {
		t.Errorf("expected 8 frames, got %d", len(frames))
	}
}

func TestGenerateSpinFrameHeight(t *testing.T) {
	frames := GenerateSpin("TEST", 4, testWidth, testHeight)
	for i, f := range frames {
		n := countLines(f)
		if n != testHeight {
			t.Errorf("spin frame %d: expected %d lines, got %d", i, testHeight, n)
		}
	}
}

func TestGenerateSpinFrameWidth(t *testing.T) {
	frames := GenerateSpin("HI", 4, testWidth, testHeight)
	for i, f := range frames {
		for j, line := range strings.Split(strings.TrimRight(f, "\n"), "\n") {
			if len([]rune(line)) != testWidth {
				t.Errorf("spin frame %d line %d: expected width %d, got %d",
					i, j, testWidth, len([]rune(line)))
			}
		}
	}
}

func TestGenerateSpinLoops(t *testing.T) {
	// Last frame + 1 step should look like first frame (same phase)
	frames := GenerateSpin("GO", 4, testWidth, testHeight)
	// Phase is frameIdx % 4, so frame[0] and frame[4] would be identical if we had 5 frames
	frames2 := GenerateSpin("GO", 8, testWidth, testHeight)
	if frames[0] != frames2[4] {
		t.Error("spin animation does not loop seamlessly")
	}
}

// --- GenerateWave ---

func TestGenerateWaveLength(t *testing.T) {
	frames := GenerateWave("HELLO", 16, testWidth, testHeight)
	if len(frames) != 16 {
		t.Errorf("expected 16 frames, got %d", len(frames))
	}
}

func TestGenerateWaveFrameHeight(t *testing.T) {
	frames := GenerateWave("HELLO", 6, testWidth, testHeight)
	for i, f := range frames {
		n := countLines(f)
		if n != testHeight {
			t.Errorf("wave frame %d: expected %d lines, got %d", i, testHeight, n)
		}
	}
}

func TestGenerateWaveFrameWidth(t *testing.T) {
	frames := GenerateWave("HI", 4, testWidth, testHeight)
	for i, f := range frames {
		for j, line := range strings.Split(strings.TrimRight(f, "\n"), "\n") {
			if len([]rune(line)) != testWidth {
				t.Errorf("wave frame %d line %d: expected width %d, got %d",
					i, j, testWidth, len([]rune(line)))
			}
		}
	}
}

func TestGenerateWaveLoops(t *testing.T) {
	frames := GenerateWave("LOOP", 8, testWidth, testHeight)
	frames2 := GenerateWave("LOOP", 16, testWidth, testHeight)
	if frames[0] != frames2[8] {
		t.Error("wave animation does not loop seamlessly")
	}
}

// --- GenerateZoom ---

func TestGenerateZoomLength(t *testing.T) {
	frames := GenerateZoom("ZOOM", 10, testWidth, testHeight)
	if len(frames) != 10 {
		t.Errorf("expected 10 frames, got %d", len(frames))
	}
}

func TestGenerateZoomFrameHeight(t *testing.T) {
	frames := GenerateZoom("ZOOM", 6, testWidth, testHeight)
	for i, f := range frames {
		n := countLines(f)
		if n != testHeight {
			t.Errorf("zoom frame %d: expected %d lines, got %d", i, testHeight, n)
		}
	}
}

func TestGenerateZoomFrameWidth(t *testing.T) {
	frames := GenerateZoom("ZM", 4, testWidth, testHeight)
	for i, f := range frames {
		for j, line := range strings.Split(strings.TrimRight(f, "\n"), "\n") {
			if len([]rune(line)) != testWidth {
				t.Errorf("zoom frame %d line %d: expected width %d, got %d",
					i, j, testWidth, len([]rune(line)))
			}
		}
	}
}

func TestGenerateZoomLoops(t *testing.T) {
	frames := GenerateZoom("X", 8, testWidth, testHeight)
	frames2 := GenerateZoom("X", 16, testWidth, testHeight)
	if frames[0] != frames2[8] {
		t.Error("zoom animation does not loop seamlessly")
	}
}

// --- Helpers ---

func TestCentreInWidth(t *testing.T) {
	result := centreInWidth("AB", 10)
	if len([]rune(result)) != 10 {
		t.Errorf("centreInWidth: expected length 10, got %d", len([]rune(result)))
	}
	if !strings.Contains(result, "AB") {
		t.Error("centreInWidth: result should contain original text")
	}
}

func TestCentreInWidthTruncates(t *testing.T) {
	result := centreInWidth("ABCDEFGHIJ", 5)
	if len([]rune(result)) != 5 {
		t.Errorf("centreInWidth truncation: expected length 5, got %d", len([]rune(result)))
	}
}

func TestLCG(t *testing.T) {
	rng := newLCG(42)
	seen := make(map[uint64]bool)
	for i := 0; i < 1000; i++ {
		v := rng.next()
		seen[v] = true
	}
	// LCG should produce many distinct values
	if len(seen) < 990 {
		t.Errorf("LCG produced too few distinct values: %d", len(seen))
	}
}
