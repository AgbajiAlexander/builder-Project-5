package main

import (
	"ascii-art/animation"
	"fmt"
)

func main() {
	anim := animation.NewAnimation("HELLO", 20)

	// Spin demo
	anim.GenerateSpinFrames()
	fmt.Println(anim.Play())

	// Wave demo
	anim.GenerateWaveFrames()
	fmt.Println(anim.Play())

	// Zoom demo
	anim.GenerateZoomFrames()
	fmt.Println(anim.Play())
}
