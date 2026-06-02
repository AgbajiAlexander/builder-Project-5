package main

import (
	"ascii-art/animation"
	"fmt"
)

func main() {
	anim := animation.NewAnimation("HELLO", 20)
	anim.GenerateSpinFrames()
	fmt.Println(anim.Play())

	anim.GenerateWaveFrames()
	fmt.Println(anim.Play())

	anim.GenerateZoomFrames()
	fmt.Println(anim.Play())
}
