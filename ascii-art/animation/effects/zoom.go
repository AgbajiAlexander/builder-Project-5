package effects

func Zoom(text string, frames int) []string {
	result := []string{}
	maxZoom := 5

	for f := 0; f < frames; f++ {
		zoomLevel := f % (2 * maxZoom)
		if zoomLevel > maxZoom {
			zoomLevel = 2*maxZoom - zoomLevel
		}
		spaces := ""
		for i := 0; i < zoomLevel; i++ {
			spaces += " "
		}
		frame := padFrame(spaces + text + spaces)
		result = append(result, frame)
	}
	return result
}
