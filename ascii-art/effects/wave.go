package effects

import (
	"math"
)

func Wave(text string, frames int) []string {
	result := []string{}
	n := len(text)

	for f := 0; f < frames; f++ {
		lines := make([]string, 10)
		for i := 0; i < n; i++ {
			offset := int(math.Sin(float64(f+i)*0.5) * 3) // vertical shift
			lineIndex := 5 + offset
			if lineIndex < 0 {
				lineIndex = 0
			}
			if lineIndex > 9 {
				lineIndex = 9
			}
			lines[lineIndex] += string(text[i])
		}
		frame := ""
		for _, l := range lines {
			frame += l + "\n"
		}
		result = append(result, frame)
	}
	return result
}
