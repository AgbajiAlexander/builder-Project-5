package effects

func Spin(text string, frames int) []string {
	result := []string{}
	n := len(text)

	for i := 0; i < frames; i++ {
		rotated := text[i%n:] + text[:i%n]
		frame := padFrame(rotated)
		result = append(result, frame)
	}
	return result
}
