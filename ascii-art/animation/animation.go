package animation

import "ascii-art/animation/effects"

type Animation struct {
	text   string
	frames int
	data   []string
}

func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		text:   text,
		frames: frames,
		data:   make([]string, 0, frames),
	}
}

func (a *Animation) GenerateSpinFrames() {
	a.data = effects.Spin(a.text, a.frames)
}

func (a *Animation) GenerateWaveFrames() {
	a.data = effects.Wave(a.text, a.frames)
}

func (a *Animation) GenerateZoomFrames() {
	a.data = effects.Zoom(a.text, a.frames)
}

func (a *Animation) GetFrame(index int) string {
	return a.data[index%a.frames]
}

func (a *Animation) Play() string {
	output := ""
	for i := 0; i < a.frames; i++ {
		output += a.GetFrame(i) + "\n---FRAME END---\n"
	}
	return output
}
