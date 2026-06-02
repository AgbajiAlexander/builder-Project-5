package effects

func padFrame(content string) string {
	lines := []string{}
	for i := 0; i < 10; i++ {
		if i == 5 {
			lines = append(lines, content)
		} else {
			lines = append(lines, "")
		}
	}
	frame := ""
	for _, l := range lines {
		frame += l + "\n"
	}
	return frame
}
