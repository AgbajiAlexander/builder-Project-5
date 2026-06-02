package effects

import (
	"math"
	"strings"
)

// bigLetters maps a character to a 3-line tall ASCII art version.
var bigLetters = map[rune][3]string{
	'A': {" /\\ ", "/--\\", "/  \\"},
	'B': {"|\\ ", "|-<", "|/ "},
	'C': {" /-", "|  ", " \\-"},
	'D': {"|\\  ", "| > ", "|/  "},
	'E': {"---", "|-", "---"},
	'F': {"---", "|-", "|  "},
	'G': {" /-", "| -", " \\-"},
	'H': {"|  |", "|--|", "|  |"},
	'I': {"-|-", " | ", "-|-"},
	'J': {"  |", "  |", "\\_|"},
	'K': {"|/ ", "|-<", "|\\ "},
	'L': {"|  ", "|  ", "|__"},
	'M': {"|V|", "| |", "| |"},
	'N': {"|\\ |", "| \\|", "|  |"},
	'O': {" O ", "O O", " O "},
	'P': {"|--", "|-<", "|  "},
	'Q': {" O ", "O O", " Oo"},
	'R': {"|--", "|-<", "|\\ "},
	'S': {" /-", " \\-", "-/ "},
	'T': {"---", " | ", " | "},
	'U': {"|  |", "|  |", " \\/ "},
	'V': {"\\ /", "\\ /", " V "},
	'W': {"| |", "|V|", "   "},
	'X': {"\\ /", " X ", "/ \\"},
	'Y': {"\\ /", " Y ", " | "},
	'Z': {"---", " / ", "---"},
	' ': {"   ", "   ", "   "},
}

// renderBigText renders text as 3-line tall ASCII art.
// Returns 3 strings (one per row).
func renderBigText(text string) [3]string {
	var rows [3]string
	for _, ch := range strings.ToUpper(text) {
		gl, ok := bigLetters[ch]
		if !ok {
			gl = bigLetters[' ']
		}
		for r := 0; r < 3; r++ {
			rows[r] += gl[r] + " "
		}
	}
	return rows
}

// buildZoomFrame creates a single 10-line frame for the zoom in/out effect.
// Scale oscillates smoothly: 0 → max → 0 → … over the total frames.
func buildZoomFrame(text string, frameIdx, totalFrames, width, height int) string {
	lines := make([]string, height)

	// scale: 0.0 (tiny) → 1.0 (full) → 0.0, looping via sine
	t := float64(frameIdx) / float64(totalFrames)
	scale := (math.Sin(t*2*math.Pi-math.Pi/2) + 1.0) / 2.0 // 0..1

	// Background: starfield that gets denser as we zoom out
	starDensity := 1.0 - scale
	rng := newLCG(uint64(frameIdx * 7))

	for row := 0; row < height; row++ {
		buf := make([]byte, width)
		for col := 0; col < width; col++ {
			if rng.nextFloat() < starDensity*0.08 {
				stars := []byte{'.', '*', '+', '\'', '`'}
				buf[col] = stars[rng.nextInt(len(stars))]
			} else {
				buf[col] = ' '
			}
		}
		lines[row] = string(buf)
	}

	// Render text block, scaled
	bigRows := renderBigText(text)
	bigWidth := len([]rune(bigRows[0]))
	bigHeight := 3

	scaledW := int(math.Round(float64(bigWidth) * scale))
	scaledH := int(math.Round(float64(bigHeight) * scale))
	if scaledW < 1 {
		scaledW = 1
	}
	if scaledH < 1 {
		scaledH = 1
	}

	// Place in centre of canvas
	startRow := (height - scaledH) / 2
	startCol := (width - scaledW) / 2

	for sr := 0; sr < scaledH; sr++ {
		// Map scaled row back to big-text row
		srcRow := int(float64(sr)/float64(scaledH)*float64(bigHeight)) % bigHeight
		srcLine := []rune(bigRows[srcRow])

		row := startRow + sr
		if row < 0 || row >= height {
			continue
		}

		lineBuf := []rune(lines[row])
		for sc := 0; sc < scaledW; sc++ {
			col := startCol + sc
			if col < 0 || col >= width {
				continue
			}
			srcCol := int(float64(sc)/float64(scaledW)*float64(len(srcLine))) % len(srcLine)
			ch := srcLine[srcCol]
			if ch != ' ' {
				// Make far-away text dimmer
				if scale < 0.4 {
					lineBuf[col] = '·'
				} else if scale < 0.7 {
					lineBuf[col] = '+'
				} else {
					lineBuf[col] = ch
				}
			}
		}
		lines[row] = string(lineBuf)
	}

	// Zoom indicator at bottom
	barWidth := width - 2
	filled := int(scale * float64(barWidth))
	if filled > barWidth {
		filled = barWidth
	}
	zoomBar := "[" + strings.Repeat("#", filled) + strings.Repeat(".", barWidth-filled) + "]"
	lines[height-1] = zoomBar

	return strings.Join(lines, "\n") + "\n"
}

// GenerateZoom returns `count` frames of the zoom in/out effect.
func GenerateZoom(text string, count, width, height int) []string {
	frames := make([]string, count)
	for i := range frames {
		frames[i] = buildZoomFrame(text, i, count, width, height)
	}
	return frames
}

// --- minimal deterministic LCG pseudo-random for frame content ---

type lcg struct{ state uint64 }

func newLCG(seed uint64) *lcg { return &lcg{state: seed + 1} }
func (l *lcg) next() uint64 {
	l.state = l.state*6364136223846793005 + 1442695040888963407
	return l.state
}
func (l *lcg) nextFloat() float64 { return float64(l.next()>>11) / (1 << 53) }
func (l *lcg) nextInt(n int) int  { return int(l.next() % uint64(n)) }
