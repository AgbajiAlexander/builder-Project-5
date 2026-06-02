package animation

import (
	"ascii-art/effects"
)

type Animation struct {
	text   string
	frames int
	data   []string
}

// Constructor
func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		text:   text,
		frames: frames,
		data:   make([]string, 0, frames),
	}
}

// Generate spin frames
func (a *Animation) GenerateSpinFrames() {
	a.data = effects.Spin(a.text, a.frames)
}

// Generate wave frames
func (a *Animation) GenerateWaveFrames() {
	a.data = effects.Wave(a.text, a.frames)
}

// Generate zoom frames
func (a *Animation) GenerateZoomFrames() {
	a.data = effects.Zoom(a.text, a.frames)
}

// Get specific frame
func (a *Animation) GetFrame(index int) string {
	return a.data[index%a.frames]
}

// Play all frames
func (a *Animation) Play() string {
	output := ""
	for i := 0; i < a.frames; i++ {
		output += a.GetFrame(i) + "\n---FRAME END---\n"
	}
	return output
}
