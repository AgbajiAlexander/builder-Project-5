package animation

import (
	"fmt"
	"strings"

	"ascii-art/effects"
)

const (
	FrameHeight = 10
	FrameWidth  = 40
)

// Animation holds the text, frame count, and generated frames.
type Animation struct {
	Text   string
	Frames int
	frames []string
}

// NewAnimation creates a new Animation with the given text and frame count.
func NewAnimation(text string, frames int) *Animation {
	if frames < 1 {
		frames = 1
	}
	return &Animation{
		Text:   text,
		Frames: frames,
	}
}

// GenerateSpinFrames fills the animation with a spinning rotation effect.
func (a *Animation) GenerateSpinFrames() {
	a.frames = effects.GenerateSpin(a.Text, a.Frames, FrameWidth, FrameHeight)
}

// GenerateWaveFrames fills the animation with a wave scroll effect.
func (a *Animation) GenerateWaveFrames() {
	a.frames = effects.GenerateWave(a.Text, a.Frames, FrameWidth, FrameHeight)
}

// GenerateZoomFrames fills the animation with a zoom in/out effect.
func (a *Animation) GenerateZoomFrames() {
	a.frames = effects.GenerateZoom(a.Text, a.Frames, FrameWidth, FrameHeight)
}

// GetFrame returns the frame at the given index (wraps around for seamless looping).
func (a *Animation) GetFrame(index int) string {
	if len(a.frames) == 0 {
		return ""
	}
	return a.frames[index%len(a.frames)]
}

// Play returns all frames joined with delay markers between them.
func (a *Animation) Play() string {
	if len(a.frames) == 0 {
		return ""
	}
	var sb strings.Builder
	for i, frame := range a.frames {
		sb.WriteString(frame)
		if i < len(a.frames)-1 {
			sb.WriteString(fmt.Sprintf("\n--- FRAME %d/%d ---\n", i+1, len(a.frames)))
		}
	}
	return sb.String()
}
