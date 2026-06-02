package main

import (
	"fmt"
	"os"
	"time"

	"ascii-art/animation"
)

func printDivider(label string) {
	fmt.Printf("\n════════════════════════════════════════\n")
	fmt.Printf("  %s\n", label)
	fmt.Printf("════════════════════════════════════════\n\n")
}

func playLive(anim *animation.Animation, effect string, loops int) {
	printDivider(fmt.Sprintf("EFFECT: %s  |  text: %q", effect, anim.Text))
	frames := anim.Frames
	for loop := 0; loop < loops; loop++ {
		for i := 0; i < frames; i++ {
			// Move cursor up FrameHeight lines to overwrite previous frame
			if loop != 0 || i != 0 {
				fmt.Printf("\033[%dA", 10)
			}
			fmt.Print(anim.GetFrame(i))
			time.Sleep(120 * time.Millisecond)
		}
	}
}

func demoPlayOutput(anim *animation.Animation, effect string) {
	printDivider(fmt.Sprintf("Play() output — %s", effect))
	output := anim.Play()
	// Print first 3 frames only to keep demo readable
	maxChars := 10 * 41 * 3
	if len(output) > maxChars {
		fmt.Print(output[:maxChars])
		fmt.Println("\n... (truncated for demo)")
	} else {
		fmt.Print(output)
	}
}

func main() {
	interactive := len(os.Args) > 1 && os.Args[1] == "--live"
	text := "ASCII"
	if len(os.Args) > 2 {
		text = os.Args[2]
	}

	// ── SPIN ──────────────────────────────────────────────────────────────────
	spinAnim := animation.NewAnimation(text, 12)
	spinAnim.GenerateSpinFrames()

	// ── WAVE ──────────────────────────────────────────────────────────────────
	waveAnim := animation.NewAnimation(text, 16)
	waveAnim.GenerateWaveFrames()

	// ── ZOOM ──────────────────────────────────────────────────────────────────
	zoomAnim := animation.NewAnimation(text, 20)
	zoomAnim.GenerateZoomFrames()

	if interactive {
		// Live terminal animation
		playLive(spinAnim, "SPIN", 3)
		time.Sleep(400 * time.Millisecond)
		playLive(waveAnim, "WAVE", 3)
		time.Sleep(400 * time.Millisecond)
		playLive(zoomAnim, "ZOOM", 3)
		fmt.Println("\nDone!")
		return
	}

	// ── Static demo ───────────────────────────────────────────────────────────
	fmt.Println("ASCII Art Animation System")
	fmt.Println("Run with --live [TEXT] for animated terminal playback.")
	fmt.Println()

	// Show first frame of each effect
	fmt.Println("── SPIN frame 0 ──")
	fmt.Print(spinAnim.GetFrame(0))
	fmt.Println()

	fmt.Println("── WAVE frame 0 ──")
	fmt.Print(waveAnim.GetFrame(0))
	fmt.Println()

	fmt.Println("── ZOOM frame 0 ──")
	fmt.Print(zoomAnim.GetFrame(0))
	fmt.Println()

	// Demo Play() output for spin
	demoPlayOutput(spinAnim, "SPIN")

	// Demonstrate seamless looping
	fmt.Println("\n── Loop check: GetFrame(0) == GetFrame(Frames) ──")
	effects := []struct {
		name string
		anim *animation.Animation
	}{
		{"SPIN", spinAnim},
		{"WAVE", waveAnim},
		{"ZOOM", zoomAnim},
	}
	for _, e := range effects {
		f0 := e.anim.GetFrame(0)
		fN := e.anim.GetFrame(e.anim.Frames) // wraps to 0
		match := f0 == fN
		fmt.Printf("  %s loop seamless: %v\n", e.name, match)
	}
}
