package effects

import (
	"math"
	"strings"
)

// spinChars cycles through characters that give a rotation feel.
var spinChars = []string{"|", "/", "-", "\\"}

// rotateLine rotates the characters in a string by the given phase,
// replacing each char with a spin-cycle character proportional to its position.
func rotateLine(text string, phase int) string {
	runes := []rune(text)
	result := make([]rune, len(runes))
	for i, r := range runes {
		if r == ' ' {
			result[i] = ' '
		} else {
			idx := (i + phase) % len(spinChars)
			result[i] = []rune(spinChars[idx])[0]
		}
	}
	return string(result)
}

// buildSpinFrame creates a single 10-line frame for the spin effect.
// The middle rows show the text with rotating characters; borders pulse with phase.
func buildSpinFrame(text string, frameIdx, totalFrames, width, height int) string {
	phase := frameIdx % len(spinChars)
	angle := float64(frameIdx) / float64(totalFrames) * 2 * math.Pi

	// Border character cycles through spin chars
	borderChar := spinChars[phase]

	lines := make([]string, height)

	// Top border (row 0)
	lines[0] = strings.Repeat(borderChar, width)

	// Row 1: side borders + decorative dots
	dot := spinChars[(phase+1)%len(spinChars)]
	inner1 := strings.Repeat(dot, width-2)
	lines[1] = borderChar + inner1 + borderChar

	// Rows 2-3: empty padding
	for r := 2; r <= 3; r++ {
		pad := strings.Repeat(" ", width-2)
		lines[r] = borderChar + pad + borderChar
	}

	// Row 4: centred text with spin overlay
	rotated := rotateLine(text, phase)
	textRow := centreInWidth(rotated, width-2)
	lines[4] = borderChar + textRow + borderChar

	// Row 5: shadow/echo of text shifted by 1 phase step
	echo := rotateLine(text, (phase+1)%len(spinChars))
	echoFaded := fadeLine(echo)
	echoRow := centreInWidth(echoFaded, width-2)
	lines[5] = borderChar + echoRow + borderChar

	// Rows 6-7: decorative spin wheel rows
	for r := 6; r <= 7; r++ {
		wheelPad := buildWheelRow(width-2, angle, r)
		lines[r] = borderChar + wheelPad + borderChar
	}

	// Row 8: mirror of row 1
	lines[8] = borderChar + inner1 + borderChar

	// Row 9: Bottom border
	lines[9] = strings.Repeat(borderChar, width)

	return strings.Join(lines, "\n") + "\n"
}

// buildWheelRow creates a decorative row that shifts with the angle to simulate rotation.
func buildWheelRow(innerWidth int, angle float64, rowIdx int) string {
	buf := make([]byte, innerWidth)
	for i := range buf {
		v := math.Sin(angle+float64(i)*0.4+float64(rowIdx)*0.7) * 1.5
		switch {
		case v > 1.0:
			buf[i] = 'O'
		case v > 0.3:
			buf[i] = 'o'
		case v > -0.3:
			buf[i] = '.'
		case v > -1.0:
			buf[i] = ','
		default:
			buf[i] = ' '
		}
	}
	return string(buf)
}

// fadeLine replaces non-space chars with lighter symbols to create an echo.
func fadeLine(s string) string {
	r := []rune(s)
	for i, c := range r {
		switch c {
		case '|', '/', '-', '\\':
			r[i] = '.'
		}
	}
	return string(r)
}

// centreInWidth pads text to exactly n chars, centred.
func centreInWidth(text string, n int) string {
	t := []rune(text)
	if len(t) >= n {
		return string(t[:n])
	}
	total := n - len(t)
	left := total / 2
	right := total - left
	return strings.Repeat(" ", left) + string(t) + strings.Repeat(" ", right)
}

// GenerateSpin returns `count` frames of the spinning effect.
func GenerateSpin(text string, count, width, height int) []string {
	frames := make([]string, count)
	for i := range frames {
		frames[i] = buildSpinFrame(text, i, count, width, height)
	}
	return frames
}
