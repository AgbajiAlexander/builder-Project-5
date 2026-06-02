package effects

import (
	"math"
	"strings"
)

// buildWaveFrame creates a single 10-line frame for the wave/scroll effect.
// The text rides on a sine-wave surface that shifts each frame.
func buildWaveFrame(text string, frameIdx, totalFrames, width, height int) string {
	lines := make([]string, height)

	// Phase advances each frame for horizontal scrolling feel
	phase := float64(frameIdx) / float64(totalFrames) * 2 * math.Pi

	// Row 0: top title bar
	title := centreInWidth("~ "+text+" ~", width)
	lines[0] = title

	// Rows 1-8: wave field
	for row := 1; row <= height-2; row++ {
		line := make([]byte, width)
		rowPhase := phase + float64(row)*0.5

		for col := 0; col < width; col++ {
			// sine wave height for this column
			waveY := math.Sin(rowPhase+float64(col)*0.25) * 3.5

			// normalise row to wave coordinate (0 = top, height-2 = bottom)
			normRow := float64(row-1) / float64(height-3) * 7.0 // 0..7

			dist := math.Abs(normRow - (3.5 + waveY))
			switch {
			case dist < 0.5:
				// text chars scroll horizontally along the crest
				charIdx := (col + frameIdx) % len([]rune(text))
				line[col] = text[charIdx%len(text)]
			case dist < 1.2:
				line[col] = '~'
			case dist < 2.0:
				line[col] = '-'
			case dist < 2.8:
				line[col] = '.'
			default:
				line[col] = ' '
			}
		}
		lines[row] = string(line)
	}

	// Row 9: bottom scroll bar showing position
	progress := int(float64(frameIdx) / float64(totalFrames) * float64(width-2))
	bottomBar := "[" + strings.Repeat("=", progress) +
		">" + strings.Repeat(" ", width-3-progress) + "]"
	if len([]rune(bottomBar)) > width {
		bottomBar = bottomBar[:width]
	}
	lines[height-1] = bottomBar

	return strings.Join(lines, "\n") + "\n"
}

// GenerateWave returns `count` frames of the wave/scroll effect.
func GenerateWave(text string, count, width, height int) []string {
	frames := make([]string, count)
	for i := range frames {
		frames[i] = buildWaveFrame(text, i, count, width, height)
	}
	return frames
}
